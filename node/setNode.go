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

import "strings"

/**
 * A node in the parse tree representing a {@code #set} construct. Evaluating
 * {@code #set ($x = 23)} will set {@code $x} to the value 23. It does not in itself produce
 * any text in the output.
 *
 * <p>Velocity supports setting values within arrays or collections, with for example
 * {@code $set ($x[$i] = $y)}. That is not currently supported here.
 */
type SetNode struct {
	ResourceName string
	LineNumber   uint
	Variable     string
	Expression   ExpressionNode
	Type         string
}

func (cons *SetNode) GetResourceName() string {
	return cons.ResourceName
}

func (cons *SetNode) GetLineNumber() uint {
	return cons.LineNumber
}

func (node *SetNode) String() string {
	return ""
}

func (node *SetNode) IsWhitespace() bool {
	return false
}

func (node *SetNode) IsHorizontalWhitespace() bool {
	return false
}

func (node *SetNode) MarkDirectiveNode() {

}

func (node *SetNode) Render(context *EvaluationContext, output *strings.Builder) {

}

func NewSetNode(id string, expression ExpressionNode) *SetNode {
	return &SetNode{
		ResourceName: expression.GetResourceName(),
		LineNumber:   expression.GetLineNumber(),
		Variable:     id,
		Expression:   expression,
		Type:         "Set",
	}
}
