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
	"fmt"
	"strings"
	"unicode"

	node "sangupta.com/velocity4go/node"
	utils "sangupta.com/velocity4go/utils"
)

const EOF = -1

func isEofNode(localNode node.Node) bool {
	_, ok := localNode.(*node.EofNode)
	return ok
}

func isEndNode(localNode node.Node) bool {
	_, ok := localNode.(*node.EndNode)
	return ok
}

func isElseIfNode(localNode node.Node) bool {
	_, ok := localNode.(*node.ElseIfNode)
	return ok
}

func isElseOrElseIfEndNode(localNode node.Node) bool {
	_, ok := localNode.(*node.ElseNode)
	if ok {
		return true
	}

	_, ok = localNode.(*node.ElseIfNode)
	if ok {
		return true
	}

	_, ok = localNode.(*node.EndNode)
	return ok
}

func isSetNodeInstance(localNode node.Node) bool {
	_, ok := localNode.(*node.SetNode)
	return ok
}

func isReferenceNode(localNode node.Node) bool {
	_, ok := localNode.(node.ReferenceNode)
	return ok
}

func isCommentNode(localNode node.Node) bool {
	_, ok := localNode.(*node.CommentNode)
	return ok
}

func isDirectiveNode(localNode node.Node) bool {
	_, ok := localNode.(node.DirectiveNode)
	return ok
}

func shouldRemoveLastNodeBeforeSet(nodes []node.Node) bool {
	length := len(nodes)
	if length < 2 {
		return false
	}

	potentialSpaceBeforeSet := nodes[length-1]
	beforeSpace := nodes[length-2]

	if isReferenceNode(beforeSpace) {
		return potentialSpaceBeforeSet.IsHorizontalWhitespace()
	}

	if isCommentNode(beforeSpace) || isDirectiveNode(beforeSpace) {
		return potentialSpaceBeforeSet.IsWhitespace()
	}

	return false
}

// ----------------
// Parser functions
// ----------------

type Parser struct {
	Chars        []rune
	c            rune
	ResourceName string
	pointer      uint `default:"0"`
	pushback     int
	macros       map[string]Macro
	line         uint
}

func (parser *Parser) Parse() (Template, error) {
	parser.c = parser.Chars[0]

	parseResult := parser.parseToStop(isEofNode, "outside any construct")
	root := node.NewConsNode(parser.ResourceName, parser.lineNumber(), parseResult.Nodes)

	return Template{
		Root:   root,
		Macros: parser.macros,
		Type:   "Template",
	}, nil
}

/**
 * Gets the next character from the reader and assigns it to {@code c}. If there are no more
 * characters, sets {@code c} to {@link #EOF} if it is not already.
 */
func (parser *Parser) next() {
	if parser.c != EOF {
		if parser.pushback >= 0 {
			parser.pointer++
			if parser.pointer >= uint(len(parser.Chars)) {
				parser.c = EOF
			} else {
				parser.c = parser.Chars[parser.pointer]
				if parser.c == '\n' {
					parser.line++
				}
			}
		} else {
			parser.c = rune(parser.pushback)
			parser.pushback = -1
		}
	}
}

/**
 * If {@code c} is a space character, keeps reading until {@code c} is a non-space character or
 * there are no more characters.
 */
func (parser *Parser) skipSpace() {
	for unicode.IsSpace(parser.c) {
		parser.next()
	}
}

/**
 * Saves the current character {@code c} to be read again, and sets {@code c} to the given
 * {@code c1}. Suppose the text contains {@code xy} and we have just read {@code y}.
 * So {@code c == 'y'}. Now if we execute {@code pushback('x')}, we will have
 * {@code c == 'x'} and the next call to {@link #next()} will set {@code c == 'y'}. Subsequent
 * calls to {@code next()} will continue reading from {@link #reader}. So the pushback
 * essentially puts us back in the state we were in before we read {@code y}.
 */
func (parser *Parser) doPushback(char rune) {
	parser.pushback = int(parser.c)
	parser.c = char
}

/**
 * Gets the next character from the reader, and if it is a space character, keeps reading until
 * a non-space character is found.
 */
func (parser *Parser) nextNonSpace() {
	parser.next()
	parser.skipSpace()
}

/**
 * Skips any space in the reader, and then throws an exception if the first non-space character
 * found is not the expected one. Sets {@code c} to the first character after that expected one.
 */
