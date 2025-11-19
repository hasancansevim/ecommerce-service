docker stop postgres-test
docker rm postgres-test

docker run --name postgres-test -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=123456 -p 6432:5432 -d postgres:latest
Write-Output "PostgreSql Starting ..."
Start-Sleep -Seconds 5

docker exec -i postgres-test psql -U postgres -d postgres -c "DROP DATABASE IF EXISTS ecommerce;"
docker exec -i postgres-test psql -U postgres -d postgres -c "CREATE DATABASE ecommerce;"
Write-Output "Database ecommerce recreated"

docker exec -i postgres-test psql -U postgres -d ecommerce -c "
    CREATE TABLE IF NOT EXISTS categories(
        id BIGSERIAL NOT NULL PRIMARY KEY,
        name VARCHAR(255) NOT NULL UNIQUE,
        description TEXT,
        is_active BOOLEAN DEFAULT true
    );

    CREATE TABLE IF NOT EXISTS stores(
        id BIGSERIAL NOT NULL PRIMARY KEY,
        name VARCHAR(255) NOT NULL UNIQUE,
        slug VARCHAR(255) NOT NULL UNIQUE,
        description TEXT,
        logo_url VARCHAR(500),
        contact_email VARCHAR(255),
        contact_phone VARCHAR(50),
        contact_address TEXT,
        is_active BOOLEAN DEFAULT true,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
    );

    CREATE TABLE IF NOT EXISTS products (
        id BIGSERIAL NOT NULL PRIMARY KEY,
        name VARCHAR(255) NOT NULL,
        slug VARCHAR(255) NOT NULL UNIQUE,
        description TEXT,
        price DECIMAL(10,2) NOT NULL,
        base_price DECIMAL(10,2) NOT NULL,
        discount DECIMAL(10,2) DEFAULT 0,
        image_url VARCHAR(500),
        meta_description VARCHAR(300),
        stock_quantity INTEGER DEFAULT 0 NOT NULL,
        is_active BOOLEAN DEFAULT true NOT NULL,
        is_featured BOOLEAN DEFAULT false NOT NULL,
        category_id BIGINT,
        store_id BIGINT NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
        FOREIGN KEY (category_id) REFERENCES categories(id),
        FOREIGN KEY (store_id) REFERENCES stores(id)
    );

    CREATE TABLE IF NOT EXISTS users (
        id BIGSERIAL NOT NULL PRIMARY KEY,
        first_name VARCHAR(255) NOT NULL,
        last_name VARCHAR(255) NOT NULL,
        email VARCHAR(255) NOT NULL UNIQUE,
        password_hash VARCHAR(255) NOT NULL,
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