FROM golang:1.14

COPY main main
COPY data data
CMD ["./main"]
