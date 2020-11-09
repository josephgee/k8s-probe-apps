FROM scratch

COPY cpu-loader /

#   nobody:nobody
USER 65534:65534
ENTRYPOINT ["/cpu-loader"]
