FROM golang:1.8



RUN mkdir -p /go/src/Monitoring/prom-json-exporter-mod/
WORKDIR /go/src/Monitoring/prom-json-exporter-mod/

COPY . /go/src/Monitoring/prom-json-exporter-mod/
RUN go-wrapper download && go-wrapper install

CMD ["go-wrapper", "run", "--interval", "15"]

EXPOSE 7979