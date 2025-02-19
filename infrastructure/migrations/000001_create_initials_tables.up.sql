CREATE TABLE contact_type (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL UNIQUE
);

CREATE TABLE document_type (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL UNIQUE
);

CREATE TABLE app_user (
    user_id SERIAL PRIMARY KEY,
    username VARCHAR(50) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    full_name VARCHAR(100) NOT NULL,
    birth_date DATE,
    registration_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE address (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    street VARCHAR(100) NOT NULL,
    number VARCHAR(10),
    complement VARCHAR(50),
    neighborhood VARCHAR(50) NOT NULL,
    city VARCHAR(50) NOT NULL,
    state CHAR(2) NOT NULL,
    zip_code VARCHAR(10) NOT NULL,
    country VARCHAR(50) DEFAULT 'Brazil',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by INT NOT NULL,
    updated_by INT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES app_user(user_id),
    FOREIGN KEY (created_by) REFERENCES app_user(user_id),
    FOREIGN KEY (updated_by) REFERENCES app_user(user_id)
);

CREATE TABLE contact (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    contact_type_id INT NOT NULL,
    value VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by INT NOT NULL,
    updated_by INT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES app_user(user_id),
    FOREIGN KEY (contact_type_id) REFERENCES contact_type(id),
    FOREIGN KEY (created_by) REFERENCES app_user(user_id),
    FOREIGN KEY (updated_by) REFERENCES app_user(user_id)
);

CREATE TABLE document (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    document_type_id INT NOT NULL,
    value VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by INT NOT NULL,
    updated_by INT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES app_user(user_id),
    FOREIGN KEY (document_type_id) REFERENCES document_type(id),
    FOREIGN KEY (created_by) REFERENCES app_user(user_id),
    FOREIGN KEY (updated_by) REFERENCES app_user(user_id)
);