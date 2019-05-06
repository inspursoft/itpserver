FROM 10.111.25.119/library/golang:1.12.4-alpine3.9

RUN mkdir -p /go/src/github.com/inspursoft/itpserver/src
COPY src/apiserver /go/src/github.com/inspursoft/itpserver/src/apiserver
WORKDIR /go/src/github.com/inspursoft/itpserver/src/apiserver
RUN go build -v -o /go/bin/apiserver && chmod u+x /go/bin/apiserver
CMD [ "/go/bin/apiserver" ]
EXPOSE 8088