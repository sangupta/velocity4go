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

package parser

import "sangupta.com/velocity4go/node"

/**
 * Denotes end of file as a rune
 */
const EOF = -1

/**
 * Check if node is a type of `node.EofNode`
 */
func isEofNode(localNode node.Node) bool {
	if localNode == nil {
		return false
	}

	_, ok := localNode.(*node.EofNode)
	return ok
}

/**
 * Check if node is a type of `node.EndNode`
 */
func isEndNode(localNode node.Node) bool {
	if localNode == nil {
		return false
	}

	_, ok := localNode.(*node.EndNode)
	return ok
}

/**
 * Check if node is a type of `node.ElseIfNode`
 */
func isElseIfNode(localNode node.Node) bool {
	if localNode == nil {
		return false
	}

	_, ok := localNode.(*node.ElseIfNode)
	return ok
}

/**
 * Check if node is a type of `node.ElseIfNode` or `node.EndNode`
 */
func isElseOrElseIfEndNode(localNode node.Node) bool {
	_, ok := localNode.(*node.ElseNode)
	if ok {
		return true
	}

	_, ok = localNode.(*node.ElseIfNode)
	if ok {
		return true
	}

	_, ok = localNode.(*node.EndNode)
	return ok
}

/**
 * Check if node is a type of `node.SetNode`
 */
func isSetNodeInstance(localNode node.Node) bool {
	if localNode == nil {
		return false
	}

	_, ok := localNode.(*node.SetNode)
	return ok
}

/**
 * Check if node is a type of `node.ReferenceNode`
 */
func isReferenceNode(localNode node.Node) bool {
	if localNode == nil {
		return false
	}

	_, ok := localNode.(node.ReferenceNode)
	return ok
}

/**
 * Check if node is a type of `node.CommentNode`
 */
func isCommentNode(localNode node.Node) bool {
	if localNode == nil {
		return false
	}

	_, ok := localNode.(*node.CommentNode)
	return ok
}

/**
 * Check if node is a type of `node.DirectiveNode`
 */
func isDirectiveNode(localNode node.Node) bool {
	if localNode == nil {
		return false
	}

	_, ok := localNode.(node.DirectiveNode)
	return ok
}

func shouldRemoveLastNodeBeforeSet(nodes []node.Node) bool {
	length := len(nodes)
	if length < 2 {
		return false
	}

	potentialSpaceBeforeSet := nodes[length-1]
	beforeSpace := nodes[length-2]

	if isReferenceNode(beforeSpace) {
		return potentialSpaceBeforeSet.IsHorizontalWhitespace()
	}

	if isCommentNode(beforeSpace) || isDirectiveNode(beforeSpace) {
		return potentialSpaceBeforeSet.IsWhitespace()
	}

	return false
}
