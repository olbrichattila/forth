// Package lexer tokenize the source code
// The lexical analyzer main purpose is to break down raw source code into manageable pieces called tokens
package lexer

// Lexer abstracts the logic from the concrete implementation
type Lexer interface {
	Tokenize(code string) ([]Token, error)
}

// New creates a new lexer
func New() Lexer {
	return &tokenizer{}
}

var tokens = map[string]TokenType{
	"+":  TokenTypeAdd,
	"-": TokenTypeSub,
	"*": TokenTypeMultiply,
	"dup": TokenTypeDup,
	"drop": TokenTypeDrop,
	"swap": TokenTypeSwap,
	".": TokenTypePrint,
	":": TokenTypeFunction,
	";": TokenEndFunc,
}

type tokenizer struct {
	code string
	pos int
}

// Tokenize will convert the source code to tokens
func (t *tokenizer) Tokenize(code string) ([]Token, error) {
	result := make([]Token, 0)

	t.code = code;
	t.pos = 0;

	for {
		if t.pos == len(code) {
			break
		}

		char := string(code[t.pos])
		if t.ignoreCharacter(char) {
			t.pos++;
			continue
		}
		
		if val, ok := t.parseIntToken(); ok {
			result = append(result, Token{t: TokenTypeNumber, value: val})
			continue
		}

		if val, ok := t.parseNameToken(); ok {
			if tk, ok := tokens[val]; ok {
				result = append(result, Token{t: tk, value: val})
			} else {
				result = append(result, Token{t: TokenTypeName, value: val})
			}
	
			continue
		}

		if tk, ok := tokens[char]; ok {
			result = append(result, Token{t: tk, value: char})
		}
		
		t.pos++;
	}

	return result, nil;
}

func (t *tokenizer) parseIntToken() (string, bool) {
	num := ""
	for {
		if t.pos == len(t.code) || !t.isInt(string(t.code[t.pos])) {
			return num, num != ""
		}
		
		num += string(string(t.code[t.pos]))
		t.pos++
	}
}

func (t *tokenizer) parseNameToken() (string, bool) {
	result := ""
	for {
		if t.pos == len(t.code) || !t.isName(string(t.code[t.pos])) {
			return result, result != ""
		}
		
		result += string(string(t.code[t.pos]))
		t.pos++
	}
}

func (t *tokenizer) isInt(s string) bool {
	return s >= "0" && s <= "9"
}

func (t *tokenizer) isName(s string) bool {
	return (s >= "a" && s <= "z") || (s >= "A" && s <= "Z")
}

func (t *tokenizer) ignoreCharacter(s string) bool {
	return s == " " || s == "\n" || s == "\r"  || s == "\t"
}
