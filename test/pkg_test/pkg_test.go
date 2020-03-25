package pkg_test

import (
	"go_project_example/pkg"
	"testing"
)

func Test_GetTextByTime(t *testing.T) {
	t.Log(string(pkg.GetTextByTime()) != "")
}
