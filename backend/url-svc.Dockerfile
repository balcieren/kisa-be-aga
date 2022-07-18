# Build stage
FROM golang:1.18-alpine as builder
WORKDIR /src

COPY go.sum go.mod ./
RUN go mod download

ADD . . 
RUN CGO_ENABLED=0 go build -o /bin/app cmd/url-service/main.go

# Execute stage
FROM scratch
WORKDIR /build
COPY --from=builder /bin/app /bin/app
ENTRYPOINT ["/bin/app"]