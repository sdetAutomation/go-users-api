FROM golang:1.14

# go-sqlite3 requires cgo to work.
# ENV CGO_ENABLED=0

EXPOSE 8080

WORKDIR /build

COPY . .

RUN go build -o go-users-api

CMD ["./go-users-api"]