FROM golang:1.26-alpine3.23 AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -trimpath -a -o server ./cmd/server/

FROM alpine:3.23
RUN apk add --no-cache tzdata
WORKDIR /app
COPY --from=builder /app/server .
COPY app.env .
COPY app_conf.yaml .
COPY --chmod=0755 start.sh .
COPY internal/db/migration ./internal/db/migration

CMD [ "/app/server" ]
