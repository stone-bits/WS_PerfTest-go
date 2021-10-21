# Build stage
FROM golang:1.15-alpine AS build

ENV PROJECT_DIR WS_PerfTest-go

ADD . /$PROJECT_DIR
WORKDIR /$PROJECT_DIR
RUN go build -o /app/WS_PerfTest-go

# Final stage
FROM golang:1.15-alpine

COPY --from=build /app /app
WORKDIR /app

ENTRYPOINT [ "./WS_PerfTest-go" ]