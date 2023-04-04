FROM golang:1.20 AS build

RUN mkdir /build

WORKDIR /build

COPY . .

RUN make build

FROM ubuntu:latest

COPY --from=build /dist/gueue /opt/gueue

CMD ["/opt/gueue", "-c", "/etc/gueue/queue.yaml"]