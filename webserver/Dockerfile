FROM golang:1.20
WORKDIR /usr/src/app
COPY go.mod ./
COPY go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN go build main.go
EXPOSE 3000
CMD ["./main"]
