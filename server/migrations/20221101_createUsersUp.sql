  CREATE TABLE users (
    id serial not null primary key,
    email varchar not null unique,
    encrypted_password varchar not null,
    userName varchar not null,
    seccondName varchar not null,
    isadmin boolean DEFAULT false
  );