func (parser *Parser) expect(expected rune) {
	parser.skipSpace()
	if parser.c == expected {
		parser.next()
		return
	}

	panic(utils.ParseException("Expected: " + string(expected) + ", found: " + string(parser.c)))
}

/**
 * Parse until reaching a {@code StopNode}. The {@code StopNode} must have one of the classes in
 * {@code stopClasses}. This method is called recursively to parse nested constructs. At the
 * top level, we expect the parse to end when it reaches {@code EofNode}. In a {@code #foreach},
 * for example, we expect the parse to end when it reaches the matching {@code #end}. In an
 * {@code #if}, the parse can end with {@code #end}, {@code #else}, or {@code #elseif}. And then
 * after {@code #else} or {@code #elseif} we will call this method again to parse the next part.
 *
 * @return the nodes that were parsed, plus the {@code StopNode} that caused parsing to stop.
 */
func (parser *Parser) parseToStop(stopClasses func(node node.Node) bool, contextDescription string) ParseResult {
	nodes := make([]node.Node, 0)
	var localNode node.Node
	for true {
		localNode = parser.parseNode()

		if stopClasses(localNode) {
			break
		}

		isSetNode := isSetNodeInstance(localNode)
		if isSetNode && shouldRemoveLastNodeBeforeSet(nodes) {
			nodes[len(nodes)-1] = localNode
		} else {
			nodes = append(nodes, localNode)
		}
	}

	stopNode, ok := localNode.(node.StopNode)
	if !ok {
		panic(errors.New("Found " + localNode.String() + " " + contextDescription))
	}

	return ParseResult{
		Nodes: nodes,
		stop:  stopNode,
	}
}

/**
 * Skip the current character if it is a newline, then parse until reaching a {@code StopNode}.
 * This is used after directives like {@code #if}, where a newline is ignored after the final
 * {@code )} in {@code #if (condition)}.
 */
func (parser *Parser) skipNewLineAndParseToStop(stopClasses func(node node.Node) bool, contextDescription string) ParseResult {
	if parser.c == '\n' {
		parser.next()
	}
	return parser.parseToStop(stopClasses, contextDescription)
}

/**
 * Parses a single node from the reader.
 */
func (parser *Parser) parseNode() node.Node {
	if parser.c == '#' {
		parser.next()

		switch parser.c {
		case '#':
			return parser.parseLineComment()

		case '*':
			return parser.parseBlockComment()

		case '[':
			return parser.parseHashSquare()

		case '{':
			return parser.parseDirective()

		default:
			if utils.IsAsciiLetter(parser.c) {
				return parser.parseDirective()
			}

			// For consistency with Velocity, we treat # not followed by a letter or one of the
			// characters above as a plain character, and we treat #$foo as a literal # followed by
			// the reference $foo.
			return parser.parsePlainTextRune('#')
		}
	}

	if parser.c == EOF {
		return node.NewEofNode(parser.ResourceName, parser.lineNumber())
	}

	return parser.parseNonDirective()
}

/**
 * Parses a line comment, which is {@code ##} followed by any number of characters
 * up to and including the next newline.
 */
func (parser *Parser) parseLineComment() node.Node {
	lineNumber := parser.lineNumber()
	for parser.c != '\n' && parser.c != EOF {
		parser.next()
	}
	parser.next()
	return node.NewCommentNode(parser.ResourceName, lineNumber)
}

func (parser *Parser) parseBlockComment() node.Node {
	utils.AssertRune(parser.c, '*')
	startLine := parser.lineNumber()
	var lastC rune
	lastC = 0
	parser.next()

	// Consistently with Velocity, we do not make it an error if a #* comment is not closed.
	for !(lastC == '*' && parser.c == '#') && parser.c != EOF {
		lastC = parser.c
		parser.next()
	}
	parser.next() // this may read EOF twice, which works

	return node.NewCommentNode(parser.ResourceName, startLine)
}

func (parser *Parser) lineNumber() uint {
	return parser.line
}

