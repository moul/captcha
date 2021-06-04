# dynamic config
ARG             BUILD_DATE
ARG             VCS_REF
ARG             VERSION

# build
FROM            golang:1.16.5-alpine as builder
RUN             apk add --no-cache git gcc musl-dev make
ENV             GO111MODULE=on
WORKDIR         /go/src/moul.io/captcha
COPY            go.* ./
RUN             go mod download
COPY            . ./
RUN             make install

# minimalist runtime
FROM alpine:3.13
LABEL           org.label-schema.build-date=$BUILD_DATE \
                org.label-schema.name="captcha" \
                org.label-schema.description="" \
                org.label-schema.url="https://moul.io/captcha/" \
                org.label-schema.vcs-ref=$VCS_REF \
                org.label-schema.vcs-url="https://github.com/moul/captcha" \
                org.label-schema.vendor="Manfred Touron" \
                org.label-schema.version=$VERSION \
                org.label-schema.schema-version="1.0" \
                org.label-schema.cmd="docker run -i -t --rm moul/captcha" \
                org.label-schema.help="docker exec -it $CONTAINER captcha --help"
COPY            --from=builder /go/bin/captcha /bin/
ENTRYPOINT      ["/bin/captcha"]
#CMD             []
