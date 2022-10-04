FROM golang:latest

MAINTAINER ReaTion "reation_11@163.com"

WORKDIR /home_user_service

COPY . /home_user_service

RUN go build user.go

EXPOSE 8080

RUN chmod +x /home_user_service/order

CMD ["/home_user_service/order -f /home_user_service/etc/produce.yaml"]
