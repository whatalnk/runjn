# Create icon
* `go get github.com/josephspurrier/goversioninfo/cmd/goversioninfo`
* Download icon image from [here](https://github.com/jupyter/design/raw/master/logo/png-1x/jupyter-sq.png)
* `magick jupyter-sq.png -define icon:auto-resize=32 icon.ico`

# Build & install
* `go get github.com/mitchellh/go-ps`
* `git clone` into `%GOPATH%\src\github.com\whatalnk\runjn`
* `cd`
* `go generate`
* `go build`
* `go install github.com/whatalnk/runjn`
* associate to `*.ipynb`