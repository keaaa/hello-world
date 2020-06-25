FROM golang AS builder

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

COPY . .
RUN go build -o main .


FROM scratch AS runner
COPY --from=builder /build/main /

EXPOSE 8000
ENTRYPOINT ["/main"]