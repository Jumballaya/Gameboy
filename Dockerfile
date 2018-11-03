# Builder
FROM golang as builder
WORKDIR /go/src/github.com/jumballaya/gameboy
COPY . .
RUN make build-linux

# Runner
FROM debian:stretch
WORKDIR /bin/
COPY --from=builder /go/src/github.com/jumballaya/gameboy/dist/gameboy_unix ./project
CMD ["gameboy"]
