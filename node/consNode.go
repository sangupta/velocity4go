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

import (
	"errors"
	"strings"
)

type ConsNode struct {
	ResourceName string
	LineNumber   uint
	Children     []Node
	Type         string
}

func (node *ConsNode) String() string {
	return ""
}

func (node *ConsNode) IsWhitespace() bool {
	return false
}

func (node *ConsNode) IsHorizontalWhitespace() bool {
	return false
}

func (node *ConsNode) Render(context *EvaluationContext, output *strings.Builder) {
	for _, localNode := range node.Children {
		if localNode == nil {
			panic(errors.New("nil node added in template parsing phase"))
		}

		localNode.Render(context, output)
	}
}

func NewConsNode(name string, line uint, children []Node) *ConsNode {
	return &ConsNode{
		ResourceName: name,
		LineNumber:   line,
		Children:     children,
	}
}

func EmptyNode(resourceName string, lineNumber uint) *ConsNode {
	return &ConsNode{
		ResourceName: resourceName,
		LineNumber:   lineNumber,
		Children:     make([]Node, 0),
		Type:         "Constant",
	}
}
