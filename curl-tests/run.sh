#!/bin/bash

set -e

host="$1"

# Wait until the Go API is accessible
until curl -s -f -o /dev/null "$host/assets"; do
  >&2 echo "Go API is unavailable - sleeping"
  sleep 2
done

>&2 echo "Go API is up - executing scripts"

# Execute the first script
sh tests-damien.sh
if [ $? -ne 0 ]; then
  echo "Script 1 failed"
  exit 1
fi

# Execute the second script
sh tests-qazar.sh
if [ $? -ne 0 ]; then
  echo "Script 2 failed"
  exit 1
fi

echo "All scripts executed successfully"