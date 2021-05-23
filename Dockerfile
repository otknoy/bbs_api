FROM golang:1.16.4 as builder

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY openapi/ openapi/
COPY service/ service/
COPY infra/ infra/
COPY main.go .
RUN CGO_ENABLED=0 go build -o bbs-api

FROM scratch
COPY --from=builder /app/bbs-api /bin/bbs-api
ENTRYPOINT ["/bin/bbs-api"]
