FROM node:alpine AS client_builder
COPY client /client
WORKDIR /client
RUN npm install
RUN npm run build

FROM golang:1.13.6-alpine AS server_builder
RUN apk add --update --no-cache git build-base
RUN mkdir /build
WORKDIR /build
COPY ./server/go.mod .
COPY ./server/go.sum .
RUN go mod download
COPY ./server .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-w" -a -o bin/main .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=server_builder /build/bin /server
COPY --from=server_builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=server_builder /build/.env /server
COPY --from=server_builder /build/message.yaml /server
COPY --from=client_builder /client/build /client/build
WORKDIR /server
EXPOSE 8399
CMD ./main api