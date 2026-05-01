#!/bin/bash

export TOKEN
export OWNER
export REPO
export GIT_COMMIT

export STATE
export TARGET_URL
export DESC
export CONTEXT

curl -L \
  -X POST \
  -H "Accept: application/vnd.github+json" \
  -H "Authorization: Bearer $TOKEN" \
  -H "X-GitHub-Api-Version: 2026-03-10" \
  https://api.github.com/repos/$OWNER/$REPO/statuses/$GIT_COMMIT \
  -d "{\"state\":\"$STATE\",\"target_url\":\"$TARGET_URL\",\"description\":\"$DESC\",\"context\":\"$CONTEXT\"}"
