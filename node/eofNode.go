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

type EofNode struct {
	ResourceName string
	LineNumber   uint
	Type         string
}

func (cons *EofNode) GetResourceName() string {
	return cons.ResourceName
}

func (cons *EofNode) GetLineNumber() uint {
	return cons.LineNumber
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

func NewEofNode(resourceName string, lineNumber uint) *EofNode {
	return &EofNode{
		ResourceName: resourceName,
		LineNumber:   lineNumber,
		Type:         "EOF",
	}
}
