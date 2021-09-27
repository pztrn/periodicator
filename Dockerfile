FROM golang:1.17.1-alpine AS build

WORKDIR /go/src/go.dev.pztrn.name/periodicator
COPY . .

ENV CGO_ENABLED=0
RUN go build -o periodicator .

FROM alpine:latest
LABEL maintainer="Stanislav N. <pztrn@pztrn.name>"

COPY --from=build /go/src/go.dev.pztrn.name/periodicator/periodicator /usr/local/bin/periodicator

RUN apk add tzdata

ENV GPT_CONFIG=/periodicator.yaml

ENTRYPOINT [ "/usr/local/bin/periodicator" ]
