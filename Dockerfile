# Stage 1: Copy the entire project
FROM alpine:latest as base

WORKDIR /app
COPY . .

# Stage 2: Build the frontend
FROM node:latest as frontend

WORKDIR /app/web
COPY --from=base /app/web/ ./
RUN npm install
RUN npm run clean
RUN npm run build

# Stage 3: Build the Go app
FROM golang:alpine as backend

WORKDIR /app
COPY --from=base /app ./

# Copy only the necessary artifacts from the frontend build stage
COPY --from=frontend /app/web/build /app/web/build

RUN go build -o /app/moonlogs

# Final Stage
FROM alpine:latest

WORKDIR /app
COPY --from=backend /app/moonlogs ./

CMD ["./moonlogs"]