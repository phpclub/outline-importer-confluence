# Build Golang binary
FROM golang:1.22 AS build-golang

WORKDIR /srv/go/github.com/phpclub/outline-importer-confluence

COPY . .
RUN go get -v && go build -v -o /usr/local/bin/outline-importer-confluence

EXPOSE 8081
CMD ["outline-importer-confluence"]
