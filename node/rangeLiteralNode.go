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

type RangeLiteralNode struct {
	ResourceName string
	LineNumber   uint
	First        ExpressionNode
	Last         ExpressionNode
	Type         string
}

func (cons *RangeLiteralNode) GetResourceName() string {
	return cons.ResourceName
}

func (cons *RangeLiteralNode) GetLineNumber() uint {
	return cons.LineNumber
}

func (node *RangeLiteralNode) String() string {
	return ""
}

func (node *RangeLiteralNode) IsWhitespace() bool {
	return false
}

func (node *RangeLiteralNode) IsHorizontalWhitespace() bool {
	return false
}

func (node *RangeLiteralNode) MarkExpressionNode() {

}

func (node *RangeLiteralNode) Render(context *EvaluationContext, output *strings.Builder) {
	renderExpression(context, output, node.Evaluate(context), false)
}

func (node *RangeLiteralNode) IsTrue(context *EvaluationContext) bool {
	return isExpressionTrue(node, context)
}

func (node *RangeLiteralNode) Evaluate(context *EvaluationContext) interface{} {
	// TODO: fix this

	// int from = first.intValue(context);
	//   int to = last.intValue(context);
	//   ImmutableSortedSet<Integer> set =
	//       (from <= to)
	//           ? ContiguousSet.closed(from, to)
	//           : ContiguousSet.closed(to, from).descendingSet();
	//   return new ForwardingSortedSet<Integer>() {
	//     @Override
	//     protected ImmutableSortedSet<Integer> delegate() {
	//       return set;
	//     }

	//     @Override
	//     public String toString() {
	//       // ContiguousSet returns [1..3] whereas Velocity uses [1, 2, 3].
	//       return set.asList().toString();
	//     }
	//   };

	return nil
}

func NewRangeLiteralNode(resourceName string, lineNumber uint, first ExpressionNode, last ExpressionNode) *RangeLiteralNode {
	return &RangeLiteralNode{
		ResourceName: resourceName,
		LineNumber:   lineNumber,
		First:        first,
		Last:         last,
		Type:         "RangeLiteral",
	}
}
