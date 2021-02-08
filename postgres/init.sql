GRANT ALL PRIVILEGES ON DATABASE microservices TO remote;
CREATE TABLE users(
    id bigserial not null
    ,login varchar
    ,name varchar
    ,surname varchar
    ,birthday varchar
    ,password varchar
);
INSERT INTO users (login,name,surname,birthday,password)
VALUES(
    'admin'
    ,'Administrator'
    ,'Administratorov'
    ,'01.01.0001'
    ,'admin!'
);