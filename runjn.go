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

	"github.com/mitchellh/go-homedir"
	"github.com/mitchellh/go-ps"
	"github.com/skratchdot/open-golang/open"
)

func normpath(p string) string {
	return strings.Replace(p, "\\", "/", -1)
}
func main() {
	isrunning := false
	home, _ := homedir.Dir()
	v, _ := ps.Processes()
	r := regexp.MustCompile(`jupyter`)
	for _, p := range v {
		if r.MatchString(p.Executable()) {
			isrunning = true
			break
		}
	}
	if isrunning {
		nbpath, _ := filepath.Rel(home, os.Args[1])
		nburl, _ := url.Parse("http://localhost:8888/notebooks/")
		nburl.Path += normpath(nbpath)
		err := open.Start(nburl.String())
		if err != nil {
			log.Fatal(err)
		}
	} else {
		nbpath, _ := filepath.Abs(os.Args[1])
		nbdir := fmt.Sprintf("--notebook-dir='%s'", normpath(home))
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
