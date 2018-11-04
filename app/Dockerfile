FROM golang:1.11.2-alpine3.7 as builder
COPY main.go .
RUN apk update && apk upgrade && apk add --no-cache git openssh
RUN go get -u github.com/gin-gonic/gin
RUN go build -o /app .

FROM alpine:3.7
CMD ["./app"]
COPY --from=builder /app .
