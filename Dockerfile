FROM golang:latest as golang
WORKDIR /app
ENV CGO_ENABLED=0
COPY . .
RUN go build -o main .

FROM alpine:latest
WORKDIR /app
COPY --from=golang /app/main .
EXPOSE 5030
CMD ["./main"]