FROM christopher
WORKDIR /go/src/github.com/boojjee/christopher
# A hack to cache `go get`
RUN go get github.com/boojjee/christopher
ADD . /go/src/github.com/boojjee/christopher
RUN go get
RUN go build
EXPOSE 8080
ENTRYPOINT ./christopher -option value args