FROM golang:alpine
LABEL maintainer=" Gagan Kalra"
COPY . /go/src/app
WORKDIR /go/src/app
RUN go mod init
RUN go build -o gk-helloworld
EXPOSE 6111
CMD ["./gk-helloworld"]