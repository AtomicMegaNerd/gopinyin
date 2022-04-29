FROM golang:1.18.1-alpine3.15 as builder

RUN apk add --no-cache git==2.34.2-r0 \
    && rm -rf /var/cache/apk/*

WORKDIR /build
COPY . /build
RUN go build .

FROM alpine:3.15

ENV USER=dockeruser
ENV PATH=/app/bin:${PATH}

# Copy our built program over and 
RUN mkdir -p /app/bin /app/conf \
    && adduser -D $USER \
    && chown $USER:$USER /app/bin
USER $USER
COPY --from=builder /build/gopinyin /app/bin
WORKDIR /data


ENTRYPOINT ["gopinyin"]
