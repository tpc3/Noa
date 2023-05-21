FROM golang:latest AS build
RUN apt update && apt install -y mecab libmecab-dev mecab-ipadic-utf8
ENV GO111MODULE on
ENV CGO_LDFLAGS -L/usr/lib/x86_64-linux-gnu -lmecab -lstdc++
ENV CGO_CFLAGS -I/usr/include
RUN ln -s /usr/local/lib/libmecab.so.2.0.0 /lib64/libmecab.so.2
ADD . /go/src/Markov/
ARG GOARCH=amd64
ENV GOARCH ${GOARCH}
ENV CGO_ENABLED 1
WORKDIR /go/src/Markov
RUN go build -o markov .
CMD ["./markov"]