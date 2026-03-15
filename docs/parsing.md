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

The following statement is parsed like this.

```go
{
    "-a * b",
    "((-a) * b)",
},
```

1. the parser calls `ParseProgram` which calls `parseStatement` for each of the
   statements, which are separated by semilcolons, until the end of file is
   encountered
2. this example is an expression statement and `parseExpressionStatement` is
   called
3. the first call to `parseExpression` with `LOWEST` will have `-` as the token
   type. we call `parsePrefixExpression` which then calls `parseIdentifier` for
   `a`. the inner call for `-` terminates because the `PREFIX` precedence is
   less than the `PRODUCT` precedence

   ```go
   ast.PrefixExpression{
       Token: token.MINUS,
       Operator: "-"
       Right: ast.Identifier{
           Token: token.IDENT,
           Value: "a"
       }
   }
   ```

4. The first call to `parseExpression` from `-` goes into the loop because the
   precedence `LOWEST` is less than `PRODUCT`. We advance the lexer and wrap the
   prefix expression (`leftExp`) inside of an infix expression. The left is
   set equal to the prefix expression while the right is set to the identifier

   ```go
   ast.InfixExpression{
       Token: token.ASTERISK,
       Operator: "*",
       Left: ast.PrefixExpression{
           Token: token.MINUS,
           Operator: "-"
           Right: ast.Identifier{
               Token: token.IDENT,
               Value: "a",
           },
       },
       Right: ast.Identifier{
           Token: token.IDENT,
           Value: "a",
       },
   }
   ```

## Prefix Operators

The prefix operators have the following form. It precedes an expression.

```text
<prefix-operator><expression>
```

## Infix Operators

The infix operators have the following form. It is sandwiched between two
expressions.

```text
<expression><infix-operator><expression>
```
