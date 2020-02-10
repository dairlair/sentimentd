FROM alpine:latest
RUN apk --no-cache add ca-certificates

RUN mkdir /sentimentd
ADD ./build/sentimentd.linux-amd64 /sentimentd/sentimentd
RUN chmod +x /sentimentd/sentimentd

ADD ./schema /sentimentd/schema

# We create empty config which should be overriden with mount docker option.
# Also you can change application config through environment variables.
RUN mkdir /etc/sentimentd
RUN touch /etc/sentimentd/sentimentd.yml

CMD ["brain", "ls"]
ENTRYPOINT ["/sentimentd/sentimentd"]