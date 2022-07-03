使用go-redis实例化客户端
key使用google/uuid生成，而value全部使用单个字符0，测试以下数量大小的数据
1w
5w
10w
20w
30w
40w
50w
每次测试完直接执行info memory查看当前内存信息，并将结果记录到./test02的对应文件中
执行flushdb命令清空缓存

