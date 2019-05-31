Rem Adjust this file to match your needs.

Rem sets host and port
Rem 80 is the default port for web servers and meight be blocked on your machine
Rem in that case, change it to anything else, e.g. 8080
set ACCWEB_HOST=0.0.0.0:80

Rem IMPORTANT must be set, use strong different passwords for each
set ACCWEB_ADMIN_PASSWORD=a
set ACCWEB_MOD_PASSWORD=b
set ACCWEB_RO_PASSWORD=c

Rem these files must exist
set ACCWEB_TOKEN_PUBLIC_KEY=secrets/token.public
set ACCWEB_TOKEN_PRIVATE_KEY=secrets/token.private

Rem create this directory in your installation directory
set ACCWEB_CONFIG_PATH=config/

Rem set the installation path of the ACC server
set ACCWEB_SERVER_DIR=C:\Program Files (x86)\Steam\steamapps\common\Assetto Corsa Competizione\server

set ACCWEB_LOGLEVEL=info
set ACCWEB_ALLOWED_ORIGINS=*
set ACCWEB_SERVER_EXE=accServer.exe

go run main.go
