# Parsing

The purpose of the parser is to read in tokens from the lexer and create an
abstract syntax tree. At this stage in the interpreter, we are checking the
sequence of characters to ensure that the tokens conform to language's grammar.

## Pratt Parsing

Pratt parsing is a top-down parsing technique that associates precedence
directly into tokens. There are two key functions used per token, the null
denotation (nud or prefix) and left denotation (led or infix). Null denotation
describes how a token behaves at the start of an expression, while left
denotation describes how a token behaves after a left-hand expression.

Recursive descent parsers encode precedence rules directly into the grammar. One
issue with this approach is that any change to the precedence rules or opeartors
may require significant restructuring of the grammar.

The Pratt parser requires token handlers to be registered and relies on
abstraction and recursion to automatically handle the parsing.

## Prefix Operators

The prefix operators have the following form. It precedes an expression.

```text
<prefix-operator><expression>
```

In the parser code, the idea is to register `parsePrefixExpression` which
creates a prefix expression using the current token, as well as the following
expression.

## Infix Operators

The infix operators have the following form. It is sandwiched between two
expressions.

```text
<expression><infix-operator><expression>
```
