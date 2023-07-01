FROM golang

LABEL maintainer="Mamoudou Ndiaye"

COPY . /go/src/app

WORKDIR /go/src/app

# RUN go mod init 

EXPOSE 8080

CMD go run . 