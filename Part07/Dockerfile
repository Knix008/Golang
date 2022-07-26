FROM golang:latest
RUN mkdir /files
COPY keyValueTCP.go /files
WORKDIR /files
RUN go build -o /files/keyValueTCP keyValueTCP.go
ENTRYPOINT ["/files/keyValueTCP","80"]