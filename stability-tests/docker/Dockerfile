ARG KARLSEND_IMAGE
ARG KARLSENMINER_IMAGE

FROM ${KARLSEND_IMAGE} as karlsend
FROM ${KARLSENMINER_IMAGE} as karlsenminer

FROM golang:1.24-alpine

RUN mkdir -p /go/src/github.com/karlsen-network/karlsend/

WORKDIR /go/src/github.com/karlsen-network/karlsend/

RUN apk add bash build-base git

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

COPY --from=karlsend /app/ /app/
COPY --from=karlsenminer /app/ /app/
ENV PATH="/app:${PATH}"

WORKDIR /go/src/github.com/karlsen-network/karlsend/stability-tests

RUN go install ./...

ENTRYPOINT ["./run/run.sh"]
