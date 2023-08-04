FROM --platform=linux/amd64 registry.ap-northeast-1.aliyuncs.com/taimer/golang:1.19 AS builder
WORKDIR /go/src/github.com/JingruiLea/ad_boost
COPY . .
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn
ENV GOFLAGS="-mod=vendor"

RUN chmod +x ./build.sh
RUN ./build.sh

FROM --platform=linux/amd64 registry.ap-northeast-1.aliyuncs.com/taimer/alpine:3.18
COPY --from=builder /go/src/github.com/JingruiLea/ad_boost/output /opt/output

ENV LANG zh_CN.UTF-8
ENV TZ=Asia/Shanghai
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

WORKDIR /opt/output

EXPOSE 9000

ENTRYPOINT ["/bin/sh", "./bootstrap.sh"]