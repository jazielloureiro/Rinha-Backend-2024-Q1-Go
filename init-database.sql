CREATE TABLE "Account" (
    "Id" INTEGER PRIMARY KEY,
    "Limit" INTEGER NOT NULL,
    "Value" INTEGER NOT NULL,
    CONSTRAINT "MinValue" CHECK ("Value" >= -"Limit")
);

CREATE TABLE "Statement" (
    "Id" SERIAL PRIMARY KEY,
    "Value" INTEGER NOT NULL,
    "Type" CHAR NOT NULL,
    "Description" VARCHAR(10) NOT NULL,
    "Date" TIMESTAMP NOT NULL
);

INSERT INTO "Account" ("Id", "Limit", "Value")
VALUES
    (1, 100000, 0),
    (2, 80000, 0),
    (3, 1000000, 0),
    (4, 10000000, 0),
    (5, 500000, 0);