#!/usr/bin/env bash

set -euo pipefail
IFS=$'\n\t'

ENVIRONMENT_FILE="bin/shared/environment.sh"
if [ -f "$ENVIRONMENT_FILE" ]; then
    source "$ENVIRONMENT_FILE"
fi

setup_environment
show_config

INSTANCE=$SPANNER_INSTANCE
DATABASE=$SPANNER_DATABASE
PROJECT=$SPANNER_PROJECT

echo "🚀 Insertando Seed Data directo en Cloud Spanner (${PROJECT})..."

gcloud config set project "$PROJECT"

echo "📝 Insertando Subscriptions..."
gcloud spanner databases execute-sql "$DATABASE" --instance="$INSTANCE" --sql="
INSERT INTO subscriptions (id, subscriber_name, event_name, source, created_at, updated_at) VALUES
('sub-uuid-0001-alpha', 'SubscriberWorkerA', 'LoadTestEventAlpha', 'LoadTestSource', CURRENT_TIMESTAMP(), CURRENT_TIMESTAMP()),
('sub-uuid-0002-bravo', 'SubscriberWorkerB', 'LoadTestEventBravo', 'LoadTestSource', CURRENT_TIMESTAMP(), CURRENT_TIMESTAMP()),
('sub-uuid-0003-charlie', 'SubscriberWorkerC', 'LoadTestEventCharlie', 'LoadTestSource', CURRENT_TIMESTAMP(), CURRENT_TIMESTAMP());"

echo "📝 Insertando Events..."
gcloud spanner databases execute-sql "$DATABASE" --instance="$INSTANCE" --sql="
INSERT INTO events (id, slug, source, payload, created_at, updated_at) VALUES
('evt-uuid-0001', 'loadtesteventalpha', 'LoadTestSource', '{\"iteration\": \"A\", \"message\": \"Seed directo\"}', CURRENT_TIMESTAMP(), CURRENT_TIMESTAMP()),
('evt-uuid-0002', 'loadtesteventbravo', 'LoadTestSource', '{\"iteration\": \"B\", \"message\": \"Seed directo\"}', CURRENT_TIMESTAMP(), CURRENT_TIMESTAMP());"

echo "📝 Insertando Queue Messages..."
gcloud spanner databases execute-sql "$DATABASE" --instance="$INSTANCE" --sql="
INSERT INTO queue (id, subscriber_name, event_id, status, created_at, updated_at) VALUES
('q-uuid-0001', 'SubscriberWorkerA', 'evt-uuid-0001', 'pending', CURRENT_TIMESTAMP(), CURRENT_TIMESTAMP()),
('q-uuid-0002', 'SubscriberWorkerB', 'evt-uuid-0002', 'processed', CURRENT_TIMESTAMP(), CURRENT_TIMESTAMP());"

echo "🎉 ¡Poblado directo en Cloud Spanner finalizado!"