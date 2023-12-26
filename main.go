package main

import (
	"context"

	"git.code.oa.com/trpc-go/trpc-go/client"
	"git.code.oa.com/trpc-go/trpc-go/log"

	pb "git.code.oa.com/trpcprotocol/test/helloworld" // 被调服务的协议生成文件 pb.go 的 git 地址，没有 push 到 git 的话，可以在 gomod 里面 replace 本地路径进行引用，如 gomod 里面加上一行：replace "git.code.oa.com/trpcprotocol/test/helloworld" => ./你的本地桩代码路径
)

func main() {
	proxy := pb.NewGreeterClientProxy()                                                             // 创建一个客户端调用代理，名词解释见客户端开发文档
	req := &pb.HelloRequest{Msg: "Hello, I am tRPC-Go client."}                                     // 填充请求参数
	rsp, err := proxy.SayHello(context.Background(), req, client.WithTarget("ip://127.0.0.1:8000")) // 调用目标地址为前面启动的服务监听的地址
	if err != nil {
		log.Errorf("could not greet: %v", err)
		return
	}
	log.Debugf("response: %v", rsp)
}
