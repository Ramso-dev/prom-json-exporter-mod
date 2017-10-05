FROM golang:1.8



RUN mkdir -p /go/src/Monitoring/prometheus-json-exporter-master/
WORKDIR /go/src/Monitoring/prometheus-json-exporter-master/

COPY . /go/src/Monitoring/prometheus-json-exporter-master/
RUN go-wrapper download && go-wrapper install

CMD ["go-wrapper", "run", "--interval", "15"]

EXPOSE 7979