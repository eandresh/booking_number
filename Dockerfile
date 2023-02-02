# build stage
FROM golang:1.19.1-alpine as builder
LABEL "com.eheredia.vendor"="Eheredia"
LABEL "com.eheredia.maintainer"=""
LABEL "version"="2023.0.1"

RUN apk update \
  && apk add bash ca-certificates git openssh gcc g++ curl libc6-compat libc-dev make pkgconf

#RUN mkdir -p -m 0600 ~/.ssh && ssh-keyscan github.com >> ~/.ssh/known_hosts

WORKDIR /go/src/digital_shift
# Copy all the Code and stuff to compile everything
COPY go.mod go.sum ./

RUN go mod download -x

# Copy all the Code and stuff to compile everything
COPY . .
# run test
RUN go test ./...
# Builds the application as a statically linked one, to allow it to run on alpine
RUN CGO_ENABLED=1 && GOOS=linux && GOARCH=amd64 && go build -tags musl,appsec  -o compiled-app ./main

# # Moving the binary to the 'final Image' to make it smaller
FROM alpine:latest

# # `service` should be replaced here as well
COPY --from=builder /go/src/digital_shift/compiled-app .
#COPY --from=builder /go/src/digital_shift/pkg/i18n pkg/i18n

ENV SERVICE="eh-digital-shift"

CMD ["./compiled-app"]
