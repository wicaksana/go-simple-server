# syntax=docker/dockerfile:1

### Build
FROM golang:1.19.0-bullseye AS build

RUN mkdir /app
WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY web ./web

COPY *.go ./
RUN go build -o server .


### Deploy
FROM gcr.io/distroless/base-debian10
WORKDIR /
COPY --from=build /app/server /server
ENV GIN_MODE=release
EXPOSE 8080
USER nonroot:nonroot
ENTRYPOINT ["/server"]