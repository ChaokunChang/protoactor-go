package main

import (
	"log"

	console "github.com/AsynkronIT/goconsole"
	"github.com/ChaokunChang/protoactor-go/actor"
	"github.com/ChaokunChang/protoactor-go/cluster"
	"github.com/ChaokunChang/protoactor-go/cluster/consul"
	"github.com/ChaokunChang/protoactor-go/examples/cluster/shared"
	"github.com/ChaokunChang/protoactor-go/remote"
)

func main() {
	// this node knows about Hello kind
	remote.Register("Hello", actor.PropsFromProducer(func() actor.Actor {
		return &shared.HelloActor{}
	}))

	cp, err := consul.New()
	if err != nil {
		log.Fatal(err)
	}
	cluster.Start("mycluster", "127.0.0.1:8080", cp)

	hello := shared.GetHelloGrain("MyGrain")

	res, err := hello.SayHello(&shared.HelloRequest{Name: "Roger"})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Message from grain: %v", res.Message)
	console.ReadLine()

	cluster.Shutdown(true)
}
