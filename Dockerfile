# Build the manager binary
FROM golang:1.19 as builder

# Copy the Go Modules manifests
COPY . /workspace

WORKDIR /workspace
# Build
RUN CGO_ENABLED=0 GOOS=linux GO111MODULE=on GOPROXY=https://mirrors.aliyun.com/goproxy/,direct go build -mod=vendor -o manager .

FROM alpine:latest
WORKDIR /
COPY --from=builder /workspace/manager .
ENTRYPOINT ["/manager"]
