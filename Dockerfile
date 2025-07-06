FROM golang:1.23
WORKDIR /API
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build
RUN mkdir -p /API/Files/Images
CMD ["./jorycia_api"]
