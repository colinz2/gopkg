package utils

import (
	"errors"
	"io/ioutil"
	"os"
	"path"
)

var (
	ErrPathNotExist = errors.New("path not exist")
)

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, ErrPathNotExist
	}
	return false, err
}

func WalkAndLinkDir(baseDir string, toDir string) error {
	if e, _ := PathExists(toDir); !e {
		os.MkdirAll(toDir, os.ModePerm)
	}

	dirL, e := ioutil.ReadDir(baseDir)
	if e != nil {
		return e
	}

	for _, v := range dirL {
		fullPath := path.Join(baseDir, v.Name())
		fullToPath := path.Join(toDir, v.Name())
		if v.IsDir() {
			e = os.MkdirAll(fullToPath, os.ModePerm)
			if e != nil {
				return e
			}
			WalkAndLinkDir(fullPath, fullToPath)
		} else {
			e = os.Symlink(fullPath, fullToPath)
			if e != nil {
				return e
			}
		}
	}
	return nil
}
