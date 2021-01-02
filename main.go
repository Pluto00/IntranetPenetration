package main

import (
	"IntranetPenetration/client"
	"IntranetPenetration/server"
	"flag"
	"fmt"
)

var host *string = flag.String("host", "127.0.0.1", "请输入服务器ip")
var remotePort *string = flag.String("remotePort", "20012", "服务器地址端口")
var localPort *string = flag.String("localPort", "80", "本地端口")
var mode *string = flag.String("mode", "client", "服务端/客户端")

func main() {
	flag.Parse()
	fmt.Printf("args=%s, num=%d\n", flag.Args(), flag.NArg())
	fmt.Println(*host, *localPort, *remotePort)
	if *mode == "server" {
		server.Run(localPort, remotePort)
	} else {
		client.Run(localPort, remotePort, host)
	}
}
