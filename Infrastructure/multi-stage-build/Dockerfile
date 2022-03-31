FROM golang:1.18 AS build

WORKDIR /golang

COPY main.go ./
COPY go.mod ./

RUN go build -o /app

FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /app /app

USER nonroot:nonroot

EXPOSE 8080

ENTRYPOINT [ "/app" ]