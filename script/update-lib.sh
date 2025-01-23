#!/usr/bin/bash
user="mhthrh"
repo="common-lib"

ver="$(curl -s https://api.github.com/repos/$user/$repo/commits | jq -r '.[0].sha')"

echo "go get github.com/mhthrh/common-lib@$ver"