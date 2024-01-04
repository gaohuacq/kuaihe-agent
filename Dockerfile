FROM harbor.19k8s.cn/kuaihe-hz/golang:1.19-alpine as builder

WORKDIR /go/src/kuaihe-agent
COPY . .

# 切换源
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

# 设置golang环境变量 编译
RUN go env -w GO111MODULE=on && \
    go env -w GOPROXY=https://goproxy.cn,direct && \
    go env -w CGO_ENABLED=0 && \
    go mod tidy && \
    go build GOOS=linux GOARCH=amd64 -o kuaihe .

FROM alpine:latest

WORKDIR /go/src/kuaihe-agent

COPY --from=0 /go/src/kuaihe-agent/kuaihe ./
COPY --from=0 /go/src/kuaihe-agent/config.yaml ./

EXPOSE 8080
ENTRYPOINT ./kuaihe -c config.yaml

