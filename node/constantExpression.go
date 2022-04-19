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

import (
	"unicode"
)

type ConstantExpressionNode struct {
	ResourceName string
	LineNumber   uint
	Value        interface{}
	Type         string
}

func (node *ConstantExpressionNode) GetResourceName() string {
	return node.ResourceName
}

func (node *ConstantExpressionNode) GetLineNumber() uint {
	return node.LineNumber
}

func (node *ConstantExpressionNode) String() string {
	if node.Value == nil {
		return "nil"
	}

	str, ok := node.Value.(string)
	if ok {
		return str
	}

	boo, ok := node.Value.(bool)
	if ok {
		if boo {
			return "true"
		}

		return "false"
	}

	intx, ok := node.Value.(int)
	if ok {
		return string(intx)
	}

	return "unknown"
}

func (node *ConstantExpressionNode) IsWhitespace() bool {
	str, ok := node.Value.(string)
	if !ok {
		return false
	}

	for _, char := range str {
		if !unicode.IsSpace(char) {
			return false
		}
	}

	return true
}

// TODO: fix this
func (node *ConstantExpressionNode) IsHorizontalWhitespace() bool {
	str, ok := node.Value.(string)
	if !ok {
		return false
	}

	for _, char := range str {
		if !unicode.IsSpace(char) {
			return false
		}
	}

	return true
}

func (node *ConstantExpressionNode) MarkExpressionNode() {

}

func NewConstantExpressionNode(resourceName string, lineNumber uint, expression string) *ConstantExpressionNode {
	return &ConstantExpressionNode{
		ResourceName: resourceName,
		LineNumber:   lineNumber,
		Value:        expression,
		Type:         "ConstantExpression",
	}
}
