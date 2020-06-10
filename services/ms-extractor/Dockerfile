FROM golang:latest as dev
LABEL maintainer="Ardit Sinani <arditsinani04@gmail.com>"

# copy the local package files to the container workspace
ADD . /go/src/mom/services/ms-extractor

# Setting up working directory
WORKDIR /go/src/mom/services/ms-extractor

# Install dependencies
COPY go.mod .
COPY go.sum .
RUN go mod download

# build binary
#RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ms-consumer .

# Listen to directory changes
RUN go get github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon --build="go build ./main.go" --command=./main