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

type ElseIfNode struct {
	ResourceName string
	LineNumber   uint
	Type         string
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

func (node *ElseIfNode) Render(context *EvaluationContext, output *strings.Builder) {

}

func NewElseIfNode(name string, line uint) *ElseIfNode {
	return &ElseIfNode{
		ResourceName: name,
		LineNumber:   line,
		Type:         "ElseIf",
	}
}
