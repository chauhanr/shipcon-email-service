FROM debian:latest

RUN mkdir /app
WORKDIR /app
COPY shipcon-email-service /app/shipcon-email-service

CMD ["./shipcon-email-service"]