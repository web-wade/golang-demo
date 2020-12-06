package engine

import (
	"golang-demo/crawler/types"
	"golang-demo/crawler/zhenai/parser"
	"testing"
)

func TestSimpleEngine_Run(t *testing.T) {
	SimpleEngine{}.Run(types.Request{
		Url:       "http://www.zhenai.com/zhenghun",
		ParseFunc: parser.ParseCityList,
	})
}
