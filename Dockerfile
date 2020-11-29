FROM golang:1.15.3
RUN ls

SHELL ["/bin/bash", "-c"]
WORKDIR /go-cors
COPY . .

RUN go build

CMD ./go-cors
