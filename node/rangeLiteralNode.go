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

type RangeLiteralNode struct {
	ResourceName string
	LineNumber   uint
	First        ExpressionNode
	Last         ExpressionNode
	Type         string
}

func (cons *RangeLiteralNode) GetResourceName() string {
	return cons.ResourceName
}

func (cons *RangeLiteralNode) GetLineNumber() uint {
	return cons.LineNumber
}

func (node *RangeLiteralNode) String() string {
	return ""
}

func (node *RangeLiteralNode) IsWhitespace() bool {
	return false
}

func (node *RangeLiteralNode) IsHorizontalWhitespace() bool {
	return false
}

func (node *RangeLiteralNode) MarkExpressionNode() {

}

func NewRangeLiteralNode(resourceName string, lineNumber uint, first ExpressionNode, last ExpressionNode) *RangeLiteralNode {
	return &RangeLiteralNode{
		ResourceName: resourceName,
		LineNumber:   lineNumber,
		First:        first,
		Last:         last,
		Type:         "RangeLiteral",
	}
}
