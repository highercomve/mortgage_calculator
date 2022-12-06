FROM golang:1.19-alpine as builder

WORKDIR /app

COPY . .

RUN apk update; apk add -U --no-cache \
    git \
    curl \
    build-base

RUN go get -d -v ./... \
    && go install -v ./...

FROM alpine

RUN apk update; apk add ca-certificates

COPY --from=builder /go/bin/mortgage_calculator /usr/bin/

ENTRYPOINT [ "/usr/bin/mortgage_calculator" ]
