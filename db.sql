CREATE DATABASE IF NOT EXISTS cursoangular;

USE cursoangular;

CREATE TABLE cursoangular.products(
    id          INT(255) auto_increment not null,
    name        VARCHAR(255),
    description TEXT,
    price       VARCHAR(255),
    image       VARCHAR(255),
    CONSTRAINT pk_products PRIMARY KEY(id)
) ENGINE=InnoDb;