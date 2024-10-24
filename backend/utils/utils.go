package utils

import (
	"log"
	"neurodvach_backend/build"
	"os/exec"
	"runtime"
	"strings"
)

func FatalfIfDebug(format string, v ...any) {
	if build.Debug {
		log.Fatalf(format, v...)
	}
	log.Printf(format, v...)
}

func OpenBrowser(url string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	default: // "linux", "freebsd", "openbsd", "netbsd"
		cmd = "xdg-open"
	}
	args = append(args, url)
	return exec.Command(cmd, args...).Start()
}

func Htmlize(text string) string {
	return `<html><body>` + strings.Replace(strings.Trim(text, "\n\r\t "), "\n", "<br>", -1) + `</body></html>`
}
