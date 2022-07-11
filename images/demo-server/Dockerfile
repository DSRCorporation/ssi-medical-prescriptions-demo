#
# Copyright SecureKey Technologies Inc. All Rights Reserved.
#
# SPDX-License-Identifier: Apache-2.0
#

ARG GO_VER
ARG ALPINE_VER

FROM golang:${GO_VER}-alpine${ALPINE_VER} as golang
RUN apk add --no-cache \
	gcc \
	musl-dev \
	git \
	libtool \
	bash \
	make;
ADD . $GOPATH/src/github.com/DSRCorporation/ssi-medical-prescriptions-demo
WORKDIR $GOPATH/src/github.com/DSRCorporation/ssi-medical-prescriptions-demo
ENV EXECUTABLES go git

FROM golang as ssi-medical-prescriptions-demo
LABEL org.opencontainers.image.source https://github.com/DSRCorporation/ssi-medical-prescriptions-demo

ARG GO_TAGS
ARG GOPROXY
RUN GO_TAGS=${GO_TAGS} GOPROXY=${GOPROXY} make demo-server


FROM alpine:${ALPINE_VER} as base
RUN apk add -U --no-cache ca-certificates
COPY --from=ssi-medical-prescriptions-demo /go/src/github.com/DSRCorporation/ssi-medical-prescriptions-demo/build/bin/demo-server /usr/local/bin
ENTRYPOINT ["demo-server"]
