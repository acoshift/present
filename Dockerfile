FROM golang:1.19.3-alpine

RUN apk add --no-cache git
WORKDIR /workspace
ADD . .
RUN go mod download
RUN go install golang.org/x/tools/cmd/present@latest

EXPOSE 8080
CMD ["present", "-http=0.0.0.0:8080", "-base=.", "-content=slide", "-use_playground"]
