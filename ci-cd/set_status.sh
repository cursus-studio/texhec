#!/bin/bash

export TOKEN
export URL

export STATE
export TARGET_URL
export DESCRIPTION
export CONTEXT

curl -L \
  -X POST \
  -H "Accept: application/vnd.github+json" \
  -H "Authorization: Bearer $TOKEN" \
  -H "X-GitHub-Api-Version: 2026-03-10" \
  $URL \
  -d "{\"state\":\"$STATE\",\"target_url\":\"$TARGET_URL\",\"description\":\"$DESCRIPTION\",\"context\":\"$CONTEXT\"}"
