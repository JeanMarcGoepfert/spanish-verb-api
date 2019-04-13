FROM golang:1.12.1

WORKDIR /src

COPY ./ ./

RUN go get -u ./

RUN go build ./

EXPOSE 8080

CMD ["spanish-api"]
