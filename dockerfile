FROM golang:1.23.1

RUN echo "Creating working dir and copying files"
RUN mkdir /app
WORKDIR /app
COPY . .

CMD ["go", "run", "httpServer.go"]