FROM golang:alpine AS builder

WORKDIR /app

COPY . .

#RUN go mod init github.com/teampui/comico-cashier
#RUN go mod tidy
RUN go mod download

RUN GOARCH=arm64 GOOS=linux CGO_ENABLED=0 go build -o first practice_k8s/first

FROM alpine
COPY --from=builder /app/first /

RUN apk add --no-cache bash

EXPOSE 6666

CMD ["/first"]
