# syntax=docker/dockerfile:1

FROM golang


WORKDIR /app


COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o main main.go

# Expose the port the app runs on
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
