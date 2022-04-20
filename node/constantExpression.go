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
	"strings"
	"unicode"

	"sangupta.com/velocity4go/utils"
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

func (node *ConstantExpressionNode) IsTrue(context *EvaluationContext) bool {
	return isExpressionTrue(node, context)
}

func (node *ConstantExpressionNode) Render(context *EvaluationContext, output *strings.Builder) {
	renderExpression(context, output, node.Evaluate(context), false)
}

func (node *ConstantExpressionNode) Evaluate(context *EvaluationContext) interface{} {
	return node.String()
}

func (node *ConstantExpressionNode) String() string {
	return utils.AsString(node.Value)
}

func (node *ConstantExpressionNode) IsSilent() bool {
	return false
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
