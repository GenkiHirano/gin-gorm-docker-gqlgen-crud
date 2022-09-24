FROM golang:1.19-alpine
RUN mkdir /go/src/app
WORKDIR /go/src/app
ADD . /go/src/app
RUN apk update && apk add git
# go.sumも追加する
COPY go.mod ./
RUN go mod download
CMD ["go", "run", "main.go"]
