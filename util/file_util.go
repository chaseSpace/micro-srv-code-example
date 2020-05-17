package util

import (
	"io/ioutil"
	"os"
)

func IterateDir(dir string, fn func(info os.FileInfo) bool) error {
	dirs, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}
	for _, dir := range dirs {
		if !fn(dir) {
			break
		}
	}
	return nil
}
