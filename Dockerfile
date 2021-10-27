# the base image to use for the container. The golang image is a Linux based image which has golang installed and no other additional programs or software
FROM golang:alpine 

RUN mkdir /app

# changes the working directory and sets it for subsequent commands
WORKDIR /app

# ADD copies the file from one location to another
# we want to copy go.mod and go.sum first so we will have all the libraries installed before copying all the files
ADD go.mod .
ADD go.sum .

# RUN will execute any commands in a new layer on top of the current image and commit the results
# the resulting committed image will be used for the next step in the Dockerfile
RUN go mod download
ADD . .

RUN go get github.com/githubnemo/CompileDaemon

# instructs that services running on Docker container is available in port 8000
EXPOSE 8000

# runs the command inside the container once a container is created from an image. You can only have one entrypoint instruction in a Dockerfile
# if multiple ENTRYPOINT instructions are used, the last one will be executed
ENTRYPOINT CompileDaemon --build="go build main.go" --command=./main