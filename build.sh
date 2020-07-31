MODULE="dev.azure.com/MAT-OCS/ConditionInsight/_git/ma.ci.go-framework-app"
export OUTPUT_NAME="go-framework-app"
TARGET_DIR=release
TARGET_PLATFORMS="darwin windows linux"
VERSION=$(git describe --tags --dirty)
GIT_SHA=$(git rev-parse HEAD)
BUILT_ON=$(date)
BUILT_BY=$(whoami)

# Caller can specify extra info in version
if [[ ! -z "$1" ]]
then
  VERSION=$VERSION-$1
fi
echo Building Version $VERSION

export CGO_ENABLED=0
export GOARCH="amd64"

# Setup ldflags
LDFLAGS="-s -w"
LDFLAGS=$LDFLAGS" -X '$MODULE/pkg.Version=$VERSION'"
LDFLAGS=$LDFLAGS" -X '$MODULE/pkg.Git=$GIT_SHA'"
LDFLAGS=$LDFLAGS" -X '$MODULE/pkg.BuiltOn=$BUILT_ON'"
LDFLAGS=$LDFLAGS" -X '$MODULE/pkg.BuiltBy=$BUILT_BY'"

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
    go build -v -ldflags="$LDFLAGS" -o $TARGET/$OUTPUT_NAME$EXT

done

cd $TARGET_DIR
find . ! -path . -type d |  cut -d "/" -f2 | awk -v name="$OUTPUT_NAME" '{ print name "_" $1 ".zip -r ./" $1 "/"  }' | xargs -L1 zip -j
#find . ! -path . -type d | xargs -L1 rm -rf

