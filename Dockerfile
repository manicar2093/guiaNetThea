FROM golang:1.16 as builder

WORKDIR /opt/app

COPY . .

RUN go build

FROM ubuntu:20.04

WORKDIR /opt/app

COPY --from=builder /opt/app/guianetThea .
COPY /templates ./templates
COPY /static ./static

RUN chmod 744 guianetThea

CMD ./guianetThea

