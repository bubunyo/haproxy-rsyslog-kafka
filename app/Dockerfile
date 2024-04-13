FROM golang:1.20-alpine
WORKDIR /app
COPY ./go.* ./
COPY ./*.go ./
RUN go build -o ./main
ENTRYPOINT ["./main"]