FROM golang:1.8-onbuild

WORKDIR /go/src/applause
COPY . .

RUN go-wrapper download
RUN go-wrapper install

CMD ["go-wrapper", "run"]
