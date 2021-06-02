{"Type":"Mixed"}
# {{.Comment}}

# SECTION-START: SetDefaults
# If different values are required, (re)set these after this section
DEFAULT_VERSION=NoVersion
DEFAULT_ARCH=amd64
DEFAULT_PLATFORM=linux
BUILD_DIR=release
# SECTION-END

# SECTION-START: Commandline
PLATFORM=()
ARCH=()
while [[ $# -gt 0 ]]; do
  key="$1"
  case $key in
      -v|--version)
      VERSION="$2"
      shift # past argument
      shift # past value
      ;;
      -a|--arch)
      ARCH+=("$2")
      shift # past argument
      shift # past value
      ;;
      -p|--platform)
      PLATFORM+=("$2")
      shift # past argument
      shift # past value
      ;;
      *)    # unknown option, treat as platform
      PLATFORM+=("$1")
      shift # past argument
      ;;
  esac
done
# Add defaults if missing
if [[ -z "$VERSION" ]]; then
  VERSION=$(git describe --tags --dirty --always)
  if [[ -z "$VERSION" ]]; then
    VERSION=$DEFAULT_VERSION
  fi
fi
if [[ ${#ARCH[@]} -eq 0 ]]; then
    ARCH=($DEFAULT_ARCH)
fi
if [[ ${#PLATFORM[@]} -eq 0 ]]; then
    PLATFORM=($DEFAULT_PLATFORM)
fi
# SECTION-END

# SECTION-START: Definitions
MODULE="{{.RepositoryPath}}"
export OUTPUT_NAME="{{.ApplicationName}}"
export CGO_ENABLED=0

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

# SECTION-START: Build
echo Build dir = $BUILD_DIR
mkdir -p $BUILD_DIR
for GOARCH in ${ARCH[@]}; do
    export GOARCH
    for GOOS in ${PLATFORM[@]}; do
        export GOOS
        export EXT=""
        if [[ ${GOOS} == "windows" ]]; then
          export EXT=".exe"
        fi
        # Version build
        export TARGET_DIR="$BUILD_DIR/$VERSION-$GOOS-$GOARCH"
        export TARGET_APP="$TARGET_DIR/$OUTPUT_NAME$EXT"
        mkdir -p $TARGET_DIR
        echo Building $TARGET_APP
        go build -v -ldflags="$LDFLAGS" -o $TARGET_APP
        echo Packaging $TARGET_DIR.zip
        zip -j1 $TARGET_DIR.zip $TARGET_APP
        # Copy app to latest
        export LATEST_DIR="$BUILD_DIR/latest-$GOOS-$GOARCH"
        mkdir -p $LATEST_DIR
        echo Copying to $LATEST_DIR
        cp $TARGET_APP $LATEST_DIR/
    done
done
# SECTION-END
