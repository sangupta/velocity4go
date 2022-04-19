/**
 * go-velocity: Velocity template engine for Go
 * https://sangupta.com/projects/go-velocity
 *
 * MIT License.
 * Copyright (c) 2022, Sandeep Gupta.
 *
 * Use of this source code is governed by a MIT style license
 * that can be found in LICENSE file in the code repository.
 */

package node

type NotExpressionNode struct {
	ResourceName string
	LineNumber   uint
	Expression   ExpressionNode
	Type         string
}

func (cons *NotExpressionNode) GetResourceName() string {
	return cons.ResourceName
}

func (cons *NotExpressionNode) GetLineNumber() uint {
	return cons.LineNumber
}

func (node *NotExpressionNode) String() string {
	return ""
}

func (node *NotExpressionNode) IsWhitespace() bool {
	return false
}

func (node *NotExpressionNode) IsHorizontalWhitespace() bool {
	return false
}

func (node *NotExpressionNode) MarkExpressionNode() {

}

func NewNotExpressionNode(expr ExpressionNode) *NotExpressionNode {
	return &NotExpressionNode{
		ResourceName: expr.GetResourceName(),
		LineNumber:   expr.GetLineNumber(),
		Expression:   expr,
		Type:         "NotExpression",
	}
}
