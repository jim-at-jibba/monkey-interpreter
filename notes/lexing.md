# Lexical analysis

- To make our code easier to work with we must go through a process `Source Code -> Tokens -> AST`
- Te first transformation is called _lexing_. Done by a lexer or tokeniser
- Tokens are easy to categorise data strinctures which are fed to a parserwhich will turn the tokens into the AST
- The basics of our lexer is to take source code as input nad output tokens
- We need `position` and `readPosition` in our Lexer struct to be able to peek futher into the input to see what is next
- the `readChar` function is to give us the next char and advance our position in the input
