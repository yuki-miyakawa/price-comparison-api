# build
FROM golang:1.22-bullseye as builder
WORKDIR /opt/app
COPY go.mod go.sum ./
RUN go mod download
# ここでmigration処理
RUN go install github.com/go-task/task/v3/cmd/task@latest
# RUN export PATH=$PATH:$(go env GOPATH)/bin
# COPY dbconfig.yml ./
# COPY .env.local ./
COPY Taskfile.yml ./
COPY database ./database
COPY sqlc.yml ./
RUN task setup
# RUN task migrate
COPY ./*.go ./
COPY ./internal ./internal
RUN sqlc generate
RUN go build -trimpath -ldflags="-s -w" -o "main"

# production
FROM gcr.io/distroless/base-debian11 as production
COPY --from=builder /opt/app/main /main
CMD ["/main"]

# develop
FROM golang:1.22-bullseye as develop
WORKDIR /app
RUN go install github.com/cosmtrek/air@latest
CMD ["air", "-c", ".air.toml"]
