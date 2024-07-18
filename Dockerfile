FROM golang:alpine
WORKDIR /app
COPY app/* .
RUN go mod download
EXPOSE 8080
CMD ["go", "run", "."]
# RUN go build -o ./build/go-rest-api
# ENTRYPOINT ["./build/go-rest-api"]