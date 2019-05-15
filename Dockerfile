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
RUN ./gen_rsa_keys.sh

# build frontend
RUN cd /go/src/github.com/assetto-corsa-web/accweb/public && npm i && npm rebuild node-sass && npm run build

FROM alpine
COPY --from=build /go/src/github.com/assetto-corsa-web/accweb /app
WORKDIR /app

# default config
ENV ACCWEB_LOGLEVEL=info
ENV ACCWEB_WATCH_BUILD_JS=false
ENV ACCWEB_ALLOWED_ORIGINS=*
ENV ACCWEB_HOST=0.0.0.0:8080
ENV ACCWEB_ADMIN_PASSWORD=
ENV ACCWEB_MOD_PASSWORD=
ENV ACCWEB_RO_PASSWORD=
ENV ACCWEB_TOKEN_PUBLIC_KEY=/app/secrets/token.public
ENV ACCWEB_TOKEN_PRIVATE_KEY=/app/secrets/token.private
ENV ACCWEB_CONFIG_PATH=/app/config/
ENV ACCWEB_SERVER_DIR=/acc/TODO/
ENV ACCWEB_SERVER_EXE=TODO

# expose ACC installation and accweb configuration directory
VOLUME ["/acc", "/app/config"]

CMD ["/app/main"]
