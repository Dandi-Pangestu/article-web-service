FROM golang:1.15-alpine
ARG SERVICE_NAME

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

WORKDIR /app
RUN mkdir -p $SERVICE_NAME
RUN cd $SERVICE_NAME
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main ./cmd/$SERVICE_NAME

EXPOSE 8080

CMD ["./main"]
