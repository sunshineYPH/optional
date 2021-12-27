package optional

import (
	"fmt"
	"reflect"
)

type Optional struct {
	data      interface{}
	hasValue  bool
	valueType string
}

func New(value interface{}) Optional {
	var optional Optional
	optional.data = value
	if value == nil {
		optional.hasValue = false
	} else {
		optional.hasValue = true
	}

	if v := reflect.TypeOf(value); v != nil {
		optional.valueType = reflect.TypeOf(value).Name()
	}

	return optional
}

// How does go implement static method
// - no static method, use function instead
func Nil() Optional {
	return New(nil)
}

func (op *Optional) Assign(value interface{}) Optional {
	op.data = value
	op.hasValue = true
	op.valueType = reflect.TypeOf(value).String()

	return *op
}

func (op Optional) String() string {
	if op.hasValue {
		str := fmt.Sprintf("%v", op.data)
		return str
	} else {
		return ""
	}
}

func (op Optional) HasValue() bool {
	return op.hasValue
}

func (op Optional) Value() interface{} {
	return op.data
}

func (op Optional) ValueOr(defaultValue interface{}) interface{} {
	if op.HasValue() {
		return op.Value()
	} else {
		return New(defaultValue).Value()
	}
}

func (op Optional) Equal(value interface{}) bool {
	switch value.(type) {
	case Optional:
		return reflect.DeepEqual(op.data, value.(Optional).Value())		
	default:
		return reflect.DeepEqual(op.data, value)
	}	
}
