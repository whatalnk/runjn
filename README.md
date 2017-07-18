# Build & install
* `go get github.com/mitchellh/go-homedir`
* `go get github.com/skratchdot/open-golang/open`
* `go get github.com/whatalnk/runjn`
* `cd $GOPATH/src/github.com/whatalnk/runjn`
* `go build`
* `go install github.com/whatalnk/runjn`

## Windows
* Associate to `*.ipynb`

## MacOS
* Use Automator

```
open -a Terminal $HOME/.go/bin/runjn $1

```

## Icon (Windows)
Before `go build`

* `go get github.com/josephspurrier/goversioninfo/cmd/goversioninfo`
* Download icon image from [here](https://github.com/jupyter/design/raw/master/logo/png-1x/jupyter-sq.png)
* `magick jupyter-sq.png -define icon:auto-resize=32 icon.ico`
* `go generate`