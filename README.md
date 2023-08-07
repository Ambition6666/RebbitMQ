# coderzh.github.io
```go
//先初始化下mq用户名,密码,ip地址,端口
//调用InitMqConfig
InitMqConfig(username,password,host,port)
//然后获取mq的对象,调用函数
//生产者
  rm := mq.NewRabbitMQPubSub("nb")
	rm.PublishPub("hello")
	defer rm.Destory()
//消费者
  rm := mq.NewRabbitMQPubSub("nb")
	mes := rm.RecieveSub()
	go func() {
		for v := range mes {
			fmt.Printf("%s\n", v.Body)
		}
	}()
	select {}
```
### 支持简单调用,worker,pubsub,route,topic五种调用
