FROM golang:alpine as builder
WORKDIR /app
COPY src .
RUN go build -o graphpg main.go

FROM alpine

ENV HOST http://localhost:9000/graphql
ENV TITLE "GraphQL Playground"
ENV PORT 8080
ENV THEME dark

WORKDIR /app
COPY --from=builder /app/graphpg .
COPY --from=builder /app/static src/static
COPY --from=builder /app/index.tmpl src/index.tmpl
EXPOSE 8080
ENTRYPOINT  ["./graphpg"]