FROM golang:1.22.0-alpine3.19 as builder

WORKDIR /build
COPY . /build
RUN go build .

FROM alpine:3.19

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
