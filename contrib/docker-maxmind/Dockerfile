FROM --platform=${BUILDPLATFORM} golang:1-alpine AS builder

ARG TARGETPLATFORM
ARG BUILDPLATFORM

WORKDIR $GOPATH/src/github.com/tracyhatemice/gogeoip-lookup

# Add libraries
RUN apk add --no-cache git upx

# Copy & build
ADD . $GOPATH/src/github.com/tracyhatemice/gogeoip-lookup

RUN set -eux; \
        case ${TARGETPLATFORM} in \
            "linux/amd64")  GOARCH=amd64  ;; \
            # arm64 and arm64v8 are equivalent in go and do not require a goarm
            # https://github.com/golang/go/wiki/GoArm
            "linux/arm64" | "linux/arm/v8")  GOARCH=arm64  ;; \
            "linux/ppc64le")  GOARCH=ppc64le  ;; \
            "linux/arm/v6") GOARCH=arm GOARM=6  ;; \
            "linux/arm/v7") GOARCH=arm GOARM=7 ;; \
        esac; \
        printf "Building geoip-lookup-service for arch ${GOARCH}\n"; \
        CGO_ENABLED=0 GOARCH=${GOARCH} go build -a -ldflags="-s -w" -o /geoip-lookup-service github.com/tracyhatemice/gogeoip-lookup/src/cmd; \
        upx --best --lzma /geoip-lookup-service

FROM scratch

COPY --from=builder /geoip-lookup-service /geoip-lookup-service

ENTRYPOINT ["/geoip-lookup-service"]
