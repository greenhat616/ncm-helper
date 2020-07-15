package pad

import (
	"testing"
	"testing/quick"
	"unicode/utf8"
)

func TestUTF8LeftEqualWithSameLength(t *testing.T) {
	f := func(a string, pad string) bool {
		slen := utf8.RuneCountInString(a)
		padded := UTF8.Left(a, slen, pad)
		return padded == a
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestUTF8RightEqualWithSameLength(t *testing.T) {
	f := func(a string, pad string) bool {
		slen := utf8.RuneCountInString(a)
		padded := UTF8.Right(a, slen, pad)
		return padded == a
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestUTF8LeftEqualWithShorterLength(t *testing.T) {
	f := func(a string, pad string) bool {
		slen := utf8.RuneCountInString(a)
		padded := UTF8.Left(a, slen-3, pad)
		return padded == a
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestUTF8RightEqualWithShorterLength(t *testing.T) {
	f := func(a string, pad string) bool {
		slen := utf8.RuneCountInString(a)
		padded := UTF8.Right(a, slen-3, pad)
		return padded == a
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestUTF8LeftEqualWithGreaterLength(t *testing.T) {
	f := func(a string, pad string) bool {
		slen := utf8.RuneCountInString(a) + 3
		padded := UTF8.Left(a, slen, pad)
		return padded == times(pad, 3)+a
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestUTF8RightEqualWithGreaterLength(t *testing.T) {
	f := func(a string, pad string) bool {
		slen := utf8.RuneCountInString(a) + 3
		padded := UTF8.Right(a, slen, pad)
		return padded == a+times(pad, 3)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}
