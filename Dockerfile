############################
# STEP 1 build executable binary
############################
FROM --platform=${BUILDPLATFORM} golang:1.15.2-alpine3.12 AS builder
# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/src/rest-apis/gorilla-rest-api/
ENV CGO_ENABLED=0
COPY . .
# Fetch dependencies.
# Using go get.
ARG TARGETOS
ARG TARGETARCH
RUN GOOS=${TARGETOS} GOARCH=${TARGETARCH} go get -d -v && go build -o /go/bin/gorilla-rest-api
# Build the binary.
# RUN go build -o /go/bin/gorilla-rest-api
############################
# STEP 2 build a small image
############################
FROM scratch
# Copy our static executable.
COPY --from=builder /go/bin/gorilla-rest-api /go/bin/gorilla-rest-api

EXPOSE 5000
# Run the hello binary.
ENTRYPOINT ["/go/bin/gorilla-rest-api"]