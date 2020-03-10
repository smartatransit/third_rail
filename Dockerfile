FROM golang:1.13

COPY third_rail third_rail

CMD ["./third_rail"]
