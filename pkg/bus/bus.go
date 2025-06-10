package bus

import (
	"context"
	"fmt"
	"reflect"
	"sync"
)

type Bus struct {
	handlers map[reflect.Type]interface{}
	mu sync.RWMutex
}

func NewBus() *Bus {
	return &Bus{
		handlers: make(map[reflect.Type]interface{}),
	}
}

func (b *Bus)RegisterHandler(handler interface{}, msg interface{}){
	b.mu.Lock()
	defer b.mu.Unlock()

	msgType := reflect.TypeOf(msg)
	b.handlers[msgType] = handler
}

func (b *Bus)Dispatch(ctx context.Context, msg interface{}) (interface{}, error){
	b.mu.Lock()
	defer b.mu.Unlock()
	
	msgType := reflect.TypeOf(msg)
	
	handler, ok := b.handlers[msgType]
	if !ok {
		return nil, fmt.Errorf("any handler find for the type: %s", msgType)
	}

	handlerValue := reflect.ValueOf(handler)

	method := handlerValue.MethodByName("Handle")
	if !method.IsValid() {
		return nil, fmt.Errorf("handler for %s doesn't have the method 'handle'", msgType)
	}

	result := method.Call([]reflect.Value{reflect.ValueOf(ctx), reflect.ValueOf(msg)})
	if len(result) != 2 {
		return nil, fmt.Errorf("handle method has more than 2 values")
	}

	if !result[1].IsNil() {
		return nil, result[1].Interface().(error)
	}

	return result[0].Interface(), nil
}