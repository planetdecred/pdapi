FROM golang:1.16 as daemon

COPY . /go/src
WORKDIR /go/src/cmd/pdapi
RUN env GO111MODULE=on go build -v

FROM node:lts as gui

WORKDIR /root
COPY ./cmd/pdapi /root
RUN npm install
RUN npm run build

FROM golang:1.16
WORKDIR /
COPY --from=daemon /go/src/cmd/pdapi/dcrdata /dcrdata
COPY --from=daemon /go/src/cmd/pdapi/views /views
COPY --from=gui /root/public /public

EXPOSE 7777
CMD [ "/dcrdata" ]
