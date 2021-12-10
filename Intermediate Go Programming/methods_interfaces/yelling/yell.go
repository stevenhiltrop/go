package yelling

import "strings"

type Yeller interface {
	String() string
	Change(s string)
	Blank()
	Len() int
}

type LoudString struct {
	Yeller
	s string
}

func NewLoudString() LoudString {
	return LoudString{}
}

func (ls *LoudString) String() string {
	return ls.s
}

func (ls *LoudString) Change(s string) {
	ls.s = strings.ToUpper(s)
}

func (ls *LoudString) Blank() {
	ls.s = ""
}
