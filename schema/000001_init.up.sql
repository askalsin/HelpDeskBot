CREATE TABLE "public.users" (
    "id" serial NOT NULL,
    "chat_id" bigint NOT NULL UNIQUE,
    "name" varchar(255) NOT NULL,
    "type_account" varchar(15) NOT NULL DEFAULT 'client',
    "phone_number" varchar(20) NOT NULL,
    "email" varchar(320) NOT NULL,
    "organization" varchar(255) NOT NULL,
    "address" varchar(255) NOT NULL,
    CONSTRAINT "users_pk" PRIMARY KEY ("id")
);

CREATE TABLE "public.issues" (
    "id" serial NOT NULL,
    "chat_id" bigint NOT NULL REFERENCES "public.users"(chat_id),
    "issue" varchar(255) NOT NULL UNIQUE,
    "assignee" varchar(255),
    "status" varchar(50) NOT NULL,
    CONSTRAINT "issues_pk" PRIMARY KEY ("id")
);

CREATE TABLE "public.groups" (
    "id" serial NOT NULL,
    "chat_id" bigint NOT NULL UNIQUE,
    CONSTRAINT "groups_pk" PRIMARY KEY ("id")
);