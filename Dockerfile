FROM golang

ADD . /go/src/github.com/346285234/bbs-server
WORKDIR /go/src/github.com/346285234/bbs-server

RUN go get github.com/julienschmidt/httprouter
RUN go get github.com/go-sql-driver/mysql
RUN go install github.com/346285234/bbs-server

ENTRYPOINT /go/bin/bbs-server

EXPOSE 8080