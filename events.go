package event

// CreateEventBus 创建一个事件总线
func CreateEventBus() *Bus {
	return &Bus{savers: make(map[string]*Event)}
}

// GlobalBus 全局总线
var GlobalBus = CreateEventBus()

// CreateEvent 创建一个事件
func CreateEvent(eventType string,
	handlers ...handlerFunc) *Event {
	return GlobalBus.NewEvent(eventType, handlers...)
}

// Publish 发布一个全局事件, 发布事件后，会返回该事件所有处理函数的返回值
func Publish(eventType string, payload interface{}) ([]interface{}, bool) {
	return GlobalBus.Publish(eventType, payload)
}

// Subscribe 订阅一个全局事件
func Subscribe(ev *Event) {
	GlobalBus.Subscribe(ev)
}

// Unsubscribe 取消订阅
func Unsubscribe(eventType string) {
	GlobalBus.Unsubscribe(eventType)
}

// Error 获取
func Error() string {
	return GlobalBus.Error()
}
