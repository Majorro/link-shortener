FROM golang:1.17.7-alpine3.15

RUN GOROOT=/go

WORKDIR /go/src/snaLinks
ADD . .
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -a -o snaLinks .

RUN sleep 10

EXPOSE 15001
RUN chmod 777 /go/src/snaLinks
ENTRYPOINT [ "/go/src/snaLinks/snaLinks" ]