FROM golang:latest

ENV CGO_ENABLED 0
ENV GOOS linux
ENV PORT=9100

COPY ./codebase /app
WORKDIR /app

RUN go get -d \
  && go get github.com/lib/pq \
  && go build -o main.go

CMD ["go", "run", "main.go"]
