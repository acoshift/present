FROM golang:1.9

RUN go get golang.org/x/tools/cmd/present

ADD nevp /slides/nevp

EXPOSE 3999
WORKDIR /slides
CMD ["present", "-http=0.0.0.0:3999"]

