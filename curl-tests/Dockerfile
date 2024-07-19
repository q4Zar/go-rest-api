# Use an official lightweight image of Alpine Linux as a base
FROM alpine:latest

# Install curl and jq
RUN apk --no-cache add curl jq

# Set a working directory
WORKDIR /usr/src/curl-tests

# Copy your script into the container
COPY . .

# Make your script executable
RUN chmod +x *.sh

# Run your script by default when the container starts
CMD ["sh", "./run.sh", "http://go-api:8080"]