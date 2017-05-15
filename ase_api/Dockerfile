########### DEV ###########

#FROM golang:1.8-alpine
# Can't use alpine, need git
FROM golang:1.8

# https://github.com/gin-gonic/gin
RUN go get github.com/gin-gonic/gin

# https://gopkg.in/gorethink/gorethink.v3
RUN go get gopkg.in/gorethink/gorethink.v3

# Expose the application on port 8080
EXPOSE 8080

# Set the entry point
CMD ["go", "run", "main.go"]
