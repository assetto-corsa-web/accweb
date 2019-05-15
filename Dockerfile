FROM golang AS build
ADD . /go/src/github.com/assetto-corsa-web/accweb
WORKDIR /go/src/github.com/assetto-corsa-web/accweb
RUN apt update && \
	apt upgrade -y && \
	apt install curl  -y
RUN curl -sL https://deb.nodesource.com/setup_8.x -o nodesource_setup.sh && bash nodesource_setup.sh
RUN apt-get install -y nodejs

# build backend
ENV GOPATH=/go
RUN go build -ldflags "-s -w" main.go

# build frontend
RUN cd /go/src/github.com/assetto-corsa-web/accweb/public && npm i && npm rebuild node-sass && npm run build

FROM alpine
COPY --from=build /go/src/github.com/assetto-corsa-web/accweb /app
WORKDIR /app

# default config
# TODO

# expose ACC installation and accweb configuration directory
VOLUME ["/acc", "/app/config"]

CMD ["/app/main"]
