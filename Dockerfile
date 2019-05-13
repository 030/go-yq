FROM golang:1.12.4-alpine as builder
COPY main.go go.mod go.sum ./go-yq/
WORKDIR go-yq
RUN adduser -D -g '' go-yq && \
    apk add git && \
    CGO_ENABLED=0 go build && \
    cp go-yq /go-yq && \
    chmod 100 /go-yq

FROM scratch
COPY --from=builder /etc/group /etc/group
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder --chown=go-yq:go-yq /go-yq /usr/local/go-yq
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
USER go-yq
ENTRYPOINT ["/usr/local/go-yq"]
