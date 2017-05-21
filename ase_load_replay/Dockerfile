#FROM golang:1.8-alpine
# Can't use alpine, need git
FROM golang:1.8

# Copy are own sources in
ADD ./ /go/src/github.com/arddor/advse17/ase_load_replay

# Get all the great dependencies
RUN go get github.com/arddor/advse17/ase_load_replay

# Install sources
RUN go install github.com/arddor/advse17/ase_load_replay

# Set the entry point
ENTRYPOINT /go/bin/ase_load_replay
