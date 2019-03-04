FROM registry.cn-hangzhou.aliyuncs.com/choerodon-tools/golang:1.9.4-alpine3.7 as builder
WORKDIR /go/src/github.com/choerodon/c7n-gateway
COPY . .
RUN go build .

FROM registry.cn-hangzhou.aliyuncs.com/choerodon-tools/alpine:c7n
WORKDIR /
COPY --from=builder /go/src/github.com/choerodon/c7n-gateway .
CMD ["/c7n-gateway", "metrics-server"]
