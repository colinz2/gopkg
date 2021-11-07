package utils

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func FormatPath(s string) string {
	switch runtime.GOOS {
	case "windows":
		return strings.Replace(s, "/", "\\", -1)
	case "darwin", "linux":
		return strings.Replace(s, "\\", "/", -1)
	default:
		return s
	}
}

func copyDir(src string, dest string) (e error) {
	src = FormatPath(src)
	dest = FormatPath(dest)

	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("xcopy", src, dest, "/I", "/E")
	case "darwin", "linux":
		cmd = exec.Command("cp", "-R", src, dest)
	}

	_, e = cmd.Output()
	if e != nil {
		return
	}
	return
}

func CopyFile2(dst, src string) (err error) {
	content, err := ioutil.ReadFile(src)
	if err != nil {
		return
	}
	err = ioutil.WriteFile(dst, content, os.ModePerm)
	if err != nil {
		return
	}
	return
}

func CopyFile(dstFilePath string, srcFilePath string) (written int64, err error) {
	srcFile, err := os.Open(srcFilePath)
	if err != nil {
		return
	}
	defer srcFile.Close()
	reader := bufio.NewReader(srcFile)

	dstFile, err := os.OpenFile(dstFilePath, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return
	}
	writer := bufio.NewWriter(dstFile)
	defer dstFile.Close()
	return io.Copy(writer, reader)
}
