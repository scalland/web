FROM golang:-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /app/bin/ .

FROM scratch
COPY --from=builder /app/bin/ /
COPY --from=builder /app/configs /configs
COPY --from=builder /app/migrations /migrations
EXPOSE 0
ENTRYPOINT ["/", "serve"]
