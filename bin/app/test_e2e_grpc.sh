#!/usr/bin/env bash

set -euo pipefail
IFS=$'\n\t'
ENVIRONMENT_FILE="bin/shared/environment.sh"
source $ENVIRONMENT_FILE

SERVER="localhost:30000"
SERVICE="event.Eventservice"

echo "🚀 Starting E2E gRPC Test for $SERVICE at $SERVER..."

# Ensure the default directory exists in case the server does not have the environment variables
mkdir -p /tmp/events

# Check dependencies
if ! command -v grpcurl &> /dev/null; then
    echo "❌ grpcurl is not installed. Run 'make support-install-grpc-tools' first."
    exit 1
fi

if ! command -v jq &> /dev/null; then
    echo "❌ jq is not installed. Please install it (e.g., sudo apt install jq / brew install jq)."
    exit 1
fi

echo "--------------------------------------------------"
echo "0️⃣  Creating a new Subscription..."
# Creamos la suscripción antes para que el Fan-Out se entere al crear el evento
SUB_PAYLOAD='{
  "subscriber_name": "marco",
  "event_name": "EventTest",
  "source": "EventSource"
}'
grpcurl -plaintext -d "$SUB_PAYLOAD" $SERVER $SERVICE/CreateSubscription
echo "✅ Subscription created successfully."
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

echo "3️⃣  Pulling Messages from Queue..."
# El worker pide sus mensajes pendientes basándose en el nombre y origen del evento
PULL_PAYLOAD='{
  "event_name": "EventTest",
  "source": "EventSource",
  "limit": 1
}'
PULL_RESP=$(grpcurl -plaintext -d "$PULL_PAYLOAD" $SERVER $SERVICE/PullMessages)
echo "$PULL_RESP"

# Extraemos el queueId del primer mensaje (grpcurl mapea snake_case a lowerCamelCase)
QUEUE_ID=$(echo "$PULL_RESP" | jq -r '.messages[0].queueId')

if [ -z "$QUEUE_ID" ] || [ "$QUEUE_ID" == "null" ]; then
    echo "❌ Error: Could not extract queueId from PullMessages response."
    exit 1
fi
echo "✅ Successfully pulled message with Queue ID: $QUEUE_ID"
echo "--------------------------------------------------"

echo "3️⃣b Acknowledging Message..."
# Confirmamos el procesamiento usando el ID único de la cola obtenido en el paso anterior
ACK_PAYLOAD="{\"queue_id\": \"$QUEUE_ID\"}"
ACK_RESP=$(grpcurl -plaintext -d "$ACK_PAYLOAD" $SERVER $SERVICE/AckMessage)
echo "$ACK_RESP"

ACK_SUCCESS=$(echo "$ACK_RESP" | jq -r '.success')
if [ "$ACK_SUCCESS" != "true" ]; then
    echo "❌ Error: AckMessage returned success=false."
    exit 1
fi
echo "✅ Message acknowledged successfully."
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