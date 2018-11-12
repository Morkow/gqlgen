FROM golang:1.10

RUN curl -L -o /bin/dep https://github.com/golang/dep/releases/download/v0.4.1/dep-linux-amd64 && chmod +x /bin/dep
RUN go get -u github.com/alecthomas/gometalinter github.com/Morkow/gorunpkg
RUN gometalinter --install

WORKDIR /go/src/github.com/Morkow/gqlgen

COPY Gopkg.* /go/src/github.com/Morkow/gqlgen/
RUN dep ensure -v --vendor-only

COPY . /go/src/github.com/Morkow/gqlgen/
