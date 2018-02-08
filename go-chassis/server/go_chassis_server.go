package main

import (
	"context"
	"flag"
	"fmt"
	"runtime"
	"time"

	"github.com/ServiceComb/go-chassis"
	_ "github.com/ServiceComb/go-chassis/core/registry/file"
	"github.com/rpcx-ecosystem/rpcx-benchmark/go-chassis/schema"
)

var (
	delay = flag.Duration("delay", 0, "delay to mock business processing")
)

type Hello struct{}

func (h *Hello) Say(ctx context.Context, args *schema.BenchmarkMessage) (*schema.BenchmarkMessage, error) {
	args.Field1 = "OK"
	args.Field2 = 100
	reply := *args
	if *delay > 0 {
		time.Sleep(*delay)
	} else {
		runtime.Gosched()
	}
	return &reply, nil
}

func main() {
	flag.Parse()

	chassis.RegisterSchema("highway", &Hello{})
	if err := chassis.Init(); err != nil {
		fmt.Printf("Init failed %s", err)
		return
	}
	chassis.Run()
}
