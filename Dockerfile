# Multi-stage build for Starbucks Coffee Shop

# Stage 1: Build frontend
FROM node:20-alpine AS frontend-builder
WORKDIR /app/frontend
COPY frontend/package*.json ./
RUN npm ci
COPY frontend/ ./
RUN npm run build

# Stage 2: Build backend
FROM golang:1.25-alpine AS backend-builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o starbucks .

# Stage 3: Production image
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=backend-builder /app/starbucks .
COPY --from=frontend-builder /app/frontend/dist ./frontend/dist
EXPOSE 8080
CMD ["./starbucks"]
