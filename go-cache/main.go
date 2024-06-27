package main

import (
	"fmt"
	"github.com/patrickmn/go-cache"
	"time"
)

const (
	DefaultExpiration = 10 * time.Second
	ClearInterval     = 1 * time.Hour
)

func main() {
	lc := cache.New(DefaultExpiration, ClearInterval)
	lc.Set("foo", "bar", cache.DefaultExpiration)
	lc.Set("baz", 42, cache.NoExpiration)

	var foo interface{}
	var found bool
	// 获取值
	foo, found = lc.Get("foo")
	if found {
		fmt.Println(foo)
	}

	var fooStr string
	// 获取值， 并断言
	if x, found := lc.Get("foo"); found {
		fooStr = x.(string)
		fmt.Println(fooStr)
	}
	// 对结构体指针进行操作
	var my *MyStruct
	lc.Set("foo", &MyStruct{Name: "NameName"}, cache.DefaultExpiration)
	if x, found := lc.Get("foo"); found {
		my = x.(*MyStruct)
		// ...
	}
	fmt.Println(my)
}

type MyStruct struct {
	Name string
}
