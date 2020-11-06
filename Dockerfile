FROM golang:1.14 AS builder

WORKDIR /src
COPY go.mod go.mod
COPY go.sum go.sum
COPY pkg/ pkg/
COPY vendor/ vendor/
COPY data/ data/
COPY docs/ docs/
COPY cmd/ cmd/
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -mod=vendor -a -installsuffix cgo -o third_rail cmd/third_rail/main.go
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -mod=vendor -a -installsuffix cgo -o dbinit cmd/dbinit/main.go

FROM alpine:3.11
RUN apk --no-cache add ca-certificates

COPY --from=builder /src/third_rail /bin/third_rail
COPY --from=builder /src/dbinit /bin/dbinit
CMD ["/bin/third_rail"]
