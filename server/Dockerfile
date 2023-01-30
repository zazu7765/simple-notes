# build SERVER
FROM golang:1.19-alpine
WORKDIR /server
COPY go.* ./
RUN go mod download
COPY *.go ./
RUN go build -o /notes-server
EXPOSE 8080
CMD ["/notes-server"]
