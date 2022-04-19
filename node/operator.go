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

type Operator struct {
	Symbol     string
	Precedence uint
}

func (op *Operator) GetSymbol() string {
	return op.Symbol
}

func (op *Operator) GetSymbolLen() int {
	return len(op.Symbol)
}

func (op *Operator) GetPrecendence() uint {
	return op.Precedence
}

func (op *Operator) String() string {
	return op.Symbol
}

func (op *Operator) GetFirstRune() rune {
	return []rune(op.Symbol)[0]
}

func (op *Operator) GetSecondRune() rune {
	return []rune(op.Symbol)[1]
}

/**
 * True if this is an inequality operator, one of {@code < > <= >=}.
 */
func (op *Operator) IsInequality() bool {
	return op.Precedence == 4
}

var (
	/**
	 * A dummy operator with low precedence. When parsing subexpressions, we always stop when we
	 * reach an operator of lower precedence than the "current precedence". For example, when
	 * parsing {@code 1 + 2 * 3 + 4}, we'll stop parsing the subexpression {@code * 3 + 4} when
	 * we reach the {@code +} because it has lower precedence than {@code *}. This dummy operator,
	 * then, behaves like {@code +} when the minimum precedence is {@code *}. We also return it
	 * if we're looking for an operator and don't find one. If this operator is {@code ⊙}, it's as
	 * if our expressions are bracketed with it, like {@code ⊙ 1 + 2 * 3 + 4 ⊙}.
	 */
	STOP = Operator{"", 0}

	// If a one-character operator is a prefix of a two-character operator, like < and <=, then
	// the one-character operator must come first.
	OR               = Operator{"||", 1}
	AND              = Operator{"&&", 2}
	EQUAL            = Operator{"==", 3}
	NOT_EQUAL        = Operator{"!=", 3}
	LESS             = Operator{"<", 4}
	LESS_OR_EQUAL    = Operator{"<=", 4}
	GREATER          = Operator{">", 4}
	GREATER_OR_EQUAL = Operator{">=", 4}
	PLUS             = Operator{"+", 5}
	MINUS            = Operator{"-", 5}
	TIMES            = Operator{"*", 6}
	DIVIDE           = Operator{"/", 6}
	REMAINDER        = Operator{"%", 6}
)

var ALL_OPERATORS = []Operator{OR, AND, EQUAL, NOT_EQUAL, LESS, LESS_OR_EQUAL, GREATER, GREATER_OR_EQUAL, PLUS, MINUS, TIMES, DIVIDE, REMAINDER}

func GetPossibleOperators(char rune) []Operator {
	ops := make([]Operator, 0, 5)

	for _, op := range ALL_OPERATORS {
		firstRune := op.GetFirstRune()
		if firstRune == char {
			ops = append(ops, op)
		}
	}

	return ops
}
