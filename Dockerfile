FROM golang:alpine as builder

# 设置go mod proxy 国内代理
# 设置golang path
ENV GOPROXY=https://goproxy.cn,https://goproxy.io,direct \
    GO111MODULE=on \
    CGO_ENABLED=1
WORKDIR /chainweb
RUN go env -w GOPROXY=https://goproxy.cn,https://goproxy.io,direct
COPY . .
RUN go env && go list && go build -o app main.go

EXPOSE 8000
ENTRYPOINT /chainweb/app

# 根据Dockerfile生成Docker镜像
# docker build -t chainweb .
# 根据Docker镜像启动Docker容器
# docker run -itd -p 8000:8000 --name chainweb chainweb
