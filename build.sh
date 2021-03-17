# Generated 2021-03-17 16:07:26 by go-framework V1.8.0
# SECTION-START: Definitions
MODULE="github.com/AtricoSoftware/go-framework-app"
export OUTPUT_NAME="go-framework"
TARGET_DIR=release
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
    export TARGET="$TARGET_DIR/$VERSION-$GOOS-$GOARCH"
    mkdir -p $TARGET
    echo Building $TARGET
    go build -v -ldflags="$LDFLAGS" -o $TARGET/$OUTPUT_NAME$EXT
    echo Packaging $TARGET.zip
    zip -j1 $TARGET.zip $TARGET/$OUTPUT_NAME$EXT
done
# SECTION-END

