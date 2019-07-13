# From https://medium.com/@pierreprinetti/the-go-dockerfile-d5d43af9ee3c
FROM golang:1.12.6 AS builder

# Download and install the latest release of dep
ADD https://github.com/golang/dep/releases/download/v0.5.0/dep-linux-amd64 /usr/bin/dep
RUN chmod +x /usr/bin/dep

# Copy the code from the host and compile it
WORKDIR /go/src/bitbucket.org/arthemismc/core
COPY Gopkg.toml Gopkg.lock ./
RUN dep ensure --vendor-only
COPY . ./
RUN CGO_ENABLED=0 make

FROM scratch
COPY --from=builder /go/src/bitbucket.org/arthemismc/core/output/core ./
ENTRYPOINT ["./core"]

