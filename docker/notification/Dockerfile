# argument for Go version
ARG GO_VERSION=1.22.7

# STAGE 1: building the executable
FROM golang:${GO_VERSION} AS builder

WORKDIR /go/src/app
ADD . /go/src/app

RUN go clean -modcache 

RUN rm go.sum

RUN go get -d -v ./...

WORKDIR /go/src/app/services/notification/cmd

RUN CGO_ENABLED=0 go build -o /go/bin/app

# Now copy it into our base image.
FROM gcr.io/distroless/base-debian11

COPY --from=builder --chown=nonroot /go/bin/app /app

COPY --chown=nonroot:nonroot --chmod=777 ./services/notification/resources /services/notification/resources

USER nonroot:nonroot

CMD ["/app"]
