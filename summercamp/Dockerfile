FROM golang:1.7.4

RUN apt-get update -qq
ADD ./builder.sh /opt/builder.sh
RUN mkdir -p /go/src/github.com/Frozen-Team/summercamp/
WORKDIR /go/src/github.com/Frozen-Team/summercamp/
RUN go get  github.com/beego/bee \
            github.com/tools/godep \
            gopkg.in/testfixtures.v1 \
            github.com/astaxie/beego/migration \
            github.com/go-sql-driver/mysql
RUN apt-get install -y postgresql-client
RUN chmod a+x /opt/builder.sh
VOLUME /tmp/build/
CMD /opt/builder.sh
