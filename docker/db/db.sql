-- for uint8, int8, int16 on program side, use smallint,
-- uint16, int32 == integer (serial if AI)
-- uint32, int64 == bigint (bigserial if AI)
-- uint64 == numeric


-- ACCOUNT-RELATED
-- ================================================================

DROP TABLE IF EXISTS "account";
CREATE TABLE "public"."account" (
    "id" bigserial NOT NULL PRIMARY KEY,
    "name" text,
    "email" text,
    "power" integer,
    "age" integer,
    "bot" jsonb,
    "created_at" timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP
) WITH (oids = false);
CREATE UNIQUE INDEX idx_account_user_id ON account(user_id);