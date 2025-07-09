FROM golang:alpine AS build

WORKDIR /build

COPY . .

RUN go build -o gobetickitz main.go

FROM alpine:3.22

WORKDIR /app

COPY --from=build /build/gobetickitz /app/gobetickitz
COPY .env /app/.env

EXPOSE 8080

ENTRYPOINT ["/app/gobetickitz"]
