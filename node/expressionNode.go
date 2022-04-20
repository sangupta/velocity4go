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

	"sangupta.com/velocity4go/utils"
)

/**
 * Marker interface to borrow inheritance
 */
type ExpressionNode interface {
	Node
	GetResourceName() string
	GetLineNumber() uint
	IsTrue(context *EvaluationContext) bool
	Evaluate(context *EvaluationContext) interface{}
	MarkExpressionNode()
}

func renderExpression(context *EvaluationContext, output *strings.Builder, rendered interface{}, silent bool) {
	if rendered == nil {
		if silent { // $!foo for example
			return
		}

		// throw evaluationException("Null value for " + this);
		panic(errors.New("Null value for rendered"))
	}

	output.WriteString(utils.AsString(rendered))
}

/**
 * True if evaluating this expression yields a value that is considered true by Velocity's
 * <a href="http://velocity.apache.org/engine/releases/velocity-1.7/user-guide.html#Conditionals">
 * rules</a>.  A value is false if it is null or equal to Boolean.FALSE.
 * Every other value is true.
 *
 * <p>Note that the text at the similar link
 * <a href="http://velocity.apache.org/engine/devel/user-guide.html#Conditionals">here</a>
 * states that empty collections and empty strings are also considered false, but that is not
 * true.
 */
func isExpressionTrue(node ExpressionNode, context *EvaluationContext) bool {
	value := node.Evaluate(context)
	boolValue, ok := value.(bool)
	if ok {
		return boolValue
	}

	return value != nil
}

/**
 * True if this is a defined value and it evaluates to true. This is the same as {@link #isTrue}
 * except that it is allowed for this to be undefined variable, in which it evaluates to false.
 * The method is overridden for plain references so that undefined is the same as false.
 * The reason is to support Velocity's idiom {@code #if ($var)}, where it is not an error
 * if {@code $var} is undefined.
 */
func isExpressionDefinedAndTrue(node ExpressionNode, context *EvaluationContext) bool {
	return node.IsTrue(context)
}

func expressionIntValue(node ExpressionNode, context *EvaluationContext) utils.Number {
	value := node.Evaluate(context)
	if value == nil {
		return utils.NilNumber()
	}

	number, ok := value.(utils.Number)
	if ok {
		return number
	}

	panic(errors.New("Arithmetic is only available on integers"))
}
