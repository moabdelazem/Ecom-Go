FROM golang:1.23.1 as build-stage

    WORKDIR /app

    COPY go.mod go.sum ./

    RUN go mod download

    COPY . .

    RUN CGO_ENABLED=0 GOOS=linux go build -o /api ./cmd/main.go

# Run the tests in the container
FROM build-stage AS run-test-stage
    RUN go test -v ./...


FROM scratch AS build-release-stage

    WORKDIR /

    COPY --from=build-stage /api /api

    EXPOSE 8080

    ENTRYPOINT ["/api"]