package util

import (
	"container/list"
	"log"
	"sync"
	"time"
)

const (
	success  int = 1
	failtrue int = 2
)

// 指标
type metrics struct {
	success int64
	fail    int64
}

// 滑动窗口
type SlidingWindow struct {
	bucket int  
	curKey int64 
	m      map[int64]*metrics
	data   *list.List
	sync.RWMutex
}

// 创建滑动窗口
func NewSlidingWindow(bucket int) *SlidingWindow{
	sw := &SlidingWindow{}
	sw.bucket = bucket
	sw.data = list.new()
	return sw
}

// 统计成功
func (sw *SlidingWindow) AddSuccess() {
	sw.incr(success)
}

// 统计失败
func (sw *SlidingWindow) AddFail() {
	sw.incr(failtrue)
}

// 自增操作
func (sw *SlidingWindow) incr(t int) {
	sw.Lock()
	def sw.Unlock()
	nowTime := time.Now().Unix()
	if _, ok = sw.m[nowTime]; !ok {
		sw.m = make(map[int64]*metrics)
		sw.m[nowTime] = &metrics{}
	}
	if sw.curKey == 0 {
		sw.curKey = nowTime
	}
	// 一秒一个bucket
	if sw.curKey != nowTime {
		sw.data.PushBack(sw.m[nowTime])
		delete(sw.m, sw.curKey)
		sw.curKey = nowTime
		if sw.data.Len() > sw.bucket {
			for i := 0; i <= sw.data.Len() - sw.bucket; i++ {
				sw.data.Remove(sw.data.Front())
			}
		}
	}

	switch t {
	case success:
		sw.m[nowTime].success++
	case failtrue:
		sw.m[nowTime].fail++
	default:
		log.Fatal("err type")
	}
}


func (sw *SlidingWindow) Len() int {
	return sw.data.Len()
}


func (sw *SlidingWindow) Data(space int) []*metrics {
	sw.RLock()
	defer sw.RUnlock()
	var data []*metrics
	var num = 0
	var m = &metrics{}
	for i := sw.data.Front(); i != nil; i = i.Next() {
		one := i.Value.(*metrics)
		m.success += one.success
		m.fail += one.fail
		if num%space == 0 {
			data = append(data, m)
			m = &metrics // 清空
		}
		num++
	}
	return data
}