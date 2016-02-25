//go:generate goversioninfo -icon=icon.ico
package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/mitchellh/go-ps"
)

func normpath(p string) string {
	return strings.Replace(p, "\\", "/", -1)
}
func main() {
	isrunning := false
	v, _ := ps.Processes()
	r := regexp.MustCompile(`jupyter`)
	for _, p := range v {
		if r.MatchString(p.Executable()) {
			isrunning = true
			break
		}
	}
	if isrunning {
		nbpath, _ := filepath.Rel(os.Getenv("Home"), os.Args[1])
		nburl, _ := url.Parse("http://localhost:8888/notebooks/")
		nburl.Path += normpath(nbpath)
		exec.Command("open", nburl.String()).Start()
	} else {
		nbpath, _ := filepath.Abs(os.Args[1])
		nbdir := fmt.Sprintf("--notebook-dir='%s'", normpath(os.Getenv("Home")))
		args := []string{"notebook", nbdir, nbpath}
		cmd := exec.Command("jupyter", args...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
		}
	}

}
