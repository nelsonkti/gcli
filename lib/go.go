package lib

import (
	"runtime"
	"strings"
)

func Version() string {
	return strings.Trim(runtime.Version(), "go")
}
