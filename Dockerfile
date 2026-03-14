FROM golang:1.26rc2-alpine

WORKDIR /usr/src/app

RUN apk update && apk add vim && apk add bash
# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./ 
RUN go mod download

RUN go install -mod=mod github.com/githubnemo/CompileDaemon
RUN go get github.com/gin-gonic/gin
RUN go get github.com/go-sql-driver/mysql 
RUN go get github.com/go-playground/validator/v10
RUN go get github.com/pressly/goose/v3
RUN go get github.com/jackc/pgx/v5
RUN go get github.com/joeshaw/envdecode
RUN go get gorm.io/gorm
RUN go get gorm.io/driver/postgres
RUN go get gorm.io/gorm/logger
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN go get github.com/rs/zerolog

COPY ./cmd ./cmd
COPY ./internal ./internal
COPY ./config ./config
COPY ./pkg ./pkg
RUN go mod tidy
RUN go build -o ./bin/migrate ./cmd/migrate

ENTRYPOINT /go/bin/CompileDaemon --build="go build -o main ./cmd/api" --command=./main -polling
