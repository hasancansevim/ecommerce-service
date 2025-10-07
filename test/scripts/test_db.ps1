# database-reset.ps1
docker stop postgres-test
docker rm postgres-test

docker run --name postgres-test -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=123456 -p 6432:5432 -d postgres:latest

Write-Output "PostgreSql Starting ..."
Start-Sleep -Seconds 10

docker exec -i postgres-test psql -U postgres -d postgres -c "DROP DATABASE IF EXISTS ecommerce;"
docker exec -i postgres-test psql -U postgres -d postgres -c "CREATE DATABASE ecommerce;"
Write-Output "Database ecommerce recreated"

docker exec -i postgres-test psql -U postgres -d ecommerce -c "
    CREATE TABLE IF NOT EXISTS products (
        id BIGSERIAL NOT NULL PRIMARY KEY,
        name VARCHAR(255) NOT NULL,
        price DOUBLE PRECISION NOT NULL,
        discount DOUBLE PRECISION,
        store VARCHAR(255) NOT NULL
    );

    CREATE TABLE IF NOT EXISTS users (
        id BIGSERIAL NOT NULL PRIMARY KEY,
        first_name VARCHAR(255) NOT NULL,
        last_name VARCHAR(255) NOT NULL,
        email VARCHAR(255) NOT NULL UNIQUE,
        password VARCHAR(255) NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
    );

    CREATE TABLE IF NOT EXISTS carts (
        id BIGSERIAL NOT NULL PRIMARY KEY,
        user_id BIGINT NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
        FOREIGN KEY (user_id) REFERENCES users(id)
    );

    CREATE TABLE IF NOT EXISTS cart_items (
        id BIGSERIAL NOT NULL PRIMARY KEY,
        cart_id BIGINT NOT NULL,
        product_id BIGINT NOT NULL,
        quantity INT NOT NULL,
        FOREIGN KEY (cart_id) REFERENCES carts(id),
        FOREIGN KEY (product_id) REFERENCES products(id)
    );

    CREATE TABLE IF NOT EXISTS orders (
        id BIGSERIAL NOT NULL PRIMARY KEY,
        user_id BIGINT NOT NULL,
        total_price DOUBLE PRECISION NOT NULL,
        status BOOLEAN NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
        FOREIGN KEY (user_id) REFERENCES users(id)
    );

    CREATE TABLE IF NOT EXISTS order_items (
        id BIGSERIAL NOT NULL PRIMARY KEY,
        order_id BIGINT NOT NULL,
        product_id BIGINT NOT NULL,
        quantity INT NOT NULL,
        price DOUBLE PRECISION NOT NULL,
        FOREIGN KEY (order_id) REFERENCES orders(id),
        FOREIGN KEY (product_id) REFERENCES products(id)
    );
"

Write-Output "All tables recreated successfully"