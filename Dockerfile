FROM golang:1.19 as builder

WORKDIR /task
COPY . .
RUN go build -o main cmd/main.go

EXPOSE 8080
CMD [ "/task/main" ]