# Parsing

> A parser is a software component that takes input data (frequently text) and builds a data structure – often some kind of parse tree, abstract syntax tree or other hierarchical structure – giving a structural representation of the input, checking for correct syntax in the process. […] The parser is often preceded by a separate lexical analyser, which creates tokens from the sequence of input characters;

- A parsing turns input into a data structure that represents that input.
- A common example of parser is `JSON.parse` in javascript
- MOst parsers produce an `abstract syntax tree`, **abstract** because some of the details visible in the source code are ommited from the AST, like whitespace, braces, comments etc
- Also called **syntactic analysis**

## Parsing strategies

- **top-down** and **bottom-up**

## Parsing expressions

- In Monkey, **expressions** produce values and **statements** dont!

- More complicated than parsing statements
- **Operator presidence** comes to mind first

```monkey
((5*5) + 10)
```

- `5*5` needs to be deeper in the ast and be parsed first
- The parser must be aware of the order of presidence
- Another interesting challenge is the same character in multiple positions `-5 - 10`. Here we have **prefix** and **infix** operators. Another example is as follows `5 * (add(2,3) + 10)`.
  - The outer pair of parens group the expressions and the inner pair denote a **call expression**. Validity now depends on context

### Pratt parser

- Main idea is associating paring functions (**semantic code**) with token types.
- Whenever a token type is encountered, the parsing functions are called to parse the appropriate expression and return the AST node that represents it.
- Each token type can have 2 types of function. One for **prefix** and one for **infix**

- Because of the 2 values, one on each side. These expressions are sometimes called "binary expressions"
