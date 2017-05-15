########### DEV ###########

#FROM golang:1.8-alpine
# Can't use alpine, need git
FROM golang:1.8

# https://github.com/dghubble/go-twitter/twitter
RUN go get github.com/dghubble/go-twitter/twitter

# https://github.com/dghubble/oauth1
RUN go get github.com/dghubble/oauth1

# https://gopkg.in/gorethink/gorethink.v3
RUN go get gopkg.in/gorethink/gorethink.v3

# https://github.com/streadway/amqp - RabbitMQ
RUN go get github.com/streadway/amqp

# Expose the application on port 8080
EXPOSE 8080

# Set the entry point
CMD ["go", "run", "main.go"]
