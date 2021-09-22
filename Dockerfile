FROM golang:1.15.5 as builder
RUN mkdir /app
ADD . /app/
WORKDIR /app
COPY fahLoadMain.go ./
RUN GOOS=linux go build -o fahLoadTest .
CMD ["/app/fahLoadTest"]

FROM alpine:latest as fahLoadTest
# Setup LD lib path
ENV LD_LIBRARY_PATH=/lib
# Install Apline dependencies
RUN apk add gcc libnsl libaio unzip openssl-dev autoconf musl-dev libc6-compat
ADD instantclient_12_2.zip /instantclient_12_2.zip
RUN unzip instantclient_12_2.zip && \
    cp -r instantclient_12_2/* /lib && \
    rm -rf instantclient_12_2.zip && \
    apk add libaio && \
    apk add libaio libnsl libc6-compat

RUN mkdir /app
COPY script.sh /app/script.sh
RUN /app/script.sh
WORKDIR /app
COPY --from=builder /app/fahLoadTest ./
CMD ["/app/fahLoadTest"]