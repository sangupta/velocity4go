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

/**
 * A node in the parse tree representing an {@code #if} construct. All instances of this class
 * have a <i>true</i> subtree and a <i>false</i> subtree. For a plain {@code #if (cond) body
 * #end}, the false subtree will be empty. For {@code #if (cond1) body1 #elseif (cond2) body2
 * #else body3 #end}, the false subtree will contain a nested {@code IfNode}, as if {@code #else
 * #if} had been used instead of {@code #elseif}.
 */
type IfNode struct {
	ResourceName string
	LineNumber   uint
	Condition    ExpressionNode
	TruePart     Node
	FalsePart    Node
	Type         string
}

func (cons *IfNode) GetResourceName() string {
	return cons.ResourceName
}

func (cons *IfNode) GetLineNumber() uint {
	return cons.LineNumber
}

func (node *IfNode) String() string {
	return ""
}

func (node *IfNode) IsWhitespace() bool {
	return false
}

func (node *IfNode) IsHorizontalWhitespace() bool {
	return false
}

func (node *IfNode) MarkDirectiveNode() {

}

func NewIfNode(resourceName string, lineNumber uint, condition ExpressionNode, truePart Node, falsePart Node) *IfNode {
	return &IfNode{
		ResourceName: resourceName,
		LineNumber:   lineNumber,
		Condition:    condition,
		TruePart:     truePart,
		FalsePart:    falsePart,
		Type:         "If",
	}
}
