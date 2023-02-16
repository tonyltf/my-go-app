FROM golang:1.17-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./
RUN go build -o /my-go-app

EXPOSE 3333

CMD [ "/my-go-app" ]
