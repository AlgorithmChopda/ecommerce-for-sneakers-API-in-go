CREATE TABLE admin (
    id INT GENERATED ALWAYS AS IDENTITY,
    username TEXT NOT NULL,
    password TEXT NOT NULL,
    created_at DATE NOT NULL DEFAULT CURRENT_DATE,
    updated_at DATE NOT NULL DEFAULT CURRENT_DATE
);

CREATE TABLE user (
id INT GENERATED ALWAYS AS IDENTITY,
first_name TEXT NOT NULL,
last_name TEXT NOT NULL,
email TEXT NOT NULL,
password TEXT NOT NULL,
created_at DATE NOT NULL DEFAULT CURRENT_DATE,
updated_at DATE NOT NULL DEFAULT CURRENT_DATE
);

CREATE TABLE organizer (
id INT GENERATED ALWAYS AS IDENTITY,
organization TEXT NOT NULL,
detail TEXT NOT NULL,
email TEXT NOT NULL,
password TEXT NOT NULL,
mobile TEXT NOT NULL,
is_verified BOOLEAN NOT NULL DEFAULT FALSE,
created_at DATE NOT NULL DEFAULT CURRENT_DATE,
updated_at DATE NOT NULL DEFAULT CURRENT_DATE
);

CREATE TABLE fundraiser (
id INT GENERATED ALWAYS AS IDENTITY,
title TEXT NOT NULL,
description TEXT NOT NULL,
organizer_id INT,
image_url TEXT NOT NULL,
video_url TEXT NOT NULL,
target_amount DECIMAL(12, 2) NOT NULL,
status TEXT CHECK(status IN ('active', 'inactive', 'banned')),
created_at DATE NOT NULL DEFAULT CURRENT_DATE,
updated_at DATE NOT NULL DEFAULT CURRENT_DATE,

CONSTRAINT fk_organizer_id
FOREIGN KEY(organizer_id)
REFERENCES organizer(id)
ON DELETE CASCADE
);

CREATE TABLE donation (
id INT GENERATED ALWAYS AS IDENTITY,
user_id INT,
fundraiser_id INT,
amount DECIMAL(12, 2) NOT NULL,
is_anonymous BOOLEAN NOT NULL DEFAULT FALSE,
created_at DATE NOT NULL DEFAULT CURRENT_DATE,

CONSTRAINT fk_user_id
FOREIGN KEY(user_id)
REFERENCES user(id)
ON DELETE CASCADE,

CONSTRAINT fk_fundraiser_id
FOREIGN KEY(fundraiser_id)
REFERENCES fundraiser(id)
ON DELETE CASCADE
);