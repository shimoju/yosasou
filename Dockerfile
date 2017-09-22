FROM golang:1.9.0

ENV LANG=C.UTF-8 \
  LC_ALL=C.UTF-8 \
  DEBIAN_FRONTEND=noninteractive

WORKDIR $GOPATH/src/github.com/shimoju/yosasou

RUN apt-get update -qq && apt-get install -qq --no-install-recommends \
    imagemagick \
    libmagickwand-dev \
  && rm -rf /var/lib/apt/lists/* \
  && go get -u github.com/golang/dep/cmd/dep \
  && go get -u github.com/codegangsta/gin

EXPOSE 3300
CMD ["gin", "--port", "3300", "--appPort", "3301", "--immediate"]
