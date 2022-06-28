FROM code.pztrn.name/containers/mirror/golang:1.18.3-alpine AS build

WORKDIR /go/src/go.dev.pztrn.name/periodicator
COPY . .

ENV CGO_ENABLED=0
RUN apk add make && make build

FROM code.pztrn.name/containers/mirror/alpine:3.16.0
LABEL maintainer="Stanislav N. <pztrn@pztrn.name>"

COPY --from=build /go/src/go.dev.pztrn.name/periodicator/periodicator /usr/local/bin/periodicator

RUN apk add tzdata

ENV GPT_CONFIG=/periodicator.yaml

ENTRYPOINT [ "/usr/local/bin/periodicator" ]
