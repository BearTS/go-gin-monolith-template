FROM golang:alpine
RUN apk --no-cache add tzdata
ENV TZ=Asia/Kolkata
RUN mkdir /app
ADD . /app/
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
RUN go build -o main .

# From here copy the built executable to the second container
FROM scratch
COPY --from=builder /app/main /app/
RUN adduser -S -D -H -h /app appuser
USER appuser
CMD ["./main"]
