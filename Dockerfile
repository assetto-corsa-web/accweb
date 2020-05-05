FROM golang AS build

LABEL maintainer="Kugel" \
	  contributors="GillesDubois" \
      version="latest" \
      description="Assetto Corsa Competizione Server Management Tool via Web Interface."

COPY . /go/src/github.com/assetto-corsa-web/accweb

WORKDIR /go/src/github.com/assetto-corsa-web/accweb

RUN apt update && \
	apt upgrade -y && \
	apt install curl  -y
RUN curl -sL https://deb.nodesource.com/setup_13.x -o nodesource_setup.sh && bash nodesource_setup.sh
RUN apt-get install -y nodejs

ENV GOPATH=/go
RUN go build -tags netgo -a -v -ldflags "-s -w" main.go 
RUN ./gen_rsa_keys.sh

RUN cd /go/src/github.com/assetto-corsa-web/accweb/public && npm i && npm rebuild node-sass && npm run build

# Final image
FROM alpine:3.9.6

COPY --from=build /go/src/github.com/assetto-corsa-web/accweb /accweb

ENV ACCWEB_HOST=localhost:8080 \
	ACCWEB_ENABLE_TLS=false \
	ACCWEB_CERT_FILE=/sslcerts/certificate.crt \
	ACCWEB_PRIV_FILE=/sslcerts/private.key \
	ACCWEB_ADMIN_PASSWORD=weakadminpassword \
	ACCWEB_MOD_PASSWORD=weakmodpassword \
	ACCWEB_RO_PASSWORD=weakropassword \
	ACCWEB_LOGLEVEL=info \
	ACCWEB_CORS=*

VOLUME /accserver /accweb /sslcerts

WORKDIR /accweb

RUN apk add gettext wine wine-dev wine-libs

EXPOSE 8080

ENTRYPOINT [ "/accweb/docker-entrypoint.sh" ] 