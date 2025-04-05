# gateway

一个简单的负载均衡器，支持通过服务名进行路由。将http、websocket数据转发到对应节点。

优势是全部使用零拷贝，性能会非常好。

预计能节约用户1d的工时，可用于摸鱼再告诉Leader写完了。

使用：调用gateway.Run，传入路由、url->服务名映射。
选项：日志库