package utils

import "strings"

//This new type and package to be able to chain calls
type String string

func (s *String) TrimSpace() *String {
    *s = String(strings.TrimSpace(string(*s)))
    return s
}

func (s *String) ToUpper() *String {
    *s = String(strings.ToUpper(string(*s)))
    return s
}

func(s *String) ToString() string {
    return string(*s)
}
