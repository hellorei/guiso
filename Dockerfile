# Arg
ARG GOARCH=amd64
ARG GOOS=linux
ARG GO_VERSION=1.22
ARG BINARY_NAME=guiso-app

# Build stage
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Install build dependencies
RUN apk add --no-cache git make

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Install all dependencies
RUN make install

# Build the Go application
RUN make build

# Final stage
FROM alpine:3.20

WORKDIR /app

ARG BINARY_NAME
ARG STRIPE_URL
ARG STRIPE_URL_1
ARG STRIPE_URL_2
ARG STRIPE_URL_3
ARG STRIPE_URL_SUB_1
ARG STRIPE_URL_SUB_2
ARG STRIPE_URL_SUB_3
ARG STRIPE_URL_SUB_MANAGE
ARG GA4_MEASUREMENT_ID
ARG GA4_API_SECRET


# Environment variables
ENV BINARY_NAME=${BINARY_NAME}
ENV PORT=8080
ENV STRIPE_URL=${STRIPE_URL}
ENV STRIPE_URL_1=${STRIPE_URL_1}
ENV STRIPE_URL_2=${STRIPE_URL_2}
ENV STRIPE_URL_3=${STRIPE_URL_3}
ENV STRIPE_URL_SUB_1=${STRIPE_URL_SUB_1}
ENV STRIPE_URL_SUB_2=${STRIPE_URL_SUB_2}
ENV STRIPE_URL_SUB_3=${STRIPE_URL_SUB_3}
ENV STRIPE_URL_SUB_MANAGE=${STRIPE_URL_SUB_MANAGE}
ENV GA4_MEASUREMENT_ID=${GA4_MEASUREMENT_ID}
ENV GA4_API_SECRET=${GA4_API_SECRET}

# Copy the binary and public files from the builder stage
COPY --from=builder /app/bin/${BINARY_NAME} ./${BINARY_NAME}
COPY --from=builder /app/public ./public
COPY --from=builder /app/locales ./locales

# Expose necessary ports (optional, depending on your app)
EXPOSE ${PORT}

# Run the application
CMD ["/bin/sh", "-c", "./${BINARY_NAME}"]