FROM golang:alpine as builder

WORKDIR /app

COPY . .

# Add "-tags production" if you are using embedded configuration
RUN go build -ldflags "-w -s"

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/docker-goyave ./docker-goyave

# Not needed if you are using embed resources
COPY resources /app/resources 

# Not needed if you are using embed configuration
COPY config.production.json ./config.production.json

RUN useradd -r -U go-exec -M -d /app -s /bin/false
RUN chown -R go-exec /app

USER go-exec

EXPOSE 80
ENV GOYAVE_ENV=production

ENTRYPOINT [ "./docker-goyave" ]