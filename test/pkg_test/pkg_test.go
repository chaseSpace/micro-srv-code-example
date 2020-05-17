package pkg_test

import (
	"go_project_example/util"
	"testing"
)

func Test_GetTextByTime(t *testing.T) {
	t.Log(string(util.GetTextByTime()) != "")
}
