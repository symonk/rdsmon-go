# Build Stage
FROM lacion/alpine-golang-buildimage:1.13 AS build-stage

LABEL app="build-rdsmon-go"
LABEL REPO="https://github.com/symonk/rdsmon-go"

ENV PROJPATH=/go/src/github.com/symonk/rdsmon-go

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:$GOROOT/bin:$GOPATH/bin

ADD . /go/src/github.com/symonk/rdsmon-go
WORKDIR /go/src/github.com/symonk/rdsmon-go

RUN make build-alpine

# Final Stage
FROM lacion/alpine-base-image:latest

ARG GIT_COMMIT
ARG VERSION
LABEL REPO="https://github.com/symonk/rdsmon-go"
LABEL GIT_COMMIT=$GIT_COMMIT
LABEL VERSION=$VERSION

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:/opt/rdsmon-go/bin

WORKDIR /opt/rdsmon-go/bin

COPY --from=build-stage /go/src/github.com/symonk/rdsmon-go/bin/rdsmon-go /opt/rdsmon-go/bin/
RUN chmod +x /opt/rdsmon-go/bin/rdsmon-go

# Create appuser
RUN adduser -D -g '' rdsmon-go
USER rdsmon-go

ENTRYPOINT ["/usr/bin/dumb-init", "--"]

CMD ["/opt/rdsmon-go/bin/rdsmon-go"]
