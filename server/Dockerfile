FROM golang:alpine as build
WORKDIR /src
COPY server.go server.go
RUN go build server.go

FROM alpine
WORKDIR /server
COPY --from=build /src/server /server/server
EXPOSE 8080
ENTRYPOINT /server/server