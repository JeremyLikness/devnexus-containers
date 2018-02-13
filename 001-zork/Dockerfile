FROM golang:alpine as build-env
RUN mkdir /src
ADD http://msinilo.pl/download/zmachine.go /src
RUN cd /src && go build -o goapp

FROM alpine
RUN mkdir /app
WORKDIR /app
COPY --from=build-env /src/goapp /app
ADD https://github.com/visnup/frotz/blob/master/lib/ZORK1.DAT?raw=true /app/zork1.dat
ENTRYPOINT ./goapp