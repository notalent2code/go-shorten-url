# Stage 1: build the app
FROM golang:alpine AS build

RUN mkdir /build

ADD . /build/

WORKDIR /build

RUN go build -o main .

# Stage 2: deploy the app
FROM alpine

RUN adduser -S -D -H -h /app appuser

USER appuser

COPY . /app

COPY --from=build /build/main /app/

WORKDIR /app

EXPOSE 4000

CMD ["./main"]