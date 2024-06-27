FROM golang:1.21 as base

WORKDIR /app
COPY . /app
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build cmd/api/main.go

FROM alpine
COPY go.mod go.sum ./
WORKDIR /app
COPY --from=base /app/main ./
EXPOSE 8080
CMD [ "./main" ]