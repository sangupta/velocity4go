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

type EofNode struct {
	ResourceName string
	LineNumber   uint
	Type         string
}

func (node *EofNode) String() string {
	return ""
}

func (node *EofNode) IsWhitespace() bool {
	return false
}

func (node *EofNode) IsHorizontalWhitespace() bool {
	return false
}

func (node *EofNode) MarkStopNode() {

}

func (node *EofNode) Render(context *EvaluationContext, output *strings.Builder) {

}

func NewEofNode(resourceName string, lineNumber uint) *EofNode {
	return &EofNode{
		ResourceName: resourceName,
		LineNumber:   lineNumber,
		Type:         "EOF",
	}
}
