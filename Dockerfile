FROM golang:1.25

WORKDIR /usr/src/app

RUN apt-get update && apt-get -y install vim
# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./ 
RUN go mod download

RUN go install -mod=mod github.com/githubnemo/CompileDaemon
RUN go get github.com/gin-gonic/gin
RUN go get github.com/go-sql-driver/mysql 
RUN go get github.com/go-playground/validator/v10
RUN go get github.com/pressly/goose/v3
RUN go get github.com/jackc/pgx/v5
COPY ./code ./code
RUN go mod tidy
RUN go build -o ./bin/migrate ./code/migrate

ENTRYPOINT /go/bin/CompileDaemon --build="go build -o main ./code" --command=./main -polling

