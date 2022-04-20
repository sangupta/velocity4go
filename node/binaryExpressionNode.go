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

type BinaryExpressionNode struct {
	ResourceName string
	LineNumber   uint
	Lhs          ExpressionNode
	Rhs          ExpressionNode
	Operator     Operator
	Type         string
}

func (cons *BinaryExpressionNode) GetResourceName() string {
	return cons.ResourceName
}

func (cons *BinaryExpressionNode) GetLineNumber() uint {
	return cons.LineNumber
}

func (node *BinaryExpressionNode) String() string {
	return ""
}

func (node *BinaryExpressionNode) IsWhitespace() bool {
	return false
}

func (node *BinaryExpressionNode) IsHorizontalWhitespace() bool {
	return false
}

func (node *BinaryExpressionNode) MarkExpressionNode() {

}

func (node *BinaryExpressionNode) Render(context *EvaluationContext, output *strings.Builder) {

}

func (node *BinaryExpressionNode) IsSilent() bool {
	return false
}

func (node *BinaryExpressionNode) IsTrue(context *EvaluationContext) bool {
	return isExpressionTrue(node, context)
}

func (node *BinaryExpressionNode) Evaluate(context *EvaluationContext) interface{} {
	switch node.Operator {
	case OR:
		return node.Lhs.IsTrue(context) || node.Rhs.IsTrue(context)

	case AND:
		return node.Lhs.IsTrue(context) && node.Rhs.IsTrue(context)

	case EQUAL:
		return node.equal(context)

	case NOT_EQUAL:
		return !node.equal(context)
	}

	return nil
}

/**
 * Returns true if {@code lhs} and {@code rhs} are equal according to Velocity.
 *
 * <p>Velocity's <a
 * href="http://velocity.apache.org/engine/releases/velocity-1.7/vtl-reference-guide.html#aifelseifelse_-_Output_conditional_on_truth_of_statements">definition
 * of equality</a> differs depending on whether the objects being compared are of the same
 * class. If so, equality comes from {@code Object.equals} as you would expect.  But if they
 * are not of the same class, they are considered equal if their {@code toString()} values are
 * equal. This means that integer 123 equals long 123L and also string {@code "123"}.  It also
 * means that equality isn't always transitive. For example, two StringBuilder objects each
 * containing {@code "123"} will not compare equal, even though the string {@code "123"}
 * compares equal to each of them.
 */
func (node *BinaryExpressionNode) equal(context *EvaluationContext) bool {
	leftValue := node.Lhs.Evaluate(context)
	rightValue := node.Rhs.Evaluate(context)

	// todo fix this
	if leftValue == nil || rightValue == nil {
		return false
	}

	if leftValue == rightValue {
		return true
	}

	// todo fix this
	return false
}

func NewBinaryExpressionNode(lhs ExpressionNode, op Operator, rhs ExpressionNode) *BinaryExpressionNode {
	return &BinaryExpressionNode{
		ResourceName: lhs.GetResourceName(),
		LineNumber:   lhs.GetLineNumber(),
		Lhs:          lhs,
		Rhs:          rhs,
		Operator:     op,
		Type:         "BinaryExpression",
	}
}
