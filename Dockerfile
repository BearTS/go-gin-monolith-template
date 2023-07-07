# Build stage
FROM golang:alpine as builder
WORKDIR /app
COPY . .
RUN go build -o main main.go

# Run stage
FROM alpine
RUN apk --no-cache add tzdata
ENV TZ=Asia/Kolkata
WORKDIR /app
COPY --from=builder /app/main .
COPY .env .
CMD [ "/app/main" ]