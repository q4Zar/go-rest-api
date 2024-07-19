#!/bin/bash

set -e
# set -x  # Print each command before executing it

host="$1"

# Wait until the Go API is accessible
until curl -s -f -o /dev/null "$host/users/ready"; do
  # >&2 echo "Go API is unavailable - sleeping"
  sleep 2
done

# >&2 echo "Go API is up - executing scripts"

# Execute the first script and display its output
echo "Executing tests-damien.sh $host"
sh tests-damien.sh $host
if [ $? -ne 0 ]; then
  echo "Script tests-damien.sh failed"
  exit 1
fi

# Execute the second script and display its output
echo "Executing tests-qazar.sh $host"
sh tests-qazar.sh $host
if [ $? -ne 0 ]; then
  echo "Script tests-qazar.sh failed"
  exit 1
fi

echo "All scripts executed successfully"