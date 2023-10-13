FROM golang:1.21-alpine AS build

RUN apk add --no-cache git make curl ca-certificates

WORKDIR /app
ADD . .

RUN make mod && make build

FROM alpine
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /app/bin/re-app /app/re-app
ENTRYPOINT [ "/app/re-app" ]
