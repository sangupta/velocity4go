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

type ConsNode struct {
	ResourceName string
	LineNumber   uint
	Children     []Node
	Type         string
}

func (cons *ConsNode) GetResourceName() string {
	return cons.ResourceName
}

func (cons *ConsNode) GetLineNumber() uint {
	return cons.LineNumber
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
