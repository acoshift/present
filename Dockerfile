FROM golang:1.12.5-alpine

RUN apk add --no-cache git
RUN go get golang.org/x/tools/cmd/present

ADD slide /slide

EXPOSE 8080
WORKDIR /slide
CMD ["present", "-http=0.0.0.0:8080"]
