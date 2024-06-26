FROM golang:alpine

WORKDIR /app

COPY . /app

RUN go build -o gohard

ENTRYPOINT [ "./gohard" ]