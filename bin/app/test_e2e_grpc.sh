#!/usr/bin/env bash

set -euo pipefail
IFS=$'\n\t'
ENVIRONMENT_FILE="bin/shared/environment.sh"
source $ENVIRONMENT_FILE

SERVER="localhost:30000"
SERVICE="event.Eventservice"

echo "🚀 Starting E2E gRPC Test for $SERVICE at $SERVER..."

mkdir -p /tmp/events

if ! command -v grpcurl &> /dev/null; then
    echo "❌ grpcurl is not installed. Run 'make support-install-grpc-tools' first."
    exit 1
fi

if ! command -v jq &> /dev/null; then
    echo "❌ jq is not installed. Please install it (e.g., sudo apt install jq / brew install jq)."
    exit 1
fi

echo "--------------------------------------------------"
echo "1️⃣  Creating a new Event..."
CREATE_PAYLOAD='{
  "name": "EventTest",
  "source": "EventSource",
  "payload": ""
}'
PLACEHOLDER_CREATE="__CUSTOM_E2E_""FIELDS_CREATE__"
echo "🔍 Placeholders found: $PLACEHOLDER_CREATE"
CREATE_PAYLOAD="${CREATE_PAYLOAD//$PLACEHOLDER_CREATE/}"
echo "🔍 Payload sent: $CREATE_PAYLOAD"
CREATE_RESP=$(grpcurl -plaintext -d "$CREATE_PAYLOAD" $SERVER $SERVICE/CreateEvent)

echo "$CREATE_RESP"
EVENT_ID=$(echo "$CREATE_RESP" | jq -r '.id')

if [ -z "$EVENT_ID" ] || [ "$EVENT_ID" == "null" ]; then
    echo "❌ Error: Could not extract the ID from the response."
    exit 1
fi
echo "✅ Successfully created with ID: $EVENT_ID"
echo "--------------------------------------------------"

echo "2️⃣  Getting Event by ID..."
grpcurl -plaintext -d "{\"id\": \"$EVENT_ID\"}" $SERVER $SERVICE/GetEvent
echo "✅ Get successful."
echo "--------------------------------------------------"

echo "4️⃣  Listing Events..."
grpcurl -plaintext -d "{\"name\": \"EventTest\", \"source\": \"EventSource\"}" $SERVER $SERVICE/AllByNameAndSource
echo "✅ List successful."
echo "--------------------------------------------------"

echo "5️⃣  Deleting Event..."
grpcurl -plaintext -d "{\"id\": \"$EVENT_ID\"}" $SERVER $SERVICE/DeleteEvent
echo "✅ Delete successful."
echo "--------------------------------------------------"

echo "6️⃣  Verifying that it was deleted..."
set +e
GET_DELETED_RESP=$(grpcurl -plaintext -d "{\"id\": \"$EVENT_ID\"}" $SERVER $SERVICE/GetEvent 2>&1)
set -e

if echo "$GET_DELETED_RESP" | grep -i -q "not found\|notfound"; then
    echo "✅ Verification successful: The record no longer exists."
else
    echo "❌ Warning: The record might still exist or an unexpected error occurred:"
    echo "$GET_DELETED_RESP"
    exit 1
fi

echo "--------------------------------------------------"
echo "🎉 All E2E tests passed successfully!"