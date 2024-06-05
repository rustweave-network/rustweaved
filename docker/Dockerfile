# -- multistage docker build: stage #1: build stage
FROM golang:1.19-alpine AS build

RUN mkdir -p /go/src/github.com/rustweave-network/rustweaved

WORKDIR /go/src/github.com/rustweave-network/rustweaved

RUN apk add --no-cache curl git openssh binutils gcc musl-dev

COPY go.mod .
COPY go.sum .


# Cache rustweaved dependencies
RUN go mod download

COPY . .

RUN go build $FLAGS -o rustweaved .

# --- multistage docker build: stage #2: runtime image
FROM alpine
WORKDIR /app

RUN apk add --no-cache ca-certificates tini

COPY --from=build /go/src/github.com/rustweave-network/rustweaved/rustweaved /app/
COPY --from=build /go/src/github.com/rustweave-network/rustweaved/infrastructure/config/sample-rustweaved.conf /app/

USER nobody
ENTRYPOINT [ "/sbin/tini", "--" ]
