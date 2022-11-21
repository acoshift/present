FROM golang:1.19.3-alpine

RUN apk add --no-cache git
ADD . /workspace/

EXPOSE 8080
WORKDIR /workspace/slide
CMD ["present", "-http=0.0.0.0:8080"]
