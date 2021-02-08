GRANT ALL PRIVILEGES ON DATABASE microservices TO remote;
CREATE TABLE "users"(
    id bigserial not null
    ,login varchar
    ,username varchar
    ,surname varchar
    ,birthday varchar
    ,password varchar
);
INSERT INTO "users"(login,username,surname,birthday,password)
VALUES(
    'admin'
    ,'Administrator'
    ,'Administratorov'
    ,'01.01.0001'
    ,'admin!'
);