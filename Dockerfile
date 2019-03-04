FROM registry.cn-hangzhou.aliyuncs.com/choerodon-tools/golang:1.9.4-alpine3.7 as builder
WORKDIR /go/src/MiniHttpFileServer
COPY . /go/src/MiniHttpFileServer
RUN go build .

FROM registry.saas.hand-china.com/tools/alpine:latest
WORKDIR /
COPY --from=builder /go/src/MiniHttpFileServer .
CMD ["/MiniHttpFileServer"]
