package utils

import (
	"errors"
	"reflect"
)

type Number struct {
	value    interface{}
	dataType reflect.Type
}

func (num *Number) IsNil() bool {
	return num.dataType == nil
}

func NilNumber() Number {
	num := Number{}
	return num
}

func NewNumber(value interface{}) Number {
	num := Number{}
	num.SetValue(value)

	return num
}

func (num *Number) HasValue() bool {
	return num.dataType != nil
}

func (num *Number) SetNil() {
	num.value = nil
	num.dataType = nil
}

func (num *Number) SetValue(value interface{}) {
	num.value = value
	num.dataType = reflect.TypeOf(value)
}

func (num *Number) asInt() (int, error) {
	i, ok := num.value.(int)
	if ok {
		return i, nil
	}

	return 0, errors.New("Cannot convert value to int")
}
