# syntax=docker/dockerfile:1

FROM golang:1.17 AS builder

ADD . /go/src/authapp
WORKDIR /go/src/authapp

COPY . /go/src/authapp

RUN go mod download
RUN CGO_ENABLED=1 GOOS=linux go build -o /app -a -ldflags '-linkmode external -extldflags "-static"' .

# RUN go build -o app .

FROM scratch

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage. Observe we also copied the .env file
COPY --from=builder /app .

# Expose port 8000 to the outside world
EXPOSE 8080

#Command to run the executable
CMD [ "./app" ]