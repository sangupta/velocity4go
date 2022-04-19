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

type ElseIfNode struct {
	ResourceName string
	LineNumber   uint
	Type         string
}

func (cons *ElseIfNode) GetResourceName() string {
	return cons.ResourceName
}

func (cons *ElseIfNode) GetLineNumber() uint {
	return cons.LineNumber
}

func (node *ElseIfNode) String() string {
	return ""
}

func (node *ElseIfNode) IsWhitespace() bool {
	return false
}

func (node *ElseIfNode) IsHorizontalWhitespace() bool {
	return false
}

func (node *ElseIfNode) MarkStopNode() {

}

func NewElseIfNode(name string, line uint) *ElseIfNode {
	return &ElseIfNode{
		ResourceName: name,
		LineNumber:   line,
		Type:         "ElseIf",
	}
}
