FROM golang:1.21.0-alpine

RUN mkdir /app
ADD . /app/
WORKDIR /app

RUN apk update && apk add ca-certificates && apk add tzdata

RUN go build -buildvcs=false -o app ./cmd/app

FROM alpine:3.17.2

WORKDIR /app/

RUN apk update && apk add ca-certificates && apk add tzdata

COPY --from=0  /app/app /app/app
COPY --from=0  /app/configs /app/configs/

EXPOSE 8000

ENTRYPOINT /app/app
