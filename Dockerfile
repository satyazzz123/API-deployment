FROM golang:1.20
WORKDIR /app
COPY . .
EXPOSE 5000
CMD ["go","run","main.go"]
