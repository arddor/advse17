# ComputeContainer
# 
# This is where the Sentiment Analysis is happening
# 
# The Image containes the compiled code, launches are way faster that way.
# Required image rebuilt on dev though!

FROM golang:1.8

# Setting up working env
RUN mkdir -p $GOPATH/src/compute
WORKDIR $GOPATH/src/compute

# Copy our code from the current folder
ADD . $GOPATH/src/compute

# Default Go Port
EXPOSE 8080

# On container build get dependencies, build and...
RUN go get -d
RUN go build

# ...run on container launch
CMD ["go", "run", "main.go"]
