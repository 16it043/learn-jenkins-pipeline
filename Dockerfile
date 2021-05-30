FROM scratch
EXPOSE 80:80
COPY helloworld .
CMD ["/helloworld"]