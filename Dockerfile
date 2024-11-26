FROM golang:1.23.2-alpine AS binary
WORKDIR /app
RUN apk update && apk upgrade && apk add --no-cache ca-certificates
RUN update-ca-certificates

COPY . .

RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -C cmd/ -ldflags="-w -s" -o server .
FROM scratch
COPY --from=binary /app/cmd /
COPY --from=binary /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
EXPOSE 8080
CMD [ "./server" ]
