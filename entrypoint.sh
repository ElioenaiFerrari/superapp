#!/bin/sh

echo "Running $APP"
if [ "$APP" = "api" ]; then
  dlv exec ./main --headless --listen=:2345 --api-version=2 --log
else
  ./main
fi