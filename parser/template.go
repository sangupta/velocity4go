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

package parser

import (
	"strings"

	"sangupta.com/velocity4go/node"
)

type Template struct {
	Root   node.Node
	Macros map[string]Macro
	Type   string
}

/**
 * Evaluate this template against the given set of
 * variable data.
 */
func (template *Template) Evaluate(variables map[string]interface{}) string {
	builder := strings.Builder{}
	context := node.EvaluationContext{
		Variables: variables,
	}

	template.Root.Render(&context, &builder)

	return builder.String()
}
