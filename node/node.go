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

/**
 * Marker interface to show inheritance.
 */
type Node interface {
	String() string
	IsWhitespace() bool
	IsHorizontalWhitespace() bool
	Render(context *EvaluationContext, output *strings.Builder)
}
