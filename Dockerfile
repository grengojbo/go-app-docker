FROM scratch

ENV HOST 0.0.0.0
ENV PORT 7000
ENV LOG_LEVEL error

EXPOSE $PORT

#COPY certs /etc/ssl/certs/
COPY bin/linux-amd64/k8sapp /

CMD ["/k8sapp"]
