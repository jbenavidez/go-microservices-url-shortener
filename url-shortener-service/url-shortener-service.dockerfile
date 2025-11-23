FROM alpine:latest

RUN mkdir /app

COPY urlShortenerApp /app

CMD [ "/app/urlShortenerApp"]