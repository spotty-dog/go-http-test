FROM gliderlabs/alpine

EXPOSE 80

ADD hello /bin/hello

CMD ["hello"]

