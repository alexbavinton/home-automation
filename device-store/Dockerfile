FROM golang:1.19 as build-stage

WORKDIR /app

COPY ./ ./
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build ./cmd/main.go 

FROM gcr.io/distroless/base-debian11 AS release-stage

COPY --from=build-stage /app/main main

ENV REDIS_HOST=redis
ENV REDIS_PORT=6379

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/main"]