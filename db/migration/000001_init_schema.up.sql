CREATE TABLE "user" (
                        "id" bigserial PRIMARY KEY,
                        "name" varchar(16) UNIQUE NOT NULL,
                        "password" varchar NOT NULL,
                        "nickname" varchar(16) NOT NULL,
                        "avatar_url" varchar NOT NULL,
                        "friend_count" int NOT NULL DEFAULT 0,
                        "friends" bigint[] ,
                        "create_at" timestamp NOT NULL DEFAULT (now()),
                        "update_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "message" (
                           "id" bigserial PRIMARY KEY,
                           "sender_id" bigInt NOT NULL,
                           "receiver_id" bigInt NOT NULL,
                           "content" text NOT NULL,
                           "send_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "friendship" (
                              "id" bigserial PRIMARY KEY,
                              "from_id" bigInt NOT NULL,
                              "to_id" bigInt NOT NULL,
                              "status" smallint NOT NULL DEFAULT 0,
                              "create_at" timestamp NOT NULL DEFAULT (now()),
                              "update_at" timestamp NOT NULL DEFAULT (now())
);

ALTER TABLE "message" ADD FOREIGN KEY ("sender_id") REFERENCES "user" ("id");

ALTER TABLE "message" ADD FOREIGN KEY ("receiver_id") REFERENCES "user" ("id");

ALTER TABLE "friendship" ADD FOREIGN KEY ("from_id") REFERENCES "user" ("id");

ALTER TABLE "friendship" ADD FOREIGN KEY ("to_id") REFERENCES "user" ("id");
