package yelling

import "strings"

type LoudString struct {
	s string
}

func NewLoudString() LoudString {
	return LoudString{}
}

func (ls *LoudString) String() string {
	return ls.s
}

func (ls LoudString) Change(s string) {
	ls.s = strings.ToUpper(s)
}

func (ls *LoudString) Blank() {
	ls.s = ""
}
