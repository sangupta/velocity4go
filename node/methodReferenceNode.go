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

/**
 * A node in the parse tree representing a method reference, like {@code $list.size()}.
 */
type MethodReferenceNode struct {
	ResourceName string
	LineNumber   uint
	Lhs          ReferenceNode
	Silent       bool
	Id           string
	Args         []ExpressionNode
	Type         string
}

func (cons *MethodReferenceNode) GetResourceName() string {
	return cons.ResourceName
}

func (cons *MethodReferenceNode) GetLineNumber() uint {
	return cons.LineNumber
}

func (node *MethodReferenceNode) String() string {
	return ""
}

func (node *MethodReferenceNode) IsWhitespace() bool {
	return false
}

func (node *MethodReferenceNode) IsHorizontalWhitespace() bool {
	return false
}

func (node *MethodReferenceNode) MarkReferenceNode() {

}

func (node *MethodReferenceNode) MarkExpressionNode() {

}

func NewMethodReferenceNode(lhs ReferenceNode, id string, args []ExpressionNode, silent bool) *MethodReferenceNode {
	return &MethodReferenceNode{
		ResourceName: lhs.GetResourceName(),
		LineNumber:   lhs.GetLineNumber(),
		Lhs:          lhs,
		Silent:       silent,
		Id:           id,
		Args:         args,
		Type:         "MethodReference",
	}
}
