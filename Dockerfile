FROM golang:1.20-alpine
WORKDIR /app
COPY . .
RUN go build -o main .
EXPOSE 5000
CMD ["./main"]
