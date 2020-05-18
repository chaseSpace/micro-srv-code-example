package pkg_test

import (
	"go_project_template/util"
	"testing"
)

func Test_GetTextByTime(t *testing.T) {
	t.Log(string(util.GetTextByTime()) != "")
}
