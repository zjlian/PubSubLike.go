# PubSubLike.go
Golang 发布订阅系统     
自用的     
相比常见的发布订阅系统，每个消息的绑定函数可以有结果返回，在触发消息后，Publish函数将会返回该与事件绑定的所有处理函数返回值作为一个集合返回
