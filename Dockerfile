FROM golang:1.25-alpine

# set working directory
WORKDIR /the/workdir/path

# copy go mod and sum files
COPY go.mod go.sum ./
RUN go mod download

# copy the source code
COPY . .

#expose port
EXPOSE 8000

#build the Go app
RUN go build -o main ./cmd

#Run the executable
CMD ["./main"]