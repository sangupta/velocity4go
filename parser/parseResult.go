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

import "sangupta.com/velocity/node"

type ParseResult struct {
	Nodes []node.Node
	stop  node.StopNode
}
