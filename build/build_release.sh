#!/bin/sh

VERSION=$1

if [ -z $VERSION ]; then
  echo "Usage: ./build/build_release.sh <Version Number, e.g. 1.2.3>"
  exit
fi

# create release directory
RDIR="releases"
VDIR="accweb_$VERSION"
DIR="$RDIR/$VDIR"
mkdir -p "$DIR"

# build frontend
cd public
npm i
npm run build

# build backend (Windows and Linux)
cd ..
CGO_ENABLED=0 GOOS=linux go build -ldflags "-s -w" -o $DIR/accweb cmd/main.go
CGO_ENABLED=0 GOOS=windows go build -ldflags "-s -w" -o $DIR/accweb.exe cmd/main.go

# copy files
cp gen_rsa_keys.sh "$DIR/gen_rsa_keys.sh"
cp gen_rsa_keys.cmd "$DIR/gen_rsa_keys.cmd"
cp LICENSE "$DIR/LICENSE"
cp README.md "$DIR/README.md"
cp build/sample_config.yml "$DIR/config.yml"

# make scripts and accweb executable
chmod +x "$DIR/accweb"
chmod +x "$DIR/accweb.exe"
chmod +x "$DIR/gen_rsa_keys.sh"
chmod +x "$DIR/gen_rsa_keys.cmd"

cd "$RDIR"
zip -r "$VDIR.zip" "$VDIR"
cd ..
rm -r $DIR

echo "done"
