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

	"path"

	"github.com/mitchellh/go-homedir"
	"github.com/skratchdot/open-golang/open"
)

func normpath(p string) string {
	return strings.Replace(p, "\\", "/", -1)
}
func main() {
	home, _ := homedir.Dir()
	jupyter, err := exec.LookPath("jupyter")
	if err != nil {
		log.Fatal(err)
	}
	out, err := exec.Command(jupyter, "notebook", "list").Output()
	re := regexp.MustCompile(`http://localhost:[0-9]+/`)
	address := re.FindString(string(out))
	if address != "" {
		nbpath, _ := filepath.Rel(home, os.Args[1])
		nburl, _ := url.Parse(address)
		nburl.Path = path.Join(nburl.Path, normpath(nbpath))
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
