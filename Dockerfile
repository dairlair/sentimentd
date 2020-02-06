########################################################################################################################
# Final stage                                                                                                          #
########################################################################################################################
FROM alpine:latest
RUN apk --no-cache add ca-certificates

ADD ./build/sentimentd.linux-amd64 /bin/sentimentd
RUN chmod +x /bin/sentimentd

ADD ./schema /var/lib/sentimentd/schema

# We create empty config which should be overriden with mount docker option.
# Also you can change application config through environment variables.
RUN mkdir /etc/sentimentd
RUN touch /etc/sentimentd/sentimentd.yml

CMD ["brain ls"]
ENTRYPOINT ["/bin/sentimentd"]