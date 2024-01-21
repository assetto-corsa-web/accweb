#!/bin/sh

VERSION=$1
COMMIT=$2

if [ -z $VERSION ]; then
  echo "Usage: ./build/build_release.sh <Version Number, e.g. 1.2.3>"
  exit
fi

if [ -z $COMMIT ]; then
  COMMIT=`git rev-parse --short HEAD`
fi

COMMIT=$(echo "${COMMIT}" | cut -c1-7)

echo "Starting to build accweb $VERSION ( $COMMIT )"
node -v

# create release directory
RDIR="releases"
VDIR="accweb_$VERSION"
DIR="$RDIR/$VDIR"
mkdir -p "$DIR"

# build frontend
cd public
# COMMIT=`git rev-parse --short HEAD`
cp src/components/end.vue src/components/end.vue.orig
sed -i "s/%VERSION%/$VERSION/g" src/components/end.vue
sed -i "s/%COMMIT%/$COMMIT/g" src/components/end.vue
npm i
npm run build
mv src/components/end.vue.orig src/components/end.vue

# build backend (Windows and Linux)
cd ..
# CGO_ENABLED=0 GOOS=linux go build -ldflags "-s -w" -o $DIR/accweb cmd/main.go
CGO_ENABLED=0 GOOS=windows go build -ldflags "-s -w" -o $DIR/accweb.exe cmd/main.go

# copy files
cp LICENSE "$DIR/LICENSE"
cp README.md "$DIR/README.md"
cp build/sample_config.yml "$DIR/config.yml"

# make scripts and accweb executable
# chmod +x "$DIR/accweb"
chmod +x "$DIR/accweb.exe"

cd "$RDIR"
zip a "$VDIR.zip" "$VDIR"
cd ..
# rm -r $DIR

echo "done"
