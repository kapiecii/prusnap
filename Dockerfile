FROM golang:1.24-bullseye

WORKDIR /app

COPY main.go .
COPY static/ ./static/
COPY templates/ ./templates/

RUN mkdir -p /app/pictures

RUN go build -o photo-viewer main.go

EXPOSE 8080

CMD ["./photo-viewer"]
