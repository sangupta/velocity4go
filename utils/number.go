/**
 * velocity4go: Velocity template engine for Go
 * https://sangupta.com/projects/velocity4go
 *
 * MIT License.
 * Copyright (c) 2022, Sandeep Gupta.
 *
 * Use of this source code is governed by a MIT style license
 * that can be found in LICENSE file in the code repository.
 */

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
