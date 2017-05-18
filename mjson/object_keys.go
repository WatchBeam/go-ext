package mjson

import "github.com/mailru/easyjson/jlexer"

// Keys lists object keys in the object.
func Keys(b []byte) ([]string, error) {
	keys := make([]string, 0, 16)
	in := jlexer.Lexer{Data: b}
	if in.IsNull() {
		return keys, nil
	}

	in.Delim('{')
	for !in.IsDelim('}') {
		keys = append(keys, in.UnsafeString())
		in.WantColon()
		in.SkipRecursive()
		in.WantComma()
	}
	in.Delim('}')

	return keys, in.Error()
}
