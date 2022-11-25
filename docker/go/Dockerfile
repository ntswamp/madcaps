FROM golang:latest

ENV PATH /usr/local/go/bin:$PATH

WORKDIR /home/madcaps
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...
 
EXPOSE 9090
CMD ["/usr/local/go/bin/go","run","main.go"]