FROM golang:1.17.7-alpine3.15

ENV SHORT_LINK_BASE=http://localhost:15001/

ENV POSTGRES_CONN_STRING=postgres://postgres:1@host.docker.internal:55000/postgres?sslmode=disable

ENV ALLOW_CORS_FROM=

RUN GOROOT=/go

WORKDIR /go/src/snaLinks

ADD . .

RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux go build -a -o snaLinks .

EXPOSE 15001

RUN chmod 777 /go/src/snaLinks

ENTRYPOINT [ "/go/src/snaLinks/snaLinks" ]
