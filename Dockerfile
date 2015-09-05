FROM nanoservice/go

ADD . /go/src/app
WORKDIR /go/src/app

RUN go get ./...
RUN go test ./...
RUN go install

EXPOSE 8080

ENTRYPOINT ["app"]
