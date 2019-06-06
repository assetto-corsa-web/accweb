# Adjust this file to match your needs.

# sets host and port
# 80 is the default port for web servers and meight be blocked on your machine
# in that case, change it to anything else, e.g. 8080
export ACCWEB_HOST=0.0.0.0:80

# IMPORTANT must be set, use strong different passwords for each
export ACCWEB_ADMIN_PASSWORD=
export ACCWEB_MOD_PASSWORD=
export ACCWEB_RO_PASSWORD=

# these files must exist
export ACCWEB_TOKEN_PUBLIC_KEY=secrets/token.public
export ACCWEB_TOKEN_PRIVATE_KEY=secrets/token.private

# create this directory in your installation directory
export ACCWEB_CONFIG_PATH=config/

# export the installation path of the ACC server
export ACCWEB_SERVER_DIR="/home/username/.steam/steam/steamapps/common/Assetto Corsa Competizione/server"

export ACCWEB_LOGLEVEL=info
export ACCWEB_ALLOWED_ORIGINS=*
export ACCWEB_SERVER_EXE=accServer.exe

./accweb
