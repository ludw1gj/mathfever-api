### STAGE 1: Build
FROM golang:1.15.3-alpine3.12 AS build
WORKDIR /go/src/app
COPY . .
RUN go mod download
RUN go build -o mathfever-api-server .

### STAGE 2: Run
FROM alpine:3.12.0
RUN mkdir /app
COPY --from=build /go/src/app/mathfever-api-server /app
CMD ["/app/mathfever-api-server"]
EXPOSE 8000
