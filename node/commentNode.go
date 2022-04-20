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

type CommentNode struct {
	ResourceName string
	LineNumber   uint
	Type         string
}

func (node *CommentNode) String() string {
	return ""
}

func (node *CommentNode) IsWhitespace() bool {
	return false
}

func (node *CommentNode) IsHorizontalWhitespace() bool {
	return false
}

func (node *CommentNode) Render(context *EvaluationContext, output *strings.Builder) {

}

func NewCommentNode(name string, line uint) *CommentNode {
	return &CommentNode{
		ResourceName: name,
		LineNumber:   line,
		Type:         "Comment",
	}
}
