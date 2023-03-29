# build SERVER
FROM golang:1.20.2-alpine3.16
WORKDIR /server
COPY go.* ./
RUN go mod download
COPY *.go ./
RUN go build -o /notes-server
EXPOSE 8080
CMD ["/notes-server"]
