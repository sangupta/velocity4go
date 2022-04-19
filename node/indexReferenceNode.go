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

/**
 * A node in the parse tree that is an indexing of a reference, like {@code $x[0]} or
 * {@code $x.foo[$i]}. Indexing is array indexing or calling the {@code get} method of a list
 * or a map.
 */
type IndexReferenceNode struct {
	ResourceName string
	LineNumber   uint
	Lhs          ReferenceNode
	Index        ExpressionNode
	Silent       bool
	Type         string
}

func (cons *IndexReferenceNode) GetResourceName() string {
	return cons.ResourceName
}

func (cons *IndexReferenceNode) GetLineNumber() uint {
	return cons.LineNumber
}

func (node *IndexReferenceNode) String() string {
	return ""
}

func (node *IndexReferenceNode) IsWhitespace() bool {
	return false
}

func (node *IndexReferenceNode) IsHorizontalWhitespace() bool {
	return false
}

func (node *IndexReferenceNode) MarkExpressionNode() {

}

func (node *IndexReferenceNode) MarkReferenceNode() {

}

func NewIndexReferenceNode(lhs ReferenceNode, index ExpressionNode, silent bool) *IndexReferenceNode {
	return &IndexReferenceNode{
		ResourceName: lhs.GetResourceName(),
		LineNumber:   lhs.GetLineNumber(),
		Lhs:          lhs,
		Index:        index,
		Silent:       silent,
		Type:         "IndexReference",
	}
}
