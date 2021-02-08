FROM postgres:latest
ENV POSTGRES_USER remote
ENV POSTGRES_PASSWORD Cfyz11005310
ENV POSTGRES_DB microservices
COPY init.sql /docker-entrypoint-initdb.d/
WORKDIR /home/remote/Go/src/Microservices/postgres/  
