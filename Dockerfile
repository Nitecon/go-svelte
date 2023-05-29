# Build the manager binary
FROM golang:alpine as builder

ENV GO111MODULE=auto \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

RUN apk --update add ca-certificates

WORKDIR /workspace
# Copy the Go Modules manifests
COPY . /workspace
# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN go mod download

# Build
RUN go build -ldflags "-s -w" -a -o app cmd/main.go

FROM node:alpine as nodebuilder
WORKDIR /web

COPY package.json /web
COPY svelte.config.js /web
COPY webpack.config.js /web
RUN npm install -g npm && npm install
COPY static /web/static
COPY view /web/view

RUN npm run build

FROM alpine:3.14

WORKDIR /
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /workspace/app .
COPY --from=nodebuilder /web/static /static
#USER nobody

ENTRYPOINT ["/app"]
EXPOSE 8080