ARG RUSTWEAVED_IMAGE
ARG RUSTWEAVEMINER_IMAGE

FROM ${RUSTWEAVED_IMAGE} as rustweaved
FROM ${RUSTWEAVEMINER_IMAGE} as rustweaveminer

FROM golang:1.19-alpine

RUN mkdir -p /go/src/github.com/rustweave-network/rustweaved

WORKDIR /go/src/github.com/rustweave-network/rustweaved

RUN apk add bash build-base git

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

COPY --from=rustweaved /app/ /app/
COPY --from=rustweaveminer /app/ /app/
ENV PATH="/app:${PATH}"

WORKDIR /go/src/github.com/rustweave-network/rustweaved/stability-tests

RUN go install ./...

ENTRYPOINT ["./run/run.sh"]