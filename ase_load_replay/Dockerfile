FROM golang:1.8-alpine

# Install git
RUN apk add --update \
        git \
    ; \
    export PATH=/usr/bin/git:"${PATH}"; \
    rm -rf /var/cache/apk/*

# Copy are own sources in
ADD ./ /go/src/github.com/arddor/advse17/ase_load_replay

# Get all the great dependencies
RUN go get github.com/arddor/advse17/ase_load_replay

# Install sources
RUN go install github.com/arddor/advse17/ase_load_replay

# Remove git
RUN apk del git

# Set the entry point
ENTRYPOINT /go/bin/ase_load_replay
