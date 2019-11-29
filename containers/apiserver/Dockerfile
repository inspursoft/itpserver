FROM golang:1.12.6-alpine3.9

RUN mkdir -p /go/src/github.com/inspursoft/itpserver/src \
 && echo http://dl-2.alpinelinux.org/alpine/edge/community/ >> /etc/apk/repositories \
 && apk --no-cache --no-progress add git \
 && go get github.com/beego/bee

COPY src/apiserver /go/src/github.com/inspursoft/itpserver/src/apiserver
WORKDIR /go/src/github.com/inspursoft/itpserver/src/apiserver

CMD [ "bee", "run", "-gendoc=true", "-downdoc=true" ]
EXPOSE 8088
