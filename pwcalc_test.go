package main

import "testing"

type testdata struct {
	alias  string
	secret string
	length int

	want string
}

var tests = []testdata{
	{"gmail.com", "mysecret", 1, "X"},
	{"gmail.com", "mysecret", 16, "XmMvBXYLVWOh2a2e"},
	{"gmail.com", "mysecret", 28, "XmMvBXYLVWOh2a2e5VVUDXwZnGU="},
}

func TestCalcPw(t *testing.T) {
	for _, c := range tests {
		pw := calcPw(c.secret, c.alias, c.length)
		if pw != c.want {
			t.Errorf("alias:%s secret:%s length:%d want:%s  but got %s",
				c.alias, c.secret, c.length, c.want, pw)
		}
	}
}
