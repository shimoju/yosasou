FROM golang:1.9.0-alpine

WORKDIR $GOPATH/src/github.com/shimoju/yosasou

RUN apk --no-cache add \
    bash \
    git \
    build-base \
    imagemagick \
    imagemagick-dev \
  && go get -u github.com/golang/dep/cmd/dep \
  && go get -u github.com/codegangsta/gin

EXPOSE 3300
CMD ["gin", "--port", "3300", "--appPort", "3301", "--immediate"]
