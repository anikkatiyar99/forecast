FROM golang:1.17.2

WORKDIR /go/home
COPY ./ ./home

RUN cd /home && go build -o forecast

CMD ["./forecast"]