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

import (
	"errors"

	"sangupta.com/velocity/node"
)

type OperatorParser struct {
	/**
	 * The operator we have just scanned, in the same way that {@link #c} is the character we have
	 * just read. If we were not able to scan an operator, this will be {@link Operator#STOP}.
	 */
	currentOperator node.Operator
}

/**
 * Parse a subexpression whose left-hand side is {@code lhs} and where we only consider
 * operators with precedence at least {@code minPrecedence}.
 *
 * @return the parsed subexpression
 */
func (op *OperatorParser) Parse(parser Parser, lhs node.ExpressionNode, minPrecedence uint) node.ExpressionNode {
	for op.currentOperator.GetPrecendence() >= minPrecedence {
		operator := op.currentOperator

		rhs := parser.parseUnaryExpression()

		op.nextOperator(parser)

		for op.currentOperator.GetPrecendence() > operator.GetPrecendence() {
			rhs = op.Parse(parser, rhs, op.currentOperator.GetPrecendence())
		}

		lhs = node.NewBinaryExpressionNode(lhs, operator, rhs)
	}

	return lhs
}

/**
 * Updates {@link #currentOperator} to be an operator read from the input,
 * or {@link Operator#STOP} if there is none.
 */
func (op *OperatorParser) nextOperator(parser Parser) {
	parser.skipSpace()

	possibleOperators := node.GetPossibleOperators(parser.c)
	if len(possibleOperators) == 0 {
		op.currentOperator = node.STOP
		return
	}

	firstChar := parser.c
	parser.next()

	var operator *node.Operator
	for _, op := range possibleOperators {
		if op.GetSymbolLen() == 1 {
			if operator != nil {
				panic(errors.New("operator must be nil"))
			}

			operator = &op
		} else if op.GetSecondRune() == parser.c {
			parser.next()
			operator = &op
		}
	}

	if operator == nil {
		panic(errors.New("Expected possibleOperators, not just " + string(firstChar)))
	}

	op.currentOperator = *operator
}

func NewOperatorParser() *OperatorParser {
	return &OperatorParser{}
}
