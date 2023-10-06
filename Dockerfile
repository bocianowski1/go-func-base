# FROM alpine:3.16 as root-certs
# RUN apk add -U --no-cache ca-certificates
# RUN addgroup -g 1001 app
# RUN adduser app -u 1001 -D -G app /home/app

# FROM golang:1.20 as builder
# WORKDIR /app
# COPY --from=root-certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
# COPY . .
# RUN CGO_ENABLED=0 GOOS=linux go build -mod=vendor -o app .

# FROM scratch as final
# COPY --from=root-certs /etc/passwd /etc/passwd
# COPY --from=root-certs /etc/group /etc/group
# COPY --chown=1001:1001 --from=root-certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
# COPY --chown=1001:1001 --from=builder /app/app /app
# USER app
# ENTRYPOINT ["/app/app"]

FROM mongo:latest as mongo
ENV MONGO_DATA_DIR=/data/db
RUN mkdir -p $MONGO_DATA_DIR
EXPOSE 27017
CMD ["mongod"]