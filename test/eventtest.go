package main

import (
	"errors"
	"event"
	"log"
)

func testHandler1(text interface{}) (interface{}, error) {
	log.Println(text, 1)
	return nil, nil
}
func testHandler2(text interface{}) (interface{}, error) {
	log.Println(text, 2)
	return nil, errors.New("testHandler2 错误处理测试")
}
func testHandler3(text interface{}) (interface{}, error) {
	log.Println(text, 3)
	return text, nil
}

func test() {
	results, perfect := event.Publish("TEST", "TEST TEST")
	if !perfect {
		log.Println(event.Error())
	}
	log.Println(results)

	// 触发不存在事件
	results2, perfect := event.Publish("蛤？", nil)
	if !perfect {
		log.Println(event.Error())
	}
	log.Println(results2)
}

func main() {
	testev := event.CreateEvent("TEST", testHandler1, testHandler2, testHandler3)
	event.Subscribe(testev)
	test()
}
