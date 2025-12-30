package app

import (
	"fmt"
	"strings"
)

func ErrNotImplemented(cmd string, args []string) error {
	suffix := ""
	if len(args) > 0 {
		suffix = " " + strings.Join(args, " ")
	}
	return fmt.Errorf("%s%s: not implemented", cmd, suffix)
}
