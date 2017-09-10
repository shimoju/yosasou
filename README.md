# YOSASOU

LGTM Generator

## Setup

```bash
brew install go imagemagick
go get -u github.com/golang/dep/cmd/dep

git clone https://github.com/shimoju/yosasou.git "${GOPATH:-~/go}/src/github.com/shimoju/yosasou"
cd "${GOPATH:-~/go}/src/github.com/shimoju/yosasou"
dep ensure
```

## Run

```bash
go run server.go
```

Access [http://localhost:1323/](http://localhost:1323/)
