FROM golang:1.23
WORKDIR /API
COPY . .
RUN go mod tidy
RUN go mod download
RUN mkdir -p /API/Files/Images
RUN go build
CMD ["./jorycia_api"]
