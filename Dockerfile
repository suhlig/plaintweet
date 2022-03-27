FROM scratch
COPY plaintweet /
ENTRYPOINT ["/plaintweet"]
