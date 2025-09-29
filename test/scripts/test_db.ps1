docker run --name postgres-test -e POSTGRES_USER=postgres -e  POSTGRES_PASSWORD=123456 -p 6432:5432 -d postgres:latest

Write-Output "PostgreSql Starting ..."
Start-Sleep -Seconds 10

docker exec -it postgres-test psql -U postgres -d postgres -c "Create Database ecommerce"
Write-Output "Database ecommerce created"

docker exec -it postgres-test psql -U postgres -d ecommerce -c "
    create table if not exists products
    (
        id bigserial not null primary key,
        name varchar(255) not null,
        price double precision not null,
        discount double precision,
        store varchar(255) not null
    );
"
Write-Output "Table products created"