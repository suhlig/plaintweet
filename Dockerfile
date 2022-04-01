FROM alpine
RUN apk --no-cache add ca-certificates
ARG DIST=.
COPY ${DIST}/plaintweet /
ENTRYPOINT ["/plaintweet"]
