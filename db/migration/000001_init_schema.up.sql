CREATE TABLE "user" (
                        "id" bigserial PRIMARY KEY,
                        "name" varchar(16) UNIQUE NOT NULL,
                        "password" varchar NOT NULL,
                        "nickname" varchar(16) NOT NULL,
                        "avatarURL" varchar NOT NULL,
                        "friendCount" int NOT NULL DEFAULT 0,
                        "friends" bigint[] NOT NULL,
                        "createAt" timestamp NOT NULL DEFAULT (now()),
                        "updateAt" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "message" (
                           "id" bigserial PRIMARY KEY,
                           "senderId" bigInt NOT NULL,
                           "receiverId" bigInt NOT NULL,
                           "content" text NOT NULL,
                           "sendAt" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "friendship" (
                              "id" bigserial PRIMARY KEY,
                              "fromId" bigInt NOT NULL,
                              "toId" bigInt NOT NULL,
                              "status" smallint NOT NULL DEFAULT 0,
                              "createAt" timestamp NOT NULL DEFAULT (now()),
                              "updateAt" timestamp NOT NULL DEFAULT (now())
);

ALTER TABLE "message" ADD FOREIGN KEY ("senderId") REFERENCES "user" ("id");

ALTER TABLE "message" ADD FOREIGN KEY ("receiverId") REFERENCES "user" ("id");

ALTER TABLE "friendship" ADD FOREIGN KEY ("fromId") REFERENCES "user" ("id");

ALTER TABLE "friendship" ADD FOREIGN KEY ("toId") REFERENCES "user" ("id");
