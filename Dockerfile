# Build the Go Binary
FROM registry.suse.com/bci/golang:1.18 as build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 go build -o /dc-hello-world

# Create image to bundle app
# Using an image with tools, like a shell, rather than a scratch image to
# enable easier debugging.
FROM registry.suse.com/bci/bci-micro:15.3

COPY --from=build /dc-hello-world /usr/local/bin/dc-hello-world

EXPOSE 3000

CMD ["/usr/local/bin/dc-hello-world"]