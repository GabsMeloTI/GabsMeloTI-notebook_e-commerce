FROM golang:alpine

WORKDIR /sales

COPY . .

RUN go get -d -v ./...
RUN go build -o api .

EXPOSE 8001

CMD ["./api"]

