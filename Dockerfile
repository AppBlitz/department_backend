FROM golang:1.26 as build

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -v -o /usr/local/bin/app ./...

FROM alpine:latest

WORKDIR ./app

COPY --from=builder /usr/src/app/app .

CMD ["./app"]
