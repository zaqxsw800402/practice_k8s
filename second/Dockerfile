FROM golang:alpine AS builder

WORKDIR /app

COPY . .

#RUN go mod init github.com/teampui/comico-cashier
#RUN go mod tidy
RUN go mod download

RUN GOARCH=arm64 GOOS=linux CGO_ENABLED=0 go build -o second practice_k8s/second

FROM scratch
COPY --from=builder /app/second /

EXPOSE 6666

CMD ["/second"]
