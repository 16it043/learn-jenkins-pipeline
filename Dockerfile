FROM golang:latest

EXPOSE 80:80

COPY helloworld .

CMD ["./helloworld"]