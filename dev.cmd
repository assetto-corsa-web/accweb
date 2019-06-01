Rem This file is for local development only!
Rem It configures and starts accweb for local development.

set ACCWEB_LOGLEVEL=debug
set ACCWEB_WATCH_BUILD_JS=true
set ACCWEB_ALLOWED_ORIGINS=*
set ACCWEB_HOST=localhost:8080
set ACCWEB_ADMIN_PASSWORD=admin
set ACCWEB_MOD_PASSWORD=mod
set ACCWEB_RO_PASSWORD=ro
set ACCWEB_TOKEN_PUBLIC_KEY=secrets/token.public
set ACCWEB_TOKEN_PRIVATE_KEY=secrets/token.private
set ACCWEB_CONFIG_PATH=config/
Rem set ACCWEB_SERVER_DIR=dev_server/
Rem set ACCWEB_SERVER_EXE=main.exe
set ACCWEB_SERVER_DIR=C:\Program Files (x86)\Steam\steamapps\common\Assetto Corsa Competizione\server
set ACCWEB_SERVER_EXE=accServer.exe

go run main.go
