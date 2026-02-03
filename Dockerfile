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
COPY ./code .

ENTRYPOINT /go/bin/CompileDaemon --build="go build -o main ./code" --command=./main -polling

#RUN go build -v -o /usr/bin/app ./...


