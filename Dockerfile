FROM golang:1.13.1-alpine

WORKDIR /app

COPY /server /app

COPY /ui/dist ui/dist

RUN go build -o . aquarium

EXPOSE 5000

CMD ["/app/aquarium"]