# To compile this image manually run:
#
# $ GO111MODULE=on GOOS=linux GOARCH=amd64 go build && docker build -t oryd/hydra:v1.0.0-rc.7_oryOS.10 . && rm hydra
FROM alpine:3.10

RUN apk add -U --no-cache ca-certificates

COPY run.sh /
RUN chmod a+x /run.sh

ENTRYPOINT ["/bin/sh"]

CMD ["/run.sh"]

FROM scratch

COPY --from=0 /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY hydra /usr/bin/hydra

USER 1000

ENTRYPOINT ["hydra"]
CMD ["serve", "all"]
