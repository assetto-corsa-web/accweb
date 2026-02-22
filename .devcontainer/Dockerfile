FROM golang:1.25-bookworm

RUN apt-get update && export DEBIAN_FRONTEND=noninteractive

RUN apt-get install -y git build-essential zip curl apt-transport-https ca-certificates gnupg2
RUN curl -fsSL https://deb.nodesource.com/setup_22.x | bash -
RUN apt-get install -y nodejs
RUN npm install -g pnpm

RUN go install golang.org/x/tools/gopls@latest \
    && go install github.com/go-delve/delve/cmd/dlv@latest \
    && go install honnef.co/go/tools/cmd/staticcheck@latest \
    && go install github.com/swaggo/swag/cmd/swag@latest \
    && curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/HEAD/install.sh | sh -s -- -b $(go env GOPATH)/bin v2.7.2

RUN mkdir -p /app && cd /app && git clone https://github.com/assetto-corsa-web/fake-accserver.git

RUN cd /app/fake-accserver/logs \
    && wget -q https://gist.github.com/pedrofaria/ab057054e31eb5c7960a79fe3116c3d9/raw/e42339008b0ddb17989cd5e0d2c76e0e0041521c/logs_20220315_201400_1636636807.log \
    && cd /app/fake-accserver && make build-for-osx

