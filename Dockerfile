FROM golang:1.20.10-alpine3.17 as builder
WORKDIR /go/src/k8s-mch/server
COPY . .

RUN go env -w GO111MODULE=on \
   && go env -w GOPROXY=https://goproxy.cn,direct \
   && go env -w CGO_ENABLED=0 \
   && go env \
   && go mod tidy \
   && go build -o server .

FROM alpine:latest

LABEL MAINTAINER="michenghua@qq.com"

WORKDIR /go/src/k8s-mch/server
COPY --from=0 /go/src/k8s-mch/server/config.yaml ./config.yaml
COPY --from=0 /go/src/k8s-mch/server/.kube/config ./.kube/config
COPY --from=0 /go/src/k8s-mch/server/server ./
EXPOSE 8081
ENTRYPOINT ./server