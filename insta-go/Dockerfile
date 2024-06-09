FROM golang:latest
WORKDIR /app

COPY . .

RUN go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
RUN sqlc generate

RUN go run github.com/99designs/gqlgen

RUN go get
RUN go build -o /app/insta-go

EXPOSE 8080

CMD ["/app/insta-go"]
