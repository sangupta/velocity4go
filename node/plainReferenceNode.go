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

func NewPlainReferenceNode(name string, line uint, id string, silent bool) *PlainReferenceNode {
	return &PlainReferenceNode{
		ResourceName: name,
		LineNumber:   line,
		Id:           id,
		Silent:       silent,
		Type:         "PlainReference",
	}
}
