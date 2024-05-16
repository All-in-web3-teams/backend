# Use an official Go runtime as a parent image
FROM golang:1.22.2

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

#代理
RUN go env -w GOPROXY=https://goproxy.cn

RUN go mod tidy


# Command to run the executable
CMD ["go", "run", "main.go"]

