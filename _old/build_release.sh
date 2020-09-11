#!/bin/bash

VERSION=$1

if [ -z $VERSION ]; then
  echo "Usage: ./build_release.sh <Version Number, e.g. 1.2.3>"
  exit
fi

# create release directory
DIR="accweb_$VERSION"
mkdir "$DIR"

# build frontend
cd public
npm i
npm run build

# build backend (Windows and Linux)
cd ..
go build -ldflags "-s -w" main.go
GOOS=windows go build -ldflags "-s -w" main.go

# copy files
mv main "$DIR/accweb"
mv main.exe "$DIR/accweb.exe"
cp gen_rsa_keys.sh "$DIR/gen_rsa_keys.sh"
cp gen_rsa_keys.cmd "$DIR/gen_rsa_keys.cmd"
cp LICENSE "$DIR/LICENSE"
cp README.md "$DIR/README.md"
cp sample_config.yml "$DIR/config.yml"
cp -r public "$DIR/public"
rm -rf "$DIR/public/node_modules"
rm -r "$DIR/public/src"
rm "$DIR/public/.gitignore"
rm "$DIR/public/package.json"
rm "$DIR/public/package-lock.json"
rm "$DIR/public/webpack.config.js"
rm "$DIR/public/static/main.scss"

# make scripts and accweb executable
chmod +x "$DIR/accweb"
chmod +x "$DIR/accweb.exe"
chmod +x "$DIR/gen_rsa_keys.sh"
chmod +x "$DIR/gen_rsa_keys.cmd"

echo "done"
