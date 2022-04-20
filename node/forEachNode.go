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
	"reflect"
	"strings"
)

/**
 * A node in the parse tree representing a {@code #foreach} construct. While evaluating
 * {@code #foreach ($x in $things)}, {$code $x} will be set to each element of {@code $things} in
 * turn. Once the loop completes, {@code $x} will go back to whatever value it had before, which
 * might be undefined. During loop execution, the variable {@code $foreach} is also defined.
 * Velocity defines a number of properties in this variable, but here we only support
 * {@code $foreach.hasNext} and {@code $foreach.index}.
 */
type ForEachNode struct {
	ResourceName string
	LineNumber   uint
	Variable     string
	Collection   ExpressionNode
	Body         Node
	Type         string
}

func (node *ForEachNode) String() string {
	return ""
}

func (node *ForEachNode) IsWhitespace() bool {
	return false
}

func (node *ForEachNode) IsHorizontalWhitespace() bool {
	return false
}

func (node *ForEachNode) MarkDirectiveNode() {

}

func (node *ForEachNode) Render(context *EvaluationContext, output *strings.Builder) {
	collectionValue := node.Collection.Evaluate(context)

	if collectionValue == nil {
		return
	}

	typeOf := reflect.TypeOf(collectionValue)
	kind := typeOf.Kind()

	switch kind {
	case reflect.Slice:
	case reflect.Array:
	case reflect.Map:
		break

	default:
		panic(errors.New("Value is not iterable"))
	}

}

func NewForEachNode(resourceName string, lineNumber uint, id string, collection ExpressionNode, body Node) *ForEachNode {
	return &ForEachNode{
		ResourceName: resourceName,
		LineNumber:   lineNumber,
		Variable:     id,
		Collection:   collection,
		Body:         body,
		Type:         "ForEach",
	}
}
