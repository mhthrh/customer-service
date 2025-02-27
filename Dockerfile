FROM golang:1.23


LABEL maintainer="Mohsen taheri m.rozbehani@outlook.com"
LABEL version="v1.0.1"
LABEL description="com.x-bank.customer-service"

RUN apt update && apt install vim -y && apt install make -y && apt-get update && apt-get install -y bash

ENV APP_ENV=test
ENV PORT=6985
ENV IP="0.0.0.0"

EXPOSE 6985
WORKDIR app/x-bank/customer-service

COPY . .

SHELL ["/bin/bash", "-euxo", "pipefail", "-c"]

#RUN cd ./x-bank/customer-service && go mod download  && go mod verify   --modfile=x-bank/customer-service/mod.go
RUN  go mod download  && go mod verify
RUN  go build -o customer-service cli/main.go
#debug
RUN ["/bin/bash", "-c", "pwd"]
RUN ["/bin/bash", "-c", "ls -al"]
#
RUN ["make", "set-app-path-docker"]
CMD ["./customer-service"]