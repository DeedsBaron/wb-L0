FROM golang:1.17 AS builder
COPY . /buf
WORKDIR /buf
RUN make build

FROM debian:buster
COPY --from=builder /buf ./
CMD ./wb-L0