## remote_dog

本包是为了通过http调用newdog开发的。 使用方式参考main函数。

## 接口

1. NewApi() 创建调用类
2. SetBasicAuth() 设置basicauth
3. SetCommand 在dog上设置一个命令
4. StopCommand 在dog上停止一个命令
5. GetCommand在dog上获取命令列表
6. Download让dag下载某个命令的程序
7. Ping dog心跳检测
8. Status dog运行程序的状态
9. UserCurrent运行dog的用户
10. Version 运行dog的版本，git的提交信息
11. GetAllWebhook 获取所有的webhook
12. GetWebhook获取某一个命令的webhook
13.SetWebhook设置的webhook
14 .DelWebhook删除webhook
15.GetKeepLive获取保活列表
16.KeepLive设置保活
17. Upload上传文件到dog
18. IpList dog的iplist，节点发现下获取的iplist
19. OperationLog操作日志
20. Plugin_list dog上运行的插件及其端口
21. SetIpNickname设置当前dog上发现的节点的ip昵称
22. DeleteServer 删除节点发现的节点列表中某一个不在线的节点
23. Relay让同集群的某一个dog执行开始或者结束
24. Import通过文件导入命令
25. Export导出dog上的命令



