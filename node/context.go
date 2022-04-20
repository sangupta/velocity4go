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

package node

import "errors"

type EvaluationContext struct {
	Variables map[string]interface{}
}

func (context *EvaluationContext) IsVarDefined(id string) bool {
	_, ok := context.Variables[id]
	return ok
}

func (context *EvaluationContext) GetVar(id string) interface{} {
	value, ok := context.Variables[id]
	if ok {
		return value
	}

	panic(errors.New("cannot get variable which is not defined: " + id))
}

/**
 * Sets a variable in the context and returns an `undo` function to
 * do away with the side effects when needed.
 */
func (context *EvaluationContext) SetVar(id string, value interface{}) func() {
	var undo func()

	oldValue, isDefined := context.Variables[id]
	if isDefined {
		undo = func() {
			context.SetVar(id, oldValue)
		}
	} else {
		undo = func() {
			context.Remove(id)
		}
	}

	context.Variables[id] = value
	return undo
}

func (context *EvaluationContext) Remove(id string) {
	delete(context.Variables, id)
}
