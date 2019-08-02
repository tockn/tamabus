FROM golang:1.11.1
LABEL maintainer="tockn <s.taku.inthebox@gmail.com>"

RUN mkdir -p go/src/github.com/tockn/tamabus/webapp
ADD ./* $GOPATH/src/github.com/tockn/tamabus/webapp/
#RUN go get github.com/rubenv/sql-migrate/...
WORKDIR $GOPATH/src/github.com/tockn/tamabus/webapp
#ENV GO111MODULE=on
RUN go get -u github.com/golang/dep/cmd/dep
RUN dep ensure
CMD ["go", "run", "main.go"]
