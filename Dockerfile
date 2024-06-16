<<<<<<< HEAD
FROM golang:1.22.3.4-alpine3.20 as builder
=======
FROM golang:1.22.3-alpine3.20 as builder
>>>>>>> c28bd83effe447dd1a7feafd3dd6723c6c99862c

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
