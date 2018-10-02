# Builder
FROM golang as builder
WORKDIR /go/src/github.com/username/project
COPY . .
RUN make build-linux

# Runner
FROM debian:stretch
WORKDIR /bin/
COPY --from=builder /go/src/github.com/username/project/dist/project_unix ./project
CMD ["project"]
