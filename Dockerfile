FROM golang:1.15.11-alpine3.13

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn

WORKDIR /app

COPY . .

RUN go build .

EXPOSE 8081 9090

ENTRYPOINT ["./cloud-fitter conf=config.yaml"]