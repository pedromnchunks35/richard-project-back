CREATE TABLE Agent (
    Id BIGINT PRIMARY KEY,
    Name VARCHAR(255)
);

CREATE TABLE Detail (
    Id BIGINT PRIMARY KEY,
    Name VARCHAR(255),
    Value VARCHAR(255),
    Description VARCHAR(255),
    id_agent BIGINT,
    FOREIGN KEY (id_agent) REFERENCES Agent(Id)
);

CREATE TABLE Category (
    Id BIGINT PRIMARY KEY,
    Name VARCHAR(255)
);

CREATE TABLE Image (
    Id BIGINT PRIMARY KEY,
    Name VARCHAR(255),
    Base64 TEXT
);

CREATE TABLE Product (
    Id BIGINT PRIMARY KEY,
    Name VARCHAR(255),
    id_category BIGINT,
    id_image BIGINT,
    FOREIGN KEY (id_category) REFERENCES Category(Id),
    FOREIGN KEY (id_image) REFERENCES Image(Id)
);

CREATE TABLE ProductDetail (
    Id BIGINT PRIMARY KEY,
    id_detail BIGINT,
    id_product BIGINT,
    FOREIGN KEY (id_detail) REFERENCES Detail(Id),
    FOREIGN KEY (id_product) REFERENCES Product(Id)
);

CREATE TABLE Intervention (
    Id BIGINT PRIMARY KEY,
    Name VARCHAR(255),
    Description VARCHAR(255)
);

CREATE TABLE Form (
    Id BIGINT PRIMARY KEY,
    Name VARCHAR(255),
    Company VARCHAR(255),
    NIF INTEGER,
    PhoneNumber VARCHAR(50),
    Email VARCHAR(255),
    FiscalAddress VARCHAR(255),
    Address VARCHAR(255),
    WithDelivery BOOLEAN,
    Observations TEXT,
    id_intervention BIGINT,
    id_form_products BIGINT,
    FOREIGN KEY (id_intervention) REFERENCES Intervention(Id)
);

CREATE TABLE FormProducts (
    Id BIGINT PRIMARY KEY,
    id_form BIGINT,
    id_product BIGINT,
    id_image_attachment BIGINT,
    Quantity FLOAT,
    x FLOAT,
    y FLOAT,
    z FLOAT,
    FOREIGN KEY (id_form) REFERENCES Form(Id),
    FOREIGN KEY (id_product) REFERENCES Product(Id),
    FOREIGN KEY (id_image_attachment) REFERENCES Image(Id)
);
