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

type EndNode struct {
	ResourceName string
	LineNumber   uint
	Type         string
}

func (node *EndNode) String() string {
	return ""
}

func (node *EndNode) IsWhitespace() bool {
	return false
}

func (node *EndNode) IsHorizontalWhitespace() bool {
	return false
}

func (node *EndNode) MarkStopNode() {

}

func (node *EndNode) Render(context *EvaluationContext, output *strings.Builder) {

}

func NewEndNode(name string, line uint) *EndNode {
	return &EndNode{
		ResourceName: name,
		LineNumber:   line,
		Type:         "End",
	}
}
