# Build & install
* `go get github.com/mitchellh/go-ps`
* `go get github.com/mitchellh/go-homedir`
* `go get github.com/skratchdot/open-golang/open`
* `go get github.com/whatalnk/runjn`
* `cd`
* `go generate` (if you need icon)
* `go build`
* `go install github.com/whatalnk/runjn`
* associate to `*.ipynb`

## Icon 
* `go get github.com/josephspurrier/goversioninfo/cmd/goversioninfo`
* Download icon image from [here](https://github.com/jupyter/design/raw/master/logo/png-1x/jupyter-sq.png)
* `magick jupyter-sq.png -define icon:auto-resize=32 icon.ico`

# Notes
* Supported Windows only
