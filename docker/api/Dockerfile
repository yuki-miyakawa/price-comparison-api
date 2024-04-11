# build
FROM golang:1.21-bullseye as builder
WORKDIR /opt/app
COPY go.mod go.sum ./
RUN go mod download
# ここでmigration処理
COPY ./*.go ./
COPY ./internal ./internal
RUN go build -trimpath -ldflags="-s -w" -o "main"

# production
FROM gcr.io/distroless/base-debian11 as production
COPY --from=builder /opt/app/main /main
ENV TEST=TEST-ENV
CMD ["/main"]

# develop
FROM golang:1.22-bullseye as develop
WORKDIR /app
RUN go install github.com/cosmtrek/air@latest
CMD ["air", "-c", ".air.toml"]
