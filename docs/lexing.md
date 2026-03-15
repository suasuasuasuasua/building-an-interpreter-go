# Lexer

The purpose of the lexer is to transform an input program (or a string of text)
into a sequence of tokens. The tokens are created using a list of predefined
symbols like `IDENT`, `SLASH`, and `EOF`. The lexer is responsible for handling
special cases like operators (say assignment vs equality) and keywords,
determining idenitifiers, and more.

It is the pre-processing stage before the real magic happens. A lexer has no
understanding of "what is right" according to the language. The only purpose is
to transform the input text into a set of atomic units known as tokens.

## Tokens

Tokens are defined in `token/`. This is the full lexicon or token set for the
language. The lexer needs to figure out which parts of the input text are what,
as well as determining keywords and identifiers for variables.
