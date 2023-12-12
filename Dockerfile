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
