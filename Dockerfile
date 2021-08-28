FROM golang:1.14.2-alpine3.11 as builder

RUN adduser -D -g '' appuser

LABEL Maintainer="Koala DevTeam <technology@koala.id>"

#RUN apk add tzdata
RUN apk update && apk add --no-cache git ca-certificates && update-ca-certificates

WORKDIR /apps/

COPY go.* ./
RUN go mod download

COPY . .
# RUN mv files/ /

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o /go/bin/app .
# RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o /go/bin/cli cli/*

FROM alpine:latest

RUN ls
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /go/bin/app /go/bin/app
# COPY --from=builder /apps/staging.config.json /staging.config.json
# COPY --from=builder /apps/production.config.json /production.config.json
# COPY --from=builder /files /files

USER appuser

EXPOSE 8080
# CMD ["/go/bin/koala-listing-api-auth"]