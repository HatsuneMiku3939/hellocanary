# builder container
FROM golang:1.13-alpine3.10 AS builder
RUN apk update && apk add --no-cache git

WORKDIR /usr/src/app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o app

#####################
FROM alpine:3.10

# install SSL root certificates
RUN apk update && apk add ca-certificates && \
	rm -rf /var/cache/apk/* && \
	update-ca-certificates

COPY --from=builder /usr/src/app/app /var/app/

WORKDIR /var/app
CMD ["/var/app/app"]
