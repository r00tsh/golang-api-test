FROM golang:1.22.1 AS builder

WORKDIR /src

COPY src .

RUN CGO_ENABLED=0 GOOS=linux go build -o app .

# hadolint ignore=DL3006
FROM gcr.io/distroless/base-debian12

COPY --from=builder /src/app /

EXPOSE 8080

CMD ["/app"]
