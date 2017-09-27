# YOSASOU

LGTM Generator

## Setup

```bash
brew install go imagemagick
go get -u github.com/golang/dep/cmd/dep
go get -u github.com/codegangsta/gin

git clone https://github.com/shimoju/yosasou.git "${GOPATH:-~/go}/src/github.com/shimoju/yosasou"
cd "${GOPATH:-~/go}/src/github.com/shimoju/yosasou"
dep ensure
```

## Run

```bash
gin --port 3300 --appPort 3301 --immediate
```

Access [http://localhost:3300/](http://localhost:3300/)
