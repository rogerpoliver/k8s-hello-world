FROM golang:1.24

# Set the Current Working Directory inside the container
WORKDIR /app

# Initialize a new Go module
RUN go mod init hello-go

# Copy the source code into the container
COPY . .

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod tidy

# Build the Go app
RUN go build -o server .

# Command to run the executable
CMD ["./server"]