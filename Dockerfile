FROM gliderlabs/alpine

EXPOSE 80

ADD hello /hello

CMD ["hello"]

