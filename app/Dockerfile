FROM alpine:latest
WORKDIR /app
COPY main .
# Nếu có file .env hoặc thư mục config thì copy kèm vào
# COPY .env .
EXPOSE 8080
CMD ["./main"]
# RUN CGO_ENABLED=0 GOOS=linux go build -p 1 -o main .