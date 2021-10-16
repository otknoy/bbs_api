FROM golang:1.17.2 as builder

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY openapi/ openapi/
COPY domain/ domain/
COPY service/ service/
COPY interfaces/ interfaces/
COPY infra/ infra/
COPY main.go .
RUN CGO_ENABLED=0 go build -o bbs-api

FROM scratch
COPY --from=builder /app/bbs-api /bin/bbs-api
ENTRYPOINT ["/bin/bbs-api"]
