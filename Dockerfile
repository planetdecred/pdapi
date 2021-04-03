FROM golang:1.16 as daemon

COPY . /go/src
WORKDIR /go/src/cmd/pdapi
RUN env GO111MODULE=on go build -v

FROM golang:1.16
WORKDIR /
COPY --from=daemon /go/src/cmd/pdapi/dcrdata /dcrdata

EXPOSE 7777
CMD [ "/dcrdata" ]
