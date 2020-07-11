FROM golang:1.14
WORKDIR /app
COPY go.mod .
RUN go mod download
COPY . .
RUN make build
CMD ["./bin/telescope"]