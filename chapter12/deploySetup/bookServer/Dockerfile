FROM alpine
RUN apk update && apk add --no-cache supervisor
COPY supervisord.conf /etc/
COPY app /go/bin/app
RUN chmod +x /go/bin/app
CMD ["/usr/bin/supervisord", "-c", "/etc/supervisord.conf"]