// Based on gist:
// https://gist.github.com/ik5/d8ecde700972d4378d87#gistcomment-3074524

package utils

import "fmt"

var (
	Info  = Teal
	Warn  = Yellow
	Fatal = Red
)

var (
	Black   = Color("\033[1;30m%s\033[0m\n")
	Red     = Color("\033[1;31m%s\033[0m\n")
	Green   = Color("\033[1;32m%s\033[0m\n")
	Yellow  = Color("\033[1;33m%s\033[0m\n")
	Purple  = Color("\033[1;34m%s\033[0m\n")
	Magenta = Color("\033[1;35m%s\033[0m\n")
	Teal    = Color("\033[1;36m%s\033[0m\n")
	White   = Color("\033[1;37m%s\033[0m\n")
)

func Color(colorString string) func(...interface{}) (int, error) {
	sprint := func(args ...interface{}) (int, error) {
		return fmt.Printf(colorString,
			fmt.Sprint(args...))
	}
	return sprint
}
