FROM golang:latest

ENV GO111MODULE=on

WORKDIR /app

RUN go get github.com/golang-migrate/migrate/v4/database/postgres

#COPY ENV.env .
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY ENV.env .
COPY . .

RUN go build -o main .
 
# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]