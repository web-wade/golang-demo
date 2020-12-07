package main

import (
	"golang-demo/crawler/engine"
	"golang-demo/crawler/scheduler"
	"golang-demo/crawler/types"
	"golang-demo/crawler/zhenai/parser"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
	}
	e.Run(types.Request{
		Url:       "http://www.zhenai.com/zhenghun",
		ParseFunc: parser.ParseCityList,
	})
}
