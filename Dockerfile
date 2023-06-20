FROM golang:alpine
ADD . /go/src/Noa/
ENV CGO_ENABLED 0
WORKDIR /go/src/Noa
RUN go build -o noa .
CMD ./noa
