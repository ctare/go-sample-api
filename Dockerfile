FROM golang:latest

WORKDIR /app

# install go packages
RUN go install github.com/air-verse/air@latest
RUN go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest
RUN go get github.com/labstack/echo/v4
RUN go get -u gorm.io/gorm
RUN go get -u gorm.io/driver/mysql




