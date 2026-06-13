#!/usr/bin/env bash

set -euo pipefail
IFS=$'\n\t'
ENVIRONMENT_FILE="bin/shared/environment.sh"
source $ENVIRONMENT_FILE

SERVER="localhost:30000"
SERVICE="event.Eventservice"
SLUG="EventTest"
SOURCE="EventSource"
HOW_MANY=100

echo "🧹 Cleaning up database..."
docker exec -it goevents-postgres psql -U admin -d goevents -c "TRUNCATE TABLE queue, events, subscriptions RESTART IDENTITY CASCADE;" > /dev/null

echo "🚀 Starting Bulk E2E gRPC Test..."
EVENT_IDS=()

for i in $(seq 1 10); do
    SUB_PAYLOAD=$(jq -n --arg name "user_$i" --arg slug "$SLUG" --arg src "$SOURCE" \
        '{subscriber_name: $name, event_name: $slug, source: $src}')
    grpcurl -plaintext -d "$SUB_PAYLOAD" $SERVER $SERVICE/CreateSubscription > /dev/null 2>&1
done
echo "✅ 10 Subscriptions created."

for i in $(seq 1 $HOW_MANY); do
    CREATE_PAYLOAD=$(jq -n --arg slug "$SLUG" --arg src "$SOURCE" --arg pld "{\"i\": $i}" \
        '{slug: $slug, source: $src, payload: $pld}')
    
    RESP=$(grpcurl -plaintext -d "$CREATE_PAYLOAD" $SERVER $SERVICE/CreateEvent)
    EVENT_IDS+=("$(echo "$RESP" | jq -r '.id')")
    echo -ne "🛠  Creating events: $i/$HOW_MANY\r"
done
echo -e "\n✅ $HOW_MANY Events created."

PULL_PAYLOAD=$(jq -n --arg slug "$SLUG" --arg src "$SOURCE" '{event_name: $slug, source: $src}')
PULL_RESP=$(grpcurl -plaintext -d "$PULL_PAYLOAD" $SERVER $SERVICE/PullMessages)
QUEUE_IDS=$(echo "$PULL_RESP" | jq -c '[.messages[].id]')

if [ "$QUEUE_IDS" != "[]" ]; then
    ACK_PAYLOAD=$(jq -n --argjson ids "$QUEUE_IDS" '{queue_ids: $ids}')
    grpcurl -plaintext -d "$ACK_PAYLOAD" $SERVER $SERVICE/AckMessages > /dev/null
    echo "✅ Acknowledgement successful."
fi

echo "🗑️  Deleting events..."
for i in "${!EVENT_IDS[@]}"; do
    DELETE_PAYLOAD=$(jq -n --arg id "${EVENT_IDS[$i]}" '{id: $id}')
    grpcurl -plaintext -d "$DELETE_PAYLOAD" $SERVER $SERVICE/DeleteEvent > /dev/null
    echo -ne "🔥 Deleting events: $((i + 1))/$HOW_MANY\r"
done
echo -e "\n✅ All events deleted."

FINAL_PULL=$(grpcurl -plaintext -d "$PULL_PAYLOAD" $SERVER $SERVICE/PullMessages)
if [ "$(echo "$FINAL_PULL" | jq '.messages | length // 0')" -eq 0 ]; then
    echo "🎉 Test passed: Queue is empty."
else
    echo "❌ Test failed: Messages remain in queue."
    exit 1
fi

echo "🧹 Cleaning up database..."
docker exec -it goevents-postgres psql -U admin -d goevents -c "TRUNCATE TABLE queue, events, subscriptions RESTART IDENTITY CASCADE;" > /dev/null
