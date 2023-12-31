FROM golang:1.20

WORKDIR /go/src/app

COPY . .

RUN apt-get update && \
    apt-get install -y nodejs npm

RUN npm install -g puppeteer-cli

RUN go build -o puppeteer-go

EXPOSE 8080

CMD ["./puppeteer-go"]