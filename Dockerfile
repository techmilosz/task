FROM golang:1.21.0-alpine3.18 as builder
WORKDIR /app
COPY ./go.mod .
COPY ./go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /app/api ./main.go

FROM busybox:1.36.1
COPY --from=builder /app ./
ENTRYPOINT ["./api"]
