FROM golang:latest
MAINTAINER ACM@UIUC - Naren Dasan

ADD . /go/src/github.com/acm-uiuc/

RUN apt-get update && apt-get install -y ca-certificates git-core ssh

ADD keys/my_key_rsa /root/.ssh/id_rsa
RUN chmod 700 /root/.ssh/id_rsa
RUN echo "Host github.com\n\tStrictHostKeyChecking no\n" >> /root/.ssh/config
RUN git config --global url.ssh://git@github.com/.insteadOf https://github.com/


RUN go get github.com/gorilla/mux
RUN go get github.com/acm-uiuc/groot/proxy
RUN go get github.com/acm-uiuc/groot/services

RUN go install github.com/acm-uiuc/groot/server


CMD ["--port 8000"]
EXPOSE 8000


ENTRYPOINT /go/bin/server
