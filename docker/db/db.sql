-- for uint8, int8, int16 on program side, use smallint,
-- uint16, int32 == integer (serial if AI)
-- uint32, int64 == bigint (bigserial if AI)
-- uint64 == numeric


-- ACCOUNT-RELATED
-- ================================================================

DROP TABLE IF EXISTS "accounts";
CREATE TABLE "public"."accounts" (
    "wallet_address" text NOT NULL PRIMARY KEY,
    "user_id" bigserial NOT NULL,
    "username" text,
    "language" text,
    "api_token" text,
    "sst_wei" text DEFAULT 0,
    "sbg_wei" text DEFAULT 0,
    "created_at" timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP
) WITH (oids = false);
CREATE UNIQUE INDEX idx_accounts_user_id ON accounts(user_id);