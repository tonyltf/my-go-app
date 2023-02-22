FROM golang:1.17-alpine as build

WORKDIR /app

RUN apk add build-base
COPY go.mod go.sum Makefile ./
RUN go mod download
# COPY api cmd configs docs init internal ./app/
COPY . /app/
WORKDIR /app/cmd/exchange-rate
# RUN go get -u 
RUN go get -u github.com/swaggo/swag/cmd/swag
RUN swag init -g main.go
RUN go build 
# RUN CGO_ENABLED=0 GOOS=linux go build -v -a -installsuffix cgo -o swag cmd/swag/main.go



FROM golang:1.17-alpine 
# COPY Makefile ./
COPY --from=build /app/cmd/exchange-rate .
COPY --from=build /app/cmd/exchange-rate /app/cmd/exchange-rate
COPY --from=build /app/api /app/api 
COPY --from=build /app/configs /app/configs
COPY --from=build /app/docs /app/docs 
# COPY --from=build /app/swagger /app/swagger 
COPY --from=build /app/init/ /app/init 
COPY --from=build /app/internal /app/internal 
# WORKDIR /app/cmd/exchange-rate
EXPOSE 80

ENTRYPOINT [ "./exchange-rate" ]
