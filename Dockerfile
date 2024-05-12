FROM golang:latest

RUN mkdir /app
ADD . /app
WORKDIR /app

# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum
# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/engine/reference/builder/#copy
COPY *.go ./

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o main_enigmanations .

# Expose port 8080
EXPOSE 8080

ENV ENV=production

# Run Executable
CMD [ "./main_enigmanations" ]
