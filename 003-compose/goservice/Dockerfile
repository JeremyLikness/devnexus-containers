FROM golang:alpine as build-env
RUN mkdir /src
COPY webserver.go /src
RUN cd /src && go build -o webserver .

FROM alpine:latest
COPY --from=build-env /src/webserver /
ENTRYPOINT ./webserver