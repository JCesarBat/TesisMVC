CREATE TABLE  "files"(
    "ID" Bigserial PRIMARY KEY ,
    "prov_id" Bigint NOT NULL,
    "file_name" varchar(255) NOT NULL,
    "file_type" varchar(255) NOT NULL,
    "file_path" varchar(255) NOT NULL
);
ALTER TABLE "files" ADD FOREIGN KEY  ("prov_id") REFERENCES "provincia" ("id");