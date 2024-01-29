# Stage 1: Copy only the necessary dependency files
FROM golang:alpine as deps

WORKDIR /app
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download \
    && go mod tidy

# Stage 2: Copy the entire project
FROM alpine:3.14 as base

WORKDIR /app
COPY . .

# Stage 3: Build the frontend
FROM node:20 as frontend

WORKDIR /app/web
COPY --from=base /app/web/ ./

# Combine Node.js commands
RUN npm install \
    && npm run clean \
    && npm run build

# Stage 4: Build the Go app
FROM golang:alpine as backend

WORKDIR /app
COPY --from=base /app ./

# Copy only the necessary artifacts from the frontend build stage
COPY --from=frontend /app/web/build /app/web/build

RUN go build -o /app/moonlogs

# Final Stage
FROM alpine:3.14

WORKDIR /app
COPY --from=backend /app/moonlogs ./

ENTRYPOINT ["./moonlogs"]
CMD []