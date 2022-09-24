# Lexical analysis (Lexing)

- Also known as **tokenising**
- Tokens are small easily categorisable data structures that we can pass to our parser
- Some lexers will take into account white space. Like python
- Some will include line numbers and columns etc

```
let x = 5 + 5;

// might end up like this
[
  LET,
  IDENTIFIER("x"),
  EQUAL_SIGN,
  INTEGER(5),
  PLUS_SIGN,
  INTEGER(5),
  SEMICOLON
]
```

- Its important to define our tokens a we need a way to differentiate between integer, special characters etc

## Lexing

- position and readPosition are used to access characters in the input
  by using them as an index. We need both to allow us to "peek" further into
  our input and look at the next character

```go

type Lexer struct {
  input string
  position int // current position in input (points to the current char)
  readPosition int // current reading position in the input (after current char)
  ch byte
}
```

- The lexer only supports ASCII chars instead of the full UNicode range. This is to keep it simple.
- To support Unicode and UTF-8 we would need to change the `ch` to a rune type and change the way we read the next char to handle the multiple byte width
- We need to be able to skip whitepace as its only used for formatting. So languages include a token for newlines. We are skipping that
