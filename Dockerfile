FROM golang
ADD . /root/golang/src/christopher
EXPOSE 8080

#CMD ["go", "run", "/root/golang/src/christopher/server.go"]