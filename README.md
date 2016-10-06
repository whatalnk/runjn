# Build & install
* `go get github.com/mitchellh/go-ps`
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
* Set environmental variable `HOME`
    * Like `C:\Users\yourname`
* "open.exe" ?
    * Installed with [R](https://cran.r-project.org/)

```
>where open
C:\R\R-3.3.1\bin\x64\open.exe

>open --help
Usage: Rcmd open file [file ...]

  opens each file with the application given by
  the Windows file association (if any)
```
