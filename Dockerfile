FROM golang:1.16 AS builder

ENV GO111MODULE 'on'
ENV GOPROXY 'https://proxy.golang.org,direct'

RUN mkdir -p /app

WORKDIR /app

COPY . .

RUN mkdir -p bin
# For using private util repo
# RUN git config --global url."https://{username}:{key}@gitlab.com/".insteadOf "https://gitlab.com/"
# RUN go get gitlab.com/repo/util
RUN go build -o ./bin/main *.go
RUN rm -rf ./common ./components ./database ./middleware ./internal ./test
CMD ["/app/bin/main"]
