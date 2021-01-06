#!/bin/bash

VERSION=""

#get parameters
while getopts v: flag
do
  case "${flag}" in
    v) VERSION=${OPTARG};;
  esac
done

#get highest tag number, and add 0.0.1 if doesn't exist
CURRENT_VERSION=`git describe --abbrev=0 --tags 2>/dev/null`

if [[ $CURRENT_VERSION == '' ]]
then
  CURRENT_VERSION='0.0.1'
fi
echo "Current Version: $CURRENT_VERSION"

# split
VERSION_ARRAY=(${CURRENT_VERSION//./ })

#get number parts
MAJOR=${VERSION_ARRAY[0]}
MINOR=${VERSION_ARRAY[1]}
PATCH=${VERSION_ARRAY[2]}

if [[ $VERSION == 'major' ]]
then
  MAJOR=$((MAJOR+1))
elif [[ $VERSION == 'minor' ]]
then
  MINOR=$((MINOR+1))
elif [[ $VERSION == 'patch' ]]
then
  PATCH=$((PATCH+1))
else
  echo "No version type or incorrect type specified, try: -v [major, minor, patch]"
  exit 1
fi

#create new tag
NEW_TAG="$MAJOR.$MINOR.$PATCH"
echo "($VERSION) updating $CURRENT_VERSION to $NEW_TAG"

#get current hash and see if it already has a tag
GIT_COMMIT=$(git rev-parse HEAD)
NEEDS_TAG=$(git describe --contains "$GIT_COMMIT" 2>/dev/null)

if [ -z "$NEEDS_TAG" ]; then
  echo "Tagged with $NEW_TAG"
  git push --tags
  git push
else
  echo "Already a tag on this commit"
fi

exit 0