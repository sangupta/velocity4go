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

type ListLiteralNode struct {
	ResourceName string
	LineNumber   uint
	Elements     []ExpressionNode
	Type         string
}

func (cons *ListLiteralNode) GetResourceName() string {
	return cons.ResourceName
}

func (cons *ListLiteralNode) GetLineNumber() uint {
	return cons.LineNumber
}

func (node *ListLiteralNode) String() string {
	return ""
}

func (node *ListLiteralNode) IsWhitespace() bool {
	return false
}

func (node *ListLiteralNode) IsHorizontalWhitespace() bool {
	return false
}

func (node *ListLiteralNode) MarkExpressionNode() {

}

func (node *ListLiteralNode) Render(context *EvaluationContext, output *strings.Builder) {

}

func NewListLiteralNode(name string, line uint, elements []ExpressionNode) *ListLiteralNode {
	return &ListLiteralNode{
		ResourceName: name,
		LineNumber:   line,
		Elements:     elements,
		Type:         "ListLiteral",
	}
}
