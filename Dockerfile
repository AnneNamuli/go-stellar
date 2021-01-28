FROM golang:1.14
WORKDIR /go/src/github.com/AnneNamuli/go-stellar

COPY . .
ENV GO111MODULE=on
RUN go install github.com/AnneNamuli/go-stellar/tools/...
RUN go install github.com/AnneNamuli/go-stellar/services/...
