FROM golang:1.17-alpine as gobuild

WORKDIR /app
RUN apk update && apk upgrade && apk add --no-cache ca-certificates
RUN update-ca-certificates
COPY . .
RUN CGO_ENABLED=0 go build -o app .

FROM scratch
WORKDIR /app
COPY --from=gobuild /app/app .
COPY --from=gobuild /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

CMD ["./app"]
