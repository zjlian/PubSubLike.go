package event

// HandlerFunc 事件处理函数结构
type eventHandler struct {
	Do handlerFunc
}
type handlerFunc func(interface{}) (interface{}, error)

// event结构体的实例方法
type eventInterface interface {
	// addHandler 添加一个处理函数
	addHandler(handler *eventHandler)
	// removeListener 移除指定的处理函数
	removeHandler(handler *eventHandler)
	// clear 清空该事件的处理函数
	clear()
}

// Event 事件结构体
type Event struct {
	// Method 事件结构体的实例方法
	methods eventInterface
	// 事件名称
	name string
	// 事件处理函数集合
	handlers []*eventHandler
}

// 添加一个处理函数
func (o *Event) addHandler(handler *eventHandler) {
	// 检查是否已经存在相同的处理函数
	for _, h := range o.handlers {
		if h == handler {
			return
		}
	}
	o.handlers = append(o.handlers, handler)
}

// 移除指定的处理函数
func (o *Event) removeHandler(handler *eventHandler) {
	for i, h := range o.handlers {
		if h == handler {
			o.handlers = append(o.handlers[:i], o.handlers[i+1:]...)
		}
	}
}

// 清空该事件的处理函数
func (o *Event) clear() {
	o.handlers = o.handlers[:0:0]
}
