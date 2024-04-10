FROM golang:1.20-alpine
WORKDIR /app
COPY ./app/go.* ./
COPY ./app/*.go ./
RUN go build -o ./main
ENTRYPOINT ["./main"]