# syntax=docker/dockerfile:1-labs

FROM cgr.dev/chainguard/go:latest-dev as build
WORKDIR /work
COPY src/go.* src/*.go ./
# COPY  ./
# Compile without CGO to avoid dynamic linking
RUN CGO_ENABLED=0 go build -o tms .
# Compress binary
RUN apk add upx && upx --lzma --best -votms.packed tms && upx -t tms.packed

FROM cgr.dev/chainguard/static:latest
COPY --from=build /work/tms.packed /tms
CMD [ "/tms" ]
LABEL org.opencontainers.image.authors='goproslowyo@gmail.com'
LABEL org.opencontainers.image.description="Talos Metadata Service"
LABEL org.opencontainers.image.licenses='LGPL-3.0'
LABEL org.opencontainers.image.source='https://github.com/goproslowyo/tms'
LABEL org.opencontainers.image.url='https://github.com/users/goproslowyo/packages/container/package/tms'
LABEL org.opencontainers.image.vendor='GoProSlowYo'
