package mjson

import "github.com/mailru/easyjson/jlexer"

// IsArray returns true whether the byte slice may be invalid JSON or if it's a
// an array. That is, if you don't know whether a JSON message contains `T[]`
// or `T`, IsArray will tell you whether you should try to unmarshal it as T[].
func IsArray(b []byte) bool {
	in := jlexer.Lexer{Data: b}
	if in.IsDelim('[') {
		return in.Error() == nil
	}

	return false
}
