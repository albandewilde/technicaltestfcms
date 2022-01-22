FROM golang:1.17 as builder

WORKDIR /src/technicaltestfcms

COPY . .

RUN CGO_ENABLED=0 go build -o ./out/server


FROM alpine

WORKDIR /bin/technicaltestfcms

COPY --from=builder /src/technicaltestfcms/out/server .

EXPOSE 8254

CMD ["./server"]
