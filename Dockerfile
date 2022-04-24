# syntax=docker/dockerfile:1

FROM golang:1.16-buster AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /seagal-bot-go

FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /seagal-bot-go /seagal-bot-go

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/seagal-bot-go"]
