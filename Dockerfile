FROM golang:1.15.11-alpine3.13

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn

WORKDIR /app

COPY . .

RUN go build . && ls | grep -v cloud-fitter | xargs rm -rf && mkdir log/

EXPOSE 8081 9090

ENTRYPOINT ["./cloud-fitter", "-conf=config/config.yaml","-log_dir=log/"]