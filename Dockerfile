## Base image to have a container with Go version necessary to run
FROM golang:1.23

## Within image, /app is the directory to hold source code
WORKDIR /app

## Copy directly from source (our computer) to image . directory
COPY . .

## Install dependencies for application (think npm install)
RUN go mod download

## Build the app
RUN go build -o /todo

## Since our server is running on port 8080, we need to expose this port on the container
EXPOSE 8080

## Run the executable
CMD ["/todo"]