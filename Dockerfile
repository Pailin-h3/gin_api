# Use a super tiny base image
FROM alpine:latest

# Install security certs (needed if your API talks to other HTTPS services)
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the binary that Jenkins built into this image
COPY api-binary .

# Give the binary permission to run
RUN chmod +x api-binary

# Tell Docker to run your API when the container starts
CMD ["./api-binary"]