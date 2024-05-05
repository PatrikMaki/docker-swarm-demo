FROM alpine
RUN apk add curl
WORKDIR /manager
COPY client.sh .
ENTRYPOINT [ "./client.sh" ]