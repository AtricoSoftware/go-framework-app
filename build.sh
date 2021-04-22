# Generated 2021-03-30 15:32:41 by go-framework development-version
# SECTION-START: Definitions
MODULE="github.com/AtricoSoftware/go-framework-app"
export OUTPUT_NAME="go-framework"
BUILD_DIR=release
TARGET_PLATFORMS="darwin windows linux"

if [[ ! -z "$1" ]]
then
  VERSION=$1
else
  VERSION=$(git describe --tags --dirty)
fi

export CGO_ENABLED=0
export GOARCH="amd64"

# setup details
# built
BUILT_ON=$(date)
BUILT_BY=$(whoami)
# git
GIT_BRANCH=$(git rev-parse --abbrev-ref HEAD)
GIT_COMMIT=$(git rev-parse HEAD)

DETAILS="{\"Built\":{\"On\":\"$BUILT_ON\", \"By\":\"$BUILT_BY\"},\"Git\":{ \"Repository\":\"$MODULE\",\"Branch\":\"$GIT_BRANCH\",\"Commit\":\"$GIT_COMMIT\"} }"
# Setup ldflags
LDFLAGS="-s -w"
LDFLAGS=$LDFLAGS" -X '$MODULE/pkg.Version=$VERSION'"
LDFLAGS=$LDFLAGS" -X '$MODULE/pkg.BuildDetails=$DETAILS'"
# SECTION-END

go run ./create_resources.go

# SECTION-START: Build
mkdir -p $TARGET_DIR
for GOOS in $TARGET_PLATFORMS; do
    export GOOS
    export EXT=""
    if [[ ${GOOS} == "windows" ]]
    then
      export EXT=".exe"
    fi
    # Version build
    export TARGET_DIR="$BUILD_DIR/$VERSION-$GOOS-$GOARCH"
    export TARGET_APP="$TARGET_DIR/$OUTPUT_NAME$EXT"
    mkdir -p TARGET_DIR
    echo Building $TARGET_APP
    go build -v -ldflags="$LDFLAGS" -o $TARGET_APP
    echo Packaging TARGET_DIR.zip
    zip -j1 TARGET_DIR.zip TARGET_APP
    # Copy app to latest
    export LATEST_DIR="$BUILD_DIR/latest-$GOOS-$GOARCH"
    mkdir -p $LATEST_DIR
    echo Copying to $LATEST_DIR
    cp $TARGET_APP $LATEST_DIR/
done
# SECTION-END

