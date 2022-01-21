FROM golang:1.17 as builder
WORKDIR /build
COPY *.go go.mod go.sum ./
RUN go get .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-extldflags '-static'"

FROM scratch
COPY --from=builder /build/wordgame .
COPY words.txt .
ENTRYPOINT ["/wordgame"]
