package lexer

// position and readPosition are used to access characters in the input
// by using them as an index. We need both to allow us to "peek" further into
// our input and look at the next character
type Lexer struct {
	input        string
	position     int // current position in input (points to the current char)
	readPosition int // current reading position in the input (after current char)
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// readChar - give us next character in input and
// advance our position in the input
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 // ASCI code fo "NUL"
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}
