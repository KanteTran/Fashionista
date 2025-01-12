DROP TABLE IF EXISTS users;
CREATE TABLE users (
                             id VARCHAR(255) PRIMARY KEY,                             -- Auto-incrementing primary key
                             phone VARCHAR(255) UNIQUE,                      -- Unique phone identifier, indexed
                             user_name VARCHAR(100),                            -- User name, nullable
                             birthday DATE,                                     -- Birthday, nullable
                             address VARCHAR(255),                              -- Address, nullable
                             device_id VARCHAR(100),                            -- Device identifier, nullable
                             gender INT CHECK (gender IN (0, 1, 2)),            -- Gender: 0 (male), 1 (female), 2 (other), nullable
                             password_hash VARCHAR(255),                        -- Hashed password, nullable
                             created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,    -- Automatically sets time on creation
                             updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

DROP TABLE IF EXISTS stores;
CREATE TABLE stores (
                      id VARCHAR(255) PRIMARY KEY,
                      phone VARCHAR(15) NOT NULL UNIQUE,
                      store_name VARCHAR(255) NOT NULL,
                      address TEXT NOT NULL,
                      description TEXT,
                      password VARCHAR(255) NOT NULL,
                      status INT NOT NULL,
                      url_image VARCHAR(255) NOT NULL

);

DROP TABLE IF EXISTS items;
CREATE TABLE items (
                       id SERIAL PRIMARY KEY,
                       store_id  VARCHAR(255) NOT NULL,
                       name VARCHAR(255) NOT NULL,
                       url VARCHAR(255) NOT NULL,
                       image_url VARCHAR(255) NOT NULL

                    );

DROP TABLE IF EXISTS inventories;
CREATE TABLE inventories (
                       id SERIAL PRIMARY KEY,
                       store_id  VARCHAR(255) NOT NULL,
                       name VARCHAR(255) NOT NULL,
                       url VARCHAR(255) NOT NULL,
                       image_url VARCHAR(255) NOT NULL,
                        user_id VARCHAR(255) NOT NULL
);