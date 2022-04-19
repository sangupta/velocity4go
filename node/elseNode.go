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

type ElseNode struct {
	ResourceName string
	LineNumber   uint
	Type         string
}

func (cons *ElseNode) GetResourceName() string {
	return cons.ResourceName
}

func (cons *ElseNode) GetLineNumber() uint {
	return cons.LineNumber
}

func (node *ElseNode) String() string {
	return ""
}

func (node *ElseNode) IsWhitespace() bool {
	return false
}

func (node *ElseNode) IsHorizontalWhitespace() bool {
	return false
}

func (node *ElseNode) MarkStopNode() {

}

func (node *ElseNode) Render(context *EvaluationContext, output *strings.Builder) {

}

func NewElseNode(name string, line uint) *ElseNode {
	return &ElseNode{
		ResourceName: name,
		LineNumber:   line,
		Type:         "Else",
	}
}
