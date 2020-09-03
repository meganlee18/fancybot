FROM golang:latest

LABEL maintainer="Megan Lee"

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/github.com/zendesk/lab-day/fancybot

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./internal/script/start.sh"]
