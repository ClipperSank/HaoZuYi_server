CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "indexpage" varchar,
  "username" varchar NOT NULL,
  "role" varchar NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT 'now()',
  "account" varchar NOT NULL,
  "password" varchar NOT NULL,
  "age" integer,
  "birthday" date,
  "contract_count" integer,
  "houses_for_rent" integer,
  "owned_houses" integer
);

CREATE TABLE "user_search_record" (
  "id" bigserial PRIMARY KEY,
  "search_query" varchar NOT NULL,
  "user_id" bigint NOT NULL,
  "search_time" timestamp NOT NULL DEFAULT 'now()'
);

CREATE TABLE "contact" (
  "id" bigserial PRIMARY KEY,
  "renter_id" bigint NOT NULL,
  "landlord_id" bigint NOT NULL,
  "house_id" integer NOT NULL,
  "contract" text NOT NULL,
  "rent" decimal NOT NULL,
  "start_time" timestamp NOT NULL,
  "end_time" timestamp DEFAULT null,
  "renter_review" text DEFAULT null,
  "landlord_review" text DEFAULT null
);

CREATE TABLE "house" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigint NOT NULL,
  "address" varchar NOT NULL,
  "is_renting" integer,
  "price" decimal,
  "size" integer,
  "kitchen" integer,
  "bathroom" integer,
  "sleeping_room" integer,
  "created_at" timestamp DEFAULT 'now()'
);

CREATE TABLE "review" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigint NOT NULL,
  "house_id" bigint NOT NULL,
  "rating" decimal NOT NULL,
  "comment" text NOT NULL,
  "created_at" timestamp DEFAULT 'now()'
);

CREATE TABLE "house_photo" (
  "id" bigserial PRIMARY KEY,
  "house_id" bigint NOT NULL,
  "photo_url" varchar NOT NULL
);

CREATE INDEX ON "users" ("username");

CREATE INDEX ON "users" ("account");

CREATE INDEX ON "user_search_record" ("user_id");

CREATE INDEX ON "contact" ("renter_id");

CREATE INDEX ON "contact" ("landlord_id");

CREATE INDEX ON "contact" ("renter_id", "landlord_id");

COMMENT ON COLUMN "users"."indexpage" IS '主頁索引';

COMMENT ON COLUMN "users"."username" IS '用戶名';

COMMENT ON COLUMN "users"."role" IS '用戶角色';

COMMENT ON COLUMN "users"."created_at" IS '創建時間';

COMMENT ON COLUMN "users"."account" IS '賬號';

COMMENT ON COLUMN "users"."password" IS '密碼';

COMMENT ON COLUMN "users"."age" IS '年齡';

COMMENT ON COLUMN "users"."birthday" IS '生日';

COMMENT ON COLUMN "users"."contract_count" IS '合同數量';

COMMENT ON COLUMN "users"."houses_for_rent" IS '出租房屋數量';

COMMENT ON COLUMN "users"."owned_houses" IS '擁有房屋數量';

COMMENT ON COLUMN "user_search_record"."search_query" IS '搜索查詢';

COMMENT ON COLUMN "user_search_record"."user_id" IS '用戶ID';

COMMENT ON COLUMN "user_search_record"."search_time" IS '搜索時間';

COMMENT ON COLUMN "contact"."renter_id" IS '租客用戶ID';

COMMENT ON COLUMN "contact"."landlord_id" IS '房東用戶ID';

COMMENT ON COLUMN "contact"."house_id" IS '房屋ID';

COMMENT ON COLUMN "contact"."contract" IS '合同詳情';

COMMENT ON COLUMN "contact"."rent" IS '租金金額';

COMMENT ON COLUMN "contact"."start_time" IS '開始時間';

COMMENT ON COLUMN "contact"."end_time" IS '結束時間';

COMMENT ON COLUMN "contact"."renter_review" IS '租客評價';

COMMENT ON COLUMN "contact"."landlord_review" IS '房東評價';

COMMENT ON COLUMN "house"."user_id" IS '用戶ID';

COMMENT ON COLUMN "house"."address" IS '地址';

COMMENT ON COLUMN "house"."is_renting" IS '是否出租';

COMMENT ON COLUMN "house"."price" IS '價格';

COMMENT ON COLUMN "house"."size" IS '面積';

COMMENT ON COLUMN "house"."kitchen" IS '廚房';

COMMENT ON COLUMN "house"."bathroom" IS '浴室';

COMMENT ON COLUMN "house"."sleeping_room" IS '睡房';

COMMENT ON COLUMN "house"."created_at" IS '創建時間';

COMMENT ON COLUMN "review"."user_id" IS '用戶ID';

COMMENT ON COLUMN "review"."house_id" IS '房屋ID';

COMMENT ON COLUMN "review"."rating" IS '評分';

COMMENT ON COLUMN "review"."comment" IS '評論內容';

COMMENT ON COLUMN "review"."created_at" IS '創建時間';

COMMENT ON COLUMN "house_photo"."house_id" IS '房屋ID';

COMMENT ON COLUMN "house_photo"."photo_url" IS '照片URL';

ALTER TABLE "review" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "contact" ADD FOREIGN KEY ("renter_id") REFERENCES "users" ("id");

ALTER TABLE "contact" ADD FOREIGN KEY ("landlord_id") REFERENCES "users" ("id");

ALTER TABLE "house" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "house_photo" ADD FOREIGN KEY ("house_id") REFERENCES "house" ("id");

ALTER TABLE "user_search_record" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
