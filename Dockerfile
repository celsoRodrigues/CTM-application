FROM golang:alpine AS builder
LABEL maintainer="celso.rodrigues@hotmail.co.uk"

WORKDIR /go/src/celso
COPY . .
RUN go mod init test && go mod tidy && CGO_ENABLED=0 GOOS=linux go build -a --installsuffix ctm -o webserver

FROM alpine
LABEL maintainer="celso.rodrigues@hotmail.co.uk"

WORKDIR /app

ENV SERVER_PORT 80

EXPOSE ${SERVER_PORT}
RUN apk --no-cache update && apk --no-cache add ca-certificates
COPY --from=builder ["/go/src/celso/webserver", "./"] 
COPY --from=builder ["/go/src/celso/deployment.yaml", "./"]

ENTRYPOINT ["./webserver"]
