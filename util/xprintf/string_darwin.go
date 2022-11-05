//go:build darwin
// +build darwin

package xprintf

import (
	"fmt"
)

const (
	RedColor    = 31
	GreenColor  = 32
	YellowColor = 33
	BlueColor   = 34
)

func Yellow(msg string, arg ...interface{}) string {
	if len(arg) > 0 {
		return sprint(YellowColor, msg, arg)
	} else {
		return sprint(YellowColor, msg)
	}
}

func Red(msg string, arg ...interface{}) string {
	if len(arg) > 0 {
		return sprint(RedColor, msg, arg)
	} else {
		return sprint(RedColor, msg)
	}
}

func Blue(msg string, arg ...interface{}) string {
	if len(arg) > 0 {
		return sprint(BlueColor, msg, arg)
	} else {
		return sprint(BlueColor, msg)
	}
}

func Green(msg string, arg ...interface{}) string {
	if len(arg) > 0 {
		return sprint(GreenColor, msg, arg)
	} else {
		return sprint(GreenColor, msg)
	}
}

func sprint(colorValue int, msg string, arg ...interface{}) string {
	var res string

	if arg != nil {
		res = fmt.Sprintf("\x1b[%dm%s\x1b[0m %+v", colorValue, msg, arg[0])
	} else {
		res = fmt.Sprintf("\x1b[%dm%s\x1b[0m", colorValue, msg)
	}

	return res
}
