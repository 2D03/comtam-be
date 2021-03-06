FROM golang:latest
RUN mkdir /app
WORKDIR /app

ARG env
ARG config
ARG version
ENV env=${env}
ENV version=${version}
ENV config=${config}

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 8000

CMD ["./main"]
