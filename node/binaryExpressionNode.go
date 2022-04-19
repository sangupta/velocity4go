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

type BinaryExpressionNode struct {
	ResourceName string
	LineNumber   uint
	Lhs          ExpressionNode
	Rhs          ExpressionNode
	Operator     Operator
	Type         string
}

func (cons *BinaryExpressionNode) GetResourceName() string {
	return cons.ResourceName
}

func (cons *BinaryExpressionNode) GetLineNumber() uint {
	return cons.LineNumber
}

func (node *BinaryExpressionNode) String() string {
	return ""
}

func (node *BinaryExpressionNode) IsWhitespace() bool {
	return false
}

func (node *BinaryExpressionNode) IsHorizontalWhitespace() bool {
	return false
}

func (node *BinaryExpressionNode) MarkExpressionNode() {

}

func NewBinaryExpressionNode(lhs ExpressionNode, op Operator, rhs ExpressionNode) *BinaryExpressionNode {
	return &BinaryExpressionNode{
		ResourceName: lhs.GetResourceName(),
		LineNumber:   lhs.GetLineNumber(),
		Lhs:          lhs,
		Rhs:          rhs,
		Operator:     op,
		Type:         "BinaryExpression",
	}
}
