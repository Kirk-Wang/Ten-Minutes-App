FROM node:10.15.1-alpine

RUN apk add --no-cache tini && npm install http-server -g && mkdir /ten

WORKDIR /ten

COPY app/build .

EXPOSE 3000

ENTRYPOINT ["/sbin/tini", "--"]
CMD [ "http-server", "-p", "3000" ]
