FROM golang:1.14
WORKDIR /app
COPY go.mod .
copy go.sum .
RUN go mod download
COPY . .
RUN make build
CMD ["./bin/telescope"]

