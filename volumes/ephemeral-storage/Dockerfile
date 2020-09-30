FROM golang:latest as builder
LABEL maintainer="Jason Ody <jasonody@users.noreply.github.com>"
WORKDIR /go
COPY ./src .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o kvstore .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root
COPY --from=builder /go/kvstore .
EXPOSE 8000
CMD ["./kvstore"]