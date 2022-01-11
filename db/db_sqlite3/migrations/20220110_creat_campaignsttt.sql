
-- for customization Begin By Nassim

CREATE TABLE IF NOT EXISTS "campaigns_ttt" ("id" integer primary key autoincrement,"user_id" bigint,"name" varchar(255) NOT NULL ,"created_date" datetime,"completed_date" datetime,"page_id" bigint,"status" varchar(255),"url" varchar(255) , "smtp_id" bigint, "launch_date" DATETIME, send_by_date DATETIME);


CREATE TABLE IF NOT EXISTS "template_groups_ttt" ("id" integer primary key autoincrement,"campaign_id" bigint,"template_id" bigint,"group_id" bigint);
CREATE TABLE IF NOT EXISTS "mail_logs_ttt" (
    "id" integer primary key autoincrement,
    "campaign_id" integer,
    "template_d" interger,
    "user_id" integer,
    "send_date" datetime,
    "send_attempt" integer,
    "r_id" varchar(255),
    "processing" boolean);

