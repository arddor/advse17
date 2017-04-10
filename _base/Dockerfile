########### DEV ###########

#FROM golang:1.8-alpine
# beego needs the normal version
FROM golang:1.8

# Install the bee dev tool
RUN go get github.com/beego/bee

# Install dependencies
# from ALL containers

# https://github.com/gin-gonic/gin
RUN go get github.com/gin-gonic/gin

# https://github.com/cdipaolo/sentiment
RUN go get github.com/cdipaolo/sentiment

# https://labix.org/mgo
RUN go get gopkg.in/mgo.v2

# https://github.com/dghubble/go-twitter/twitter
RUN go get github.com/dghubble/go-twitter/twitter

# https://github.com/dghubble/oauth1
RUN go get github.com/dghubble/oauth1

# https://github.com/streadway/amqp - RabbitMQ
RUN go get github.com/streadway/amqp

# Expose the application on port 8080
EXPOSE 8080

# Set the entry point
# Bee runs the app and watches for changes
CMD ["bee", "run"]


# 
# https://semaphoreci.com/community/tutorials/how-to-deploy-a-go-web-application-with-docker
# 
#docker build -t base_image .
#docker run -it --rm --name ma_instance -p 8080:8080 -v ./compute:/go/src/compute -w /go/src/compute base_image