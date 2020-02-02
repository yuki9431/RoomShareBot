package main

import "regexp"

// IsNumber 引数が数値か判別する
func IsNumber(s string) (isNumber bool) {
	var correctNumber = regexp.MustCompile(`^[0-9]$`)
	correctNumber.MatchString(s)

	return
}
