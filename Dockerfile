FROM cgr.dev/chainguard/go:latest as builder
WORKDIR /app
COPY src/* .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM scratch
COPY --from=builder /app/main /app/main
EXPOSE 8080
CMD ["/app/main"]
