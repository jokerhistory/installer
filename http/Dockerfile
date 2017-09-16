M docker.io/phusion/baseimage:0.9.21
MAINTAINER dian.li <lidian@huawei.com>
RUN apt-get update && apt-get install -y tar git golang
# Usage: WORKDIR /path
RUN git clone https://github.com/Huawei/containerops.git
WORKDIR ~/containerops/component/python/
RUN tar -cvf ../checkstyle.tar -C ./analysischeckstyle/ .

ADD  main.go main.go
RUN go run main.go


