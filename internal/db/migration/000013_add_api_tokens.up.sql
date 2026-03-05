CREATE TABLE api_tokens (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    name VARCHAR(255) NOT NULL,
    token_hash VARCHAR(255) NOT NULL UNIQUE,
    token_prefix VARCHAR(10) NOT NULL,
    permissions JSONB NOT NULL DEFAULT '{
        "station_ids": null,
        "variables": null,
        "max_days_history": 365
    }',
    last_used_at timestamptz,
    expires_at timestamptz,
    created_at timestamptz DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamptz DEFAULT '0001-01-01 00:00:00Z'
);

CREATE INDEX "api_tokens_token_hash_idx" ON api_tokens(token_hash);
CREATE INDEX "api_tokens_user_id_idx" ON api_tokens(user_id);

ALTER TABLE "api_tokens" ADD CONSTRAINT "api_tokens_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE;
