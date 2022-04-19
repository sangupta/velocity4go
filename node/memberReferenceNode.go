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
 * A node in the parse tree that is a reference to a property of another reference, like
 * {@code $x.foo} or {@code $x[$i].foo}.
 */
type MemberReferenceNode struct {
	ResourceName string
	LineNumber   uint
	Lhs          ReferenceNode
	Id           string
	Silent       bool
	Type         string
}

func (cons *MemberReferenceNode) GetResourceName() string {
	return cons.ResourceName
}

func (cons *MemberReferenceNode) GetLineNumber() uint {
	return cons.LineNumber
}

func (node *MemberReferenceNode) String() string {
	return ""
}

func (node *MemberReferenceNode) IsWhitespace() bool {
	return false
}

func (node *MemberReferenceNode) IsHorizontalWhitespace() bool {
	return false
}

func (node *MemberReferenceNode) MarkReferenceNode() {

}

func (node *MemberReferenceNode) MarkExpressionNode() {

}

func (node *MemberReferenceNode) Render(context *EvaluationContext, output *strings.Builder) {

}

func NewMemberReferenceNode(lhs ReferenceNode, id string, silent bool) *MemberReferenceNode {
	return &MemberReferenceNode{
		ResourceName: lhs.GetResourceName(),
		LineNumber:   lhs.GetLineNumber(),
		Id:           id,
		Silent:       silent,
		Lhs:          lhs,
		Type:         "MemberReference",
	}
}
