FROM golang:1.21.3-alpine3.18

WORKDIR /usr/src/app

COPY ./../ .

RUN go build server.go 

#EXPOSE "${PORT}"
#CMD ["ls"]
CMD ["./server"]


