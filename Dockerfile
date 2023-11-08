FROM golang:1.20-alpine AS build-base

WORKDIR /app

COPY go.mod .

RUN go mod download

RUN go install github.com/swaggo/swag/cmd/swag@latest

COPY . .

ENV GOPATH /go

RUN swag init

RUN go build -o ./out/todo-backend

FROM alpine:3.16.2 AS production

WORKDIR /app

EXPOSE ${PORT}

COPY --from=build-base /app/out/todo-backend /app/todo-backend

CMD ["/app/todo-backend"]