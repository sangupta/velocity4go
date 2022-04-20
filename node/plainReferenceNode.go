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
	"errors"
	"strings"
)

/**
 * A node in the parse tree that is a plain reference such as {@code $x}. This node may appear
 * inside a more complex reference like {@code $x.foo}.
 */
type PlainReferenceNode struct {
	ResourceName string
	LineNumber   uint
	Id           string
	Silent       bool
	Type         string
}

func (cons *PlainReferenceNode) GetResourceName() string {
	return cons.ResourceName
}

func (cons *PlainReferenceNode) GetLineNumber() uint {
	return cons.LineNumber
}

func (node *PlainReferenceNode) String() string {
	return ""
}

func (node *PlainReferenceNode) IsWhitespace() bool {
	return false
}

func (node *PlainReferenceNode) IsHorizontalWhitespace() bool {
	return false
}

func (node *PlainReferenceNode) MarkReferenceNode() {

}

func (node *PlainReferenceNode) MarkExpressionNode() {

}

func (node *PlainReferenceNode) IsTrue(context *EvaluationContext) bool {
	return isExpressionTrue(node, context)
}

func (node *PlainReferenceNode) Render(context *EvaluationContext, output *strings.Builder) {
	renderExpression(context, output, node.Evaluate(context), node.Silent)
}

func (node *PlainReferenceNode) Evaluate(context *EvaluationContext) interface{} {
	if !context.IsVarDefined(node.Id) {
		panic(errors.New("undefined reference: " + node.Id))
	}

	return context.GetVar(node.Id)
}

func NewPlainReferenceNode(name string, line uint, id string, silent bool) *PlainReferenceNode {
	return &PlainReferenceNode{
		ResourceName: name,
		LineNumber:   line,
		Id:           id,
		Silent:       silent,
		Type:         "PlainReference",
	}
}
