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
-
