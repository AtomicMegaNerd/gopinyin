FROM golang:1.23.0-alpine3.20 AS builder

WORKDIR /build
COPY . /build
RUN go build .

FROM alpine:3.20

ENV USER=dockeruser
ENV PATH=/app/bin:${PATH}

RUN mkdir -p /app/bin /app/conf \
    && adduser -D $USER \
    && chown $USER:$USER /app/bin
USER $USER
COPY --from=builder /build/gopinyin /app/bin
WORKDIR /data

ENTRYPOINT ["gopinyin"]
