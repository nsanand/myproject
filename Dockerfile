# Create build stage based on buster image
FROM golang:1.20-buster AS builder
LABEL authors="anand"

# Create working directory under /code
WORKDIR /go/src/app

# Copy over all go config
# Copy over Go source code
COPY . /go/src/app

# Install any required modules
RUN go mod download


#RUN export GOPATH=$HOME/go
#RUN export PATH=$PATH:$GOPATH/bin

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /myproject
#RUN go build -o main .

# Run the tests in the container
FROM builder AS run-test-stage
RUN go test -v ./...

# Expose the port the HTTP server is using
#EXPOSE 8080

# Run the app binary when we run the container
ENTRYPOINT ["/myproject", "20"]
#CMD ["./main", "20"]

