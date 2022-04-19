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
