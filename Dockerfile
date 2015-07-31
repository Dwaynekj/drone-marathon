# Docker image for Drone's slack notification plugin
#
#     CGO_ENABLED=0 go build -a -tags netgo
#     docker build --rm=true -t plugins/drone-marathon .

FROM gliderlabs/alpine:3.1
RUN apk-install ca-certificates
ADD drone-marathon /bin/
ENTRYPOINT ["/bin/drone-marathon"]
