FROM golang:1.9

RUN go get golang.org/x/tools/cmd/present

ADD slide /slide

EXPOSE 3999
WORKDIR /slide
CMD ["present", "-http=0.0.0.0:3999"]

