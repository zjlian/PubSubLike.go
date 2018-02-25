package event

import (
	"errors"
)

type eventBusInterface interface {
	// New 新建实例
	NewEvent(eventType string, handler handlerFunc) *Event
	// HasEvent 检查是否存在事件
	HasEvent(eventType string) bool
	// Publish 发布事件, 发布事件后，会返回该事件所有处理函数的返回值
	Publish(eventType string, payload interface{}) []interface{}
	// Subscribe 订阅事件
	Subscribe(ev *Event)
	// Unsubscribe 取消订阅
	Unsubscribe(eventType string)

	// pushError 添加一个错误
	pushError(err error)
	// Error 字符串方式返回所有存在的错误
	Error() string
}

// Bus 事件总线
type Bus struct {
	// Methos 实例方法
	Methods eventBusInterface
	// 事件对应处理对象储存区
	savers map[string]*Event

	// 事件处理函数错误集合
	errorSet []error
}

// pushError 添加一个错误
func (o *Bus) pushError(err ...error) {
	o.errorSet = append(o.errorSet, err...)
}

// Error 字符串方式返回所有存在的错误
func (o *Bus) Error() string {
	var result = ""
	for _, e := range o.errorSet {
		result += e.Error() + "\n"
	}
	return result
}

// NewEvent 创建一个事件
func (o *Bus) NewEvent(eventType string, handlers ...handlerFunc) *Event {
	ev := &Event{name: eventType}
	for _, h := range handlers {
		h := &eventHandler{Do: h}
		ev.addHandler(h)
	}
	return ev
}

// HasEvent 检查是否存在对应名称的事件
func (o *Bus) HasEvent(eventType string) bool {
	return o.savers[eventType] != nil
}

// Publish 发布事件, 发布事件后，会返回该事件所有处理函数的返回值
func (o *Bus) Publish(eventType string, payload interface{}) ([]interface{}, bool) {
	var (
		resultSet []interface{}
		handler   *eventHandler
		perfect   = true
	)
	if !o.HasEvent(eventType) {
		o.pushError(errors.New("eventBus.go Bus.Publish 触发事件没有订阅者：\"" + eventType + "\""))
		return nil, false
	}
	for _, handler = range o.savers[eventType].handlers {
		tmp, err := handler.Do(payload)
		if err != nil {
			perfect = false
			o.pushError(err)
		}
		resultSet = append(resultSet, tmp)
	}
	return resultSet, perfect
}

// Subscribe 订阅事件
func (o *Bus) Subscribe(ev *Event) {
	if o.HasEvent(ev.name) {
		for _, h := range ev.handlers {
			o.savers[ev.name].addHandler(h)
		}
	} else {
		o.savers[ev.name] = ev
	}
}

// Unsubscribe 取消订阅
func (o *Bus) Unsubscribe(eventType string) {
	if !o.HasEvent(eventType) {
		return
	}

	o.savers[eventType].clear()
}
