FROM golang:1.21-alpine AS build

RUN apk add --no-cache git make curl ca-certificates

WORKDIR /app
ADD . .

RUN chmod +x /bin/sh

RUN make mod && make build

FROM alpine
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /app/bin/re-test /app/re-test
RUN chmod +x /bin/sh
RUN chmod +x /app/re-test
ENTRYPOINT ["/bin/sh", "/app/re-test" ]
