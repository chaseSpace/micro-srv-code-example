package pkg

import (
	"fmt"
)

// Customize your public errors
var (
	EnvNotSetErr = MyErr{id: "EnvNotSetErr"}
)

type MyErr struct {
	id         string
	extra      string // print with together, generally, set it on init.
	NewErrInfo string
}

func (e *MyErr) String() string {
	s := fmt.Sprintf("%s --> %s", e.id, e.NewErrInfo)
	if e.extra != "" {
		s = fmt.Sprintf("%s\nEXTRA:%s", s, e.extra)
	}
	return s
}

func (e *MyErr) New(s string, args ...interface{}) {
	e.NewErrInfo = fmt.Sprintf(s, args...)
}

func (e *MyErr) Panic(s string, args ...interface{}) {
	e.NewErrInfo = fmt.Sprintf(s, args...)
	panic(e.String())
}

func (e *MyErr) Extra(s string, args ...interface{}) {
	e.extra = fmt.Sprintf(s, args...)
}
