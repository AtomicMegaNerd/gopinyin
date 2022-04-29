FROM golang:1.18.1-alpine3.15 as builder

RUN apk add --no-cache musl-dev git \
    && rm -rf /var/cache/apk/*

WORKDIR /build
COPY . /build
RUN go build .

FROM alpine:3.15

ENV GID 1001
ENV UID 1001
ENV USER dockeruser
ENV PATH=/app/bin:${PATH}
ENV RUST_LOG=info

# Copy our built program over and 
RUN mkdir -p /app/bin /app/conf
COPY --from=builder /build/gopinyin /app/bin
WORKDIR /data

ENTRYPOINT ["gopinyin"]
