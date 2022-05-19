FROM golang:alpine
RUN mkdir /files
COPY helloworld.go /files
WORKDIR /files
RUN go build -o /files/helloworld helloworld.go
ENTRYPOINT ["/files/helloworld"]