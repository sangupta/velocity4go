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
 * A node in the parse tree representing a macro call. If the template contains a definition like
 * {@code #macro (mymacro $x $y) ... #end}, then a call of that macro looks like
 * {@code #mymacro (xvalue yvalue)}. The call is represented by an instance of this class. The
 * definition itself does not appear in the parse tree.
 *
 * <p>Evaluating a macro involves temporarily setting the parameter variables ({@code $x $y} in
 * the example) to thunks representing the argument expressions, evaluating the macro body, and
 * restoring any previous values that the parameter variables had.
 */
type MacroCallNode struct {
	ResourceName string
	LineNumber   uint
	Name         string
	Thunks       []ExpressionNode
	Type         string
}

func (cons *MacroCallNode) GetResourceName() string {
	return cons.ResourceName
}

func (cons *MacroCallNode) GetLineNumber() uint {
	return cons.LineNumber
}

func (node *MacroCallNode) String() string {
	return ""
}

func (node *MacroCallNode) IsWhitespace() bool {
	return false
}

func (node *MacroCallNode) IsHorizontalWhitespace() bool {
	return false
}

func (node *MacroCallNode) MarkDirectiveNode() {

}

func (node *MacroCallNode) Render(context *EvaluationContext, output *strings.Builder) {

}

func NewMacroCallNode(resourceName string, lineNumber uint, name string, thunks []ExpressionNode) *MacroCallNode {
	return &MacroCallNode{
		ResourceName: resourceName,
		LineNumber:   lineNumber,
		Name:         name,
		Thunks:       thunks,
		Type:         "MacroCall",
	}
}
