FROM golang:1.19-alpine
ENV GOPROXY=https://goproxy.cn,direct 

WORKDIR /app
COPY . .
RUN go mod download && go build -o docker-app
EXPOSE 8080

CMD [ "./docker-app" ]