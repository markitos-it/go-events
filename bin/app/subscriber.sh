#!/usr/bin/env bash

set -euo pipefail
IFS=$'\n\t'

if ! command -v grpcurl &> /dev/null; then
    echo "❌ grpcurl is not installed. Run 'make support-install-grpc-tools' or install it manually."
    exit 1
fi

if ! command -v jq &> /dev/null; then
    echo "❌ jq is not installed. Please install it (e.g., 'sudo apt install jq' or 'brew install jq')."
    exit 1
fi

ENVIRONMENT_FILE="bin/shared/environment.sh"
if [ -f "$ENVIRONMENT_FILE" ]; then
    source "$ENVIRONMENT_FILE"
fi

SUBSCRIBER_NAME=${1:?"Error: The first parameter 'subscriberName' is required."}
EVENT_SLUG=${2:?"Error: The second parameter 'eventSlug' is required."}
SOURCE_SLUG=${3:?"Error: The third parameter 'sourceSlug' is required."}

SERVER=${SERVER:-"localhost:30000"}
SERVICE="event.Eventservice"

echo "--------------------------------------------------"
echo "Subscriber will run with the following parameters:"
echo "  - Subscriber Name: $SUBSCRIBER_NAME"
echo "  - Event Slug:      $EVENT_SLUG"
echo "  - Source Slug:     $SOURCE_SLUG"
echo "--------------------------------------------------"

read -p "Do you want to continue? (y/N): " -n 1 -r
echo
if [[ ! $REPLY =~ ^[Yy]$ ]]; then
    echo "Operation cancelled by user."
    exit 0
fi

echo "--------------------------------------------------"
echo "1️⃣  Registering subscriber..."
SUB_PAYLOAD=$(cat <<EOF
{
  "subscriber_name": "$SUBSCRIBER_NAME",
  "event_name": "$EVENT_SLUG",
  "source": "$SOURCE_SLUG"
}
EOF
)
grpcurl -plaintext -d "$SUB_PAYLOAD" "$SERVER" "$SERVICE/CreateSubscription" | jq .

echo "🎉 Process finished."