func (parser *Parser) parseHashSquare() node.Node {
	// We've just seen #[ which might be the start of a #[[quoted block]]#. If the next character
	// is not another [ then it's not a quoted block, but it *is* a literal #[ followed by whatever
	// that next character is.
	utils.AssertRune(parser.c, '[')

	parser.next()

	if parser.c != '[' {
		return parser.parsePlainText("#[")
	}

	startLine := parser.lineNumber()
	parser.next()

	var sb strings.Builder
	for true {
		if parser.c == EOF {
			panic(utils.ParseException("Unterminated #[[ - did not see matching ]]#" + " in " + parser.ResourceName + " at " + fmt.Sprint(startLine)))
		}

		if parser.c == '#' {
			// This might be the last character of ]]# or it might just be a random #.
			len := sb.Len()

			current := sb.String()
			if len > 1 && current[len-1] == ']' && current[len-2] == ']' {
				parser.next()
				break
			}
		}

		sb.WriteRune(parser.c)
		parser.next()
	}

	quoted := sb.String()
	quoted = quoted[0 : sb.Len()-2]
	return node.NewConstantExpressionNode(parser.ResourceName, parser.lineNumber(), quoted)
}

/**
 * Parses a single directive token from the reader. Directives can be spelled with or without
 * braces, for example {@code #if} or {@code #{if}}. In the case of {@code #end}, {@code #else},
 * and {@code #elseif}, we return a {@link StopNode} representing just the token itself. In other
 * cases we also parse the complete directive, for example a complete {@code #foreach...#end}.
 */
func (parser *Parser) parseDirective() node.Node {
	var directive string
	if parser.c == '{' {
		parser.next()
		directive = parser.parseId("Directive inside #{...}")
		parser.expect('}')
	} else {
		directive = parser.parseId("Directive")
	}

	var localNode node.Node
	switch directive {
	case "end":
		localNode = node.NewEndNode(parser.ResourceName, parser.lineNumber())
		break

	case "if":
		return parser.parseIfOrElseIf("#if")

	case "elseif":
		localNode = node.NewElseIfNode(parser.ResourceName, parser.lineNumber())
		break

	case "else":
		localNode = node.NewElseNode(parser.ResourceName, parser.lineNumber())
		break

	case "foreach":
		return parser.parseForEach()

	case "set":
		localNode = parser.parseSet()
		break

	case "parse":
		localNode = parser.parseParse()
		break

	case "macro":
		return parser.parseMacroDefinition()

	default:
		localNode = parser.parsePossibleMacroCall(directive)
	}

	// Velocity skips a newline after any directive. In the case of #if etc, we'll have done this
	// when we stopped scanning the body at #end, so in those cases we return directly rather than
	// breaking into the code here.
	// TODO(emcmanus): in fact it also skips space before the newline, which should be implemented.
	if parser.c == '\n' {
		parser.next()
	}

	return localNode
}

/**
 * Parses plain text, which is text that contains neither {@code $} nor {@code #}. The given
 * {@code firstChar} is the first character of the plain text, and {@link #c} is the second
 * (if the plain text is more than one character).
 */
func (parser *Parser) parsePlainTextRune(firstChar rune) node.Node {
	builder := strings.Builder{}
	builder.WriteRune(firstChar)

	return parser.parsePlainTextWithBuilder(&builder)
}

/**
 * Parses plain text, which is text that contains neither {@code $} nor {@code #}. The given
 * {@code initialChars} are the first characters of the plain text, and {@link #c} is the
 * character after those.
 */
func (parser *Parser) parsePlainText(initialChars string) node.Node {
	builder := strings.Builder{}
	builder.WriteString(initialChars)

	return parser.parsePlainTextWithBuilder(&builder)
}

func (parser *Parser) parsePlainTextWithBuilder(builder *strings.Builder) node.Node {
	for true {
		if parser.c == EOF || parser.c == '$' || parser.c == '#' {
			break
		}

		// Just some random character.
		builder.WriteRune(parser.c)
		parser.next()
	}

	return node.NewConstantExpressionNode(parser.ResourceName, parser.lineNumber(), builder.String())
}

func (parser *Parser) parseNonDirective() node.Node {
	if parser.c == '$' {
		return parser.parseDollar()
	}

	firstChar := parser.c
	parser.next()
	return parser.parsePlainTextRune(firstChar)
}

func (parser *Parser) parseDollar() node.Node {
	utils.AssertRune(parser.c, '$')
	parser.next()

	silent := parser.c == '!'
	if silent {
		parser.next()
	}

	if utils.IsAsciiLetter(parser.c) || parser.c == '{' {
		return parser.parseReference(silent)
	}

	if silent {
		return parser.parsePlainText("$!")
	}

	return parser.parsePlainTextRune('$')
}

func (parser *Parser) parseIfOrElseIf(directive string) node.Node {
	startLine := parser.lineNumber()
	parser.expect('(')

	var condition node.ExpressionNode
	condition = parser.parseExpression()
	parser.expect(')')

	var parsedTruePart ParseResult

	description := "parsing " + directive + " starting on line " + fmt.Sprint(startLine)
	parsedTruePart = parser.skipNewlineAndParseToStop(isElseOrElseIfEndNode, description)

	var truePart node.Node
	truePart = node.NewConsNode(parser.ResourceName, startLine, parsedTruePart.Nodes)

	var falsePart node.Node

	if isEndNode(parsedTruePart.stop) {
		falsePart = node.EmptyNode(parser.ResourceName, parser.lineNumber())
	} else if isElseIfNode(parsedTruePart.stop) {
		falsePart = parser.parseIfOrElseIf("#elseif")
	} else {
		elseLine := parser.lineNumber()
		parsedFalsePart := parser.parseToStop(isEndNode, "parsing #else starting on line "+string(elseLine))
		falsePart = node.NewConsNode(parser.ResourceName, elseLine, parsedFalsePart.Nodes)
	}

	return node.NewIfNode(parser.ResourceName, startLine, condition, truePart, falsePart)
}

func (parser *Parser) skipNewlineAndParseToStop(stopClasses func(node node.Node) bool, contextDescription string) ParseResult {
	if parser.c == '\n' {
		parser.next()
	}

	return parser.parseToStop(stopClasses, contextDescription)
}

/**
 * Parses a {@code #foreach} token from the reader.
 *
 * <pre>{@code
 * #foreach ( $<id> in <expression> ) <body> #end
 * }</pre>
 */
func (parser *Parser) parseForEach() node.Node {
	startLine := parser.lineNumber()

	parser.expect('(')
	parser.expect('$')

	id := parser.parseId("For-each variable")

	parser.skipSpace()

	var bad bool
	bad = false

	if parser.c != 'i' {
		bad = true
	} else {
		parser.next()
		if parser.c != 'n' {
			bad = true
		}
	}

	if bad {
		panic(utils.ParseException("Expected 'in' for #foreach"))
	}

	parser.next()

	var collection node.ExpressionNode
	collection = parser.parseExpression()

	parser.expect(')')

	var parsedBody ParseResult
	parsedBody = parser.skipNewlineAndParseToStop(isEndNode, "parsing #foreach starting on line "+fmt.Sprint(startLine))

	var body node.Node
	body = node.NewConsNode(parser.ResourceName, startLine, parsedBody.Nodes)

	return node.NewForEachNode(parser.ResourceName, startLine, id, collection, body)
}

/**
 * Parses a {@code #set} token from the reader.
 *
 * <pre>{@code
 * #set ( $<id> = <expression> )
 * }</pre>
 */
func (parser *Parser) parseSet() node.Node {
	parser.expect('(')
	parser.expect('$')

	id := parser.parseId("#set variable")

	parser.expect('=')

	var expression node.ExpressionNode
	expression = parser.parseExpression()

	parser.expect(')')

	return node.NewSetNode(id, expression)
}

/**
 * Parses a {@code #parse} token from the reader.
 *
 * <pre>{@code
 * #parse ( <string-literal> )
 * }</pre>
 *
 * <p>The way this works is inconsistent with Velocity. In Velocity, the {@code #parse} directive
 * is evaluated when it is encountered during template evaluation. That means that the argument
 * can be a variable, and it also means that you can use {@code #if} to choose whether or not
 * to do the {@code #parse}. Neither of those is true in EscapeVelocity. The contents of the
 * {@code #parse} are integrated into the containing template pretty much as if they had been
 * written inline. That also means that EscapeVelocity allows forward references to macros
 * inside {@code #parse} directives, which Velocity does not.
 */
func (parser *Parser) parseParse() node.Node {
	panic(errors.New("parsing inner templated is not yet implemented"))
}

/**
 * Parses a {@code #macro} token from the reader.
 *
 * <pre>{@code
 * #macro ( <id> $<param1> $<param2> <...>) <body> #end
 * }</pre>
 *
 * <p>Macro parameters are optionally separated by commas.
 */
func (parser *Parser) parseMacroDefinition() node.Node {
	panic(errors.New("working with macro's is not yet implemented"))
}

/**
 * Parses an identifier after {@code #} that is not one of the standard directives. The assumption
 * is that it is a call of a macro that is defined in the template. Macro definitions are
 * extracted from the template during the second parsing phase (and not during evaluation of the
 * template as you might expect). This means that a macro can be called before it is defined.
 * <pre>{@code
 * #<id> ()
 * #<id> ( <expr1> )
 * #<id> ( <expr1> <expr2>)
 * #<id> ( <expr1> , <expr2>)
 * ...
 * }</pre>
 */
func (parser *Parser) parsePossibleMacroCall(directive string) node.Node {
	panic(errors.New("working with macro's is not yet implemented"))
}

/**
 * Parses a reference, which is everything that can start with a {@code $}. References can
 * optionally be enclosed in braces, so {@code $x} and {@code ${x}} are the same. Braces are
 * useful when text after the reference would otherwise be parsed as part of it. For example,
 * {@code ${x}y} is a reference to the variable {@code $x}, followed by the plain text {@code y}.
 * Of course {@code $xy} would be a reference to the variable {@code $xy}.
 * <pre>{@code
 * <reference> -> $<maybe-silent><reference-no-brace> |
 *                $<maybe-silent>{<reference-no-brace>}
 * <maybe-silent> -> <empty> | !
 * }</pre>
 *
 * <p>On entry to this method, {@link #c} is the character immediately after the {@code $}, or
 * the {@code !} if there is one.
 *
 * @param silent true if this is {@code $!}.
 */
func (parser *Parser) parseReference(silent bool) node.Node {
	if parser.c == '{' {
		parser.next()

		if !utils.IsAsciiLetter(parser.c) {
			if silent {
				return parser.parsePlainText("$!{")
			}
			return parser.parsePlainText("${")
		}
		node := parser.parseReferenceNoBrace(silent)
		parser.expect('}')
		return node
	}

	return parser.parseReferenceNoBrace(silent)
}

/**
 * Same as {@link #parseReference()}, except it really must be a reference. A {@code $} in
 * normal text doesn't start a reference if it is not followed by an identifier. But in an
 * expression, for example in {@code #if ($x == 23)}, {@code $} must be followed by an
 * identifier.
 *
 * <p>Velocity allows the {@code $!} syntax in these contexts, but it doesn't have any effect
 * since null values are allowed anyway.
 */
func (parser *Parser) parseRequiredReference() node.ReferenceNode {
	if parser.c == '!' {
		parser.next()
	}

	if parser.c == '{' {
		parser.next()
		node := parser.parseReferenceNoBrace( /* silent= */ false)
		parser.expect('}')
		return node
	}

	return parser.parseReferenceNoBrace( /* silent= */ false)
}

/**
 * Parses a reference, in the simple form without braces.
 * <pre>{@code
 * <reference-no-brace> -> <id><reference-suffix>
 * }</pre>
 */
func (parser *Parser) parseReferenceNoBrace(silent bool) node.ReferenceNode {
	id := parser.parseId("Reference")

	var lhs node.ReferenceNode
	lhs = node.NewPlainReferenceNode(parser.ResourceName, parser.lineNumber(), id, silent)

	return parser.parseReferenceSuffix(lhs, silent)
}

/**
 * Parses the modifiers that can appear at the tail of a reference.
 * <pre>{@code
 * <reference-suffix> -> <empty> |
 *                       <reference-member> |
 *                       <reference-index>
 * }</pre>
 *
 * @param lhs the reference node representing the first part of the reference
 *     {@code $x} in {@code $x.foo} or {@code $x.foo()}, or later {@code $x.y} in {@code $x.y.z}.
 */
func (parser *Parser) parseReferenceSuffix(lhs node.ReferenceNode, silent bool) node.ReferenceNode {
	switch parser.c {
	case '.':
		return parser.parseReferenceMember(lhs, silent)

	case '[':
		return parser.parseReferenceIndex(lhs, silent)

	default:
		return lhs
	}
}

/**
 * Parses a reference member, which is either a property reference like {@code $x.y} or a method
 * call like {@code $x.y($z)}.
 * <pre>{@code
 * <reference-member> -> .<id><reference-property-or-method><reference-suffix>
 * <reference-property-or-method> -> <id> |
 *                                   <id> ( <method-parameter-list> )
 * }</pre>
 *
 * @param lhs the reference node representing what appears to the left of the dot, like the
 *     {@code $x} in {@code $x.foo} or {@code $x.foo()}.
 */
func (parser *Parser) parseReferenceMember(lhs node.ReferenceNode, silent bool) node.ReferenceNode {
	return nil
}

/**
 * Parses the parameters to a method reference, like {@code $foo.bar($a, $b)}.
 * <pre>{@code
 * <method-parameter-list> -> <empty> |
 *                            <non-empty-method-parameter-list>
 * <non-empty-method-parameter-list> -> <primary> |
 *                                      <primary> , <non-empty-method-parameter-list>
 * }</pre>
 *
 * @param lhs the reference node representing what appears to the left of the dot, like the
 *     {@code $x} in {@code $x.foo()}.
 */
func (parser *Parser) parseReferenceMethodParams(lhs node.ReferenceNode, id string, silent bool) node.ReferenceNode {
	return nil
}

/**
 * Parses an index suffix to a reference, like {@code $x[$i]}.
 * <pre>{@code
 * <reference-index> -> [ <primary> ]
 * }</pre>
 *
 * @param lhs the reference node representing what appears to the left of the dot, like the
 *     {@code $x} in {@code $x[$i]}.
 */
func (parser *Parser) parseReferenceIndex(lhs node.ReferenceNode, silent bool) node.ReferenceNode {
	utils.AssertRune(parser.c, '[')
	parser.next()

	var index node.ExpressionNode
	index = parser.parsePrimary()

	if parser.c != ']' {
		panic(utils.ParseException("ParseException: Expected ]"))
	}

	parser.next()

	var reference node.ReferenceNode
	reference = node.NewIndexReferenceNode(lhs, index, silent)

	return parser.parseReferenceSuffix(reference, silent)
}

/**
 * Parses an expression, which can occur within a directive like {@code #if} or {@code #set}.
 * Arbitrary expressions <i>can't</i> appear within a reference like {@code $x[$a + $b]} or
 * {@code $x.m($a + $b)}, consistent with Velocity.
 * <pre>{@code
 * <expression> -> <and-expression> |
 *                 <expression> || <and-expression>
 * <and-expression> -> <relational-expression> |
 *                     <and-expression> && <relational-expression>
 * <equality-exression> -> <relational-expression> |
 *                         <equality-expression> <equality-op> <relational-expression>
 * <equality-op> -> == | !=
 * <relational-expression> -> <additive-expression> |
 *                            <relational-expression> <relation> <additive-expression>
 * <relation> -> < | <= | > | >=
 * <additive-expression> -> <multiplicative-expression> |
 *                          <additive-expression> <add-op> <multiplicative-expression>
 * <add-op> -> + | -
 * <multiplicative-expression> -> <unary-expression> |
 *                                <multiplicative-expression> <mult-op> <unary-expression>
 * <mult-op> -> * | / | %
 * }</pre>
 */
func (parser *Parser) parseExpression() node.ExpressionNode {
	var lhs node.ExpressionNode
	lhs = parser.parseUnaryExpression()

	return NewOperatorParser().Parse(*parser, lhs, 1)
}

/**
 * Parses an expression not containing any operators (except inside parentheses).
 * <pre>{@code
 * <unary-expression> -> <primary> |
 *                       ( <expression> ) |
 *                       ! <unary-expression>
 * }</pre>
 */
func (parser *Parser) parseUnaryExpression() node.ExpressionNode {
	parser.skipSpace()

	var localNode node.ExpressionNode

	if parser.c == '(' {
		parser.nextNonSpace()
		localNode = parser.parseExpression()
		parser.expect(')')
		parser.skipSpace()
		return localNode
	}

	if parser.c == '!' {
		parser.next()
		localNode = node.NewNotExpressionNode(parser.parseUnaryExpression())
		parser.skipSpace()
		return localNode
	}

	return parser.parsePrimary()
}

/**
 * Parses an expression containing only literals or references.
 * <pre>{@code
 * <primary> -> <reference> |
 *              <string-literal> |
 *              <integer-literal> |
 *              <boolean-literal> |
 *              <list-literal>
 * }</pre>
 */
func (parser *Parser) parsePrimary() node.ExpressionNode {
	return parser.parsePrimaryWithOptionalNull(false)
}

func (parser *Parser) parsePrimaryWithOptionalNull(nullAllowed bool) node.ExpressionNode {
	parser.skipSpace()
	var node node.ExpressionNode
	if parser.c == '$' {
		parser.next()
		node = parser.parseRequiredReference()
	} else if parser.c == '"' {
		node = parser.parseStringLiteral('"', true)
	} else if parser.c == '\'' {
		node = parser.parseStringLiteral('\'', false)
	} else if parser.c == '-' {
		// Velocity does not have a negation operator. If we see '-' it must be the start of a
		// negative integer literal.
		parser.next()
		node = parser.parseIntLiteral("-")
	} else if parser.c == '[' {
		node = parser.parseListLiteral()
	} else if utils.IsAsciiDigit(parser.c) {
		node = parser.parseIntLiteral("")
	} else if utils.IsAsciiLetter(parser.c) {
		node = parser.parseBooleanOrNullLiteral(nullAllowed)
	} else {
		panic(utils.ParseException("Expected a reference or a literal"))
	}
	parser.skipSpace()
	return node
}

/**
 * Parses a list or range literal.
 *
 * <pre>{@code
 * <list-literal> -> <empty-list> | <non-empty-list>
 * <empty-list> -> [ ]
 * <non-empty-list> -> [ <primary> <list-end>
 * <list-end> -> <range-end> | <remainder-of-list-literal>
 * <range-end> -> .. <primary> ]
 * <remainder-of-list-literal> -> <end-of-list> | , <primary> <remainder-of-list-literal>
 * <end-of-list> -> ]
 * }</pre>
 */
func (parser *Parser) parseListLiteral() node.ExpressionNode {
	utils.AssertRune(parser.c, '[')
	parser.nextNonSpace()

	if parser.c == ']' {
		parser.next()
		return node.NewListLiteralNode(parser.ResourceName, parser.lineNumber(), make([]node.ExpressionNode, 0))
	}

	var first node.ExpressionNode
	first = parser.parsePrimaryWithOptionalNull(false)

	if parser.c == '.' {
		return parser.parseRangeLiteral(first)
	}

	return parser.parseRemainderOfListLiteral(first)
}

func (parser *Parser) parseRangeLiteral(first node.ExpressionNode) node.ExpressionNode {
	utils.AssertRune(parser.c, '.')
	parser.next()
	if parser.c != '.' {
		panic(utils.ParseException("Expected two dots (..) not just one"))
	}

	parser.nextNonSpace()

	var last node.ExpressionNode
	last = parser.parsePrimaryWithOptionalNull(false)

	if parser.c != ']' {
		panic(utils.ParseException("Expected ] at end of range literal"))
	}

	parser.nextNonSpace()
	return node.NewRangeLiteralNode(parser.ResourceName, parser.lineNumber(), first, last)
}

func (parser *Parser) parseRemainderOfListLiteral(first node.ExpressionNode) node.ExpressionNode {
	return nil
}

func (parser *Parser) parseStringLiteral(quote rune, expand bool) node.ExpressionNode {
	return nil
}

func (parser *Parser) parseIntLiteral(prefix string) node.ExpressionNode {
	var sb strings.Builder
	sb.WriteString(prefix)

	for utils.IsAsciiDigit(parser.c) {
		sb.WriteRune(parser.c)
		parser.next()
	}

	str := sb.String()

	return node.NewConstantExpressionNode(parser.ResourceName, parser.lineNumber(), str)
}

func (parser *Parser) parseBooleanOrNullLiteral(nullAllowed bool) node.ExpressionNode {
	id := parser.parseId("Identifier without $")

	var value string

	switch id {
	case "true":
		value = "true"
		break

	case "false":
		value = "false"
		break

	case "null":
		if nullAllowed {
			value = ""
			break
		}

		// fall through...
	default:
		var suffix string
		if nullAllowed {
			suffix = " or null"
		} else {
			suffix = ""
		}
		panic(utils.ParseException("Identifier must be preceded by $ or be true or false" + suffix + ": " + id))
	}

	return node.NewConstantExpressionNode(parser.ResourceName, parser.lineNumber(), value)
}

func (parser *Parser) parseId(what string) string {
	if !utils.IsAsciiLetter(parser.c) {
		panic(utils.ParseException(what + " should start with an ASCII letter"))
	}

	var id strings.Builder
	for utils.IsIdChar(parser.c) {
		id.WriteRune(parser.c)
		parser.next()
	}

	return id.String()
}
