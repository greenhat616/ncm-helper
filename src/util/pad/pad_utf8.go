package pad

import "unicode/utf8"

type utf8S struct {
}

// UTF8 is a var that impl the utf8s struct, a collection of tools to pad the string in UTF8
var UTF8 = &utf8S{}

func (p *utf8S) times(str string, n int) (out string) {
	for i := 0; i < n; i++ {
		out += str
	}
	return
}

// Left left-pads the string with pad up to len runes
// len may be exceeded if
func (p *utf8S) Left(str string, len int, pad string) string {
	return p.times(pad, len-utf8.RuneCountInString(str)) + str
}

// Right right-pads the string with pad up to len runes
func (p *utf8S) Right(str string, len int, pad string) string {
	return str + p.times(pad, len-utf8.RuneCountInString(str))
}
