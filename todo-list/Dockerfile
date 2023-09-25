FROM golang:alpine3.18

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN go mod download
RUN go build -o todo-list ./cmd/main.go

CMD ["./todo-list"]