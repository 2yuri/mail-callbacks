FROM alpine

COPY ./build /opt/email-dispatcher
WORKDIR /opt/email-dispatcher

EXPOSE 8000

ENTRYPOINT /opt/email-dispatcher/mail-callbacks
