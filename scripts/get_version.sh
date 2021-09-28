#!/usr/bin/env bash

# This is a branch helper script that can be used to determine what we're using
# or deploying.
# Defaulting to version 0.1.0 if we can't get version in any other way.

# We might already have version file (e.g. this script was launched with "generate"
# parameter when building a Docker image). In that case - just output it.
if [ -f "version" ]; then
	cat version
	exit 0
fi

DELIMITER=" "
VERSION="0.1.0"

# Try to get last available version tag. This tag will replace version if found.
TAG=$(git tag | tail -n 1)
if [ "${TAG}" != "" ]; then
	VERSION="${TAG}"
fi

# We should figure out should we add branch or not. For that we should take
# last commit hash and last tag hash for comparison. We should add branch
# name only if hashes aren't same.
LAST_COMMIT_HASH=$(git log --format="%h" 2>/dev/null | head -n 1)
LAST_TAG_COMMIT_HASH=$(git log "${TAG}" --format="%h" 2>/dev/null | head -n 1)
if [ "${LAST_COMMIT_HASH}" != "${LAST_TAG_COMMIT_HASH}" ]; then
	VERSION+="${DELIMITER}$(git branch | grep "*" | awk {' print $2 '})"
fi

# Add commit hash.
VERSION+="${DELIMITER}${LAST_COMMIT_HASH}"

# Add build date.
VERSION+="${DELIMITER}$(date +'%Y%m%d-%H%M')"

case $1 in
	generate)
		echo "${VERSION}" > version
	;;
	*)
		echo "${VERSION}"
	;;
esac
