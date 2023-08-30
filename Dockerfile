FROM golang:alpine

WORKDIR /app

COPY ./ ./

RUN go build -o main cmd/segmentation-service/main.go

ENTRYPOINT ./main