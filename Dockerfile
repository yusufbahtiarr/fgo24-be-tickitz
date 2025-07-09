FROM golang:alpine AS build

WORKDIR /build
COPY . .
RUN go build -o goapp main.go

FROM alpine:3.22

WORKDIR /app
COPY --from=build /build/goapp /app/goapp

ENTRYPOINT ["/app/goapp"]
