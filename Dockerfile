# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang

MAINTAINER Joël Vimenet

# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/joelvim/sensit

# Workdir is the projects dir
WORKDIR /go/src/github.com/joelvim/sensit

# Install godep
RUN go get github.com/tools/godep

# Install dependencies
RUN godep restore

# Build the sensit command inside the container.
RUN godep go install github.com/joelvim/sensit

# Run the sensit daemon command by default when the container starts.
ENTRYPOINT /go/bin/sensit daemon

# Document that the service listens on port 8080.
EXPOSE 8080
