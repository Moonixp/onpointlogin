CREATE DATABASE onpointusers;    
USE onpointusers;

CREATE TABLE users (
    fullname VARCHAR(255),
    last_login_time VARCHAR(255),
    last_login_date VARCHAR(255)
);

ALTER TABLE users ADD COLUMN id INT AUTO_INCREMENT PRIMARY KEY;

INSERT INTO users (fullname) VALUES
    ('Oh Hanie'),
    ('Marfo Philip Kwarteng Ansah'),
    ('Katakyie Kojo Desu Achempong'),
    ('Okutu Benjamin'),
    ('Anim-Ofori Calvin'),
    ('Desu emmanuella Konadu'),
    ('Armah Richard'),
    ('Puni Portia'),
    ('Gordon Joseph Abeiku'),
    ('Ampomah Russell Hagan'),
    ('Eunice Yeboah'),
    ('Solomon Ankrah'),
    ('Peter Asamoah'),
    ('Jeffrey Amponsah'),
    ('Emmanuel Affum'),
    ('Andoh Solomon'),
    ('OH Hyun Wo'),
    ('Gideon Ackney'),
    ('Mathias Lawson'),
    ('Jose Litoro'),
    ('Beatrice'),
    ('Christabel'),
    ('Adams'),
    ('Dennis Kwame Rabbi'),
    ('Amadi Victor Uchechukwu'),
    ('Okutu Bright Teye'),
    ('Arday Sandra Naa Aduarh');

CREATE TABLE logintimes (
    id INT NOT NULL,
    FOREIGN KEY (id) REFERENCES users(id),
    time VARCHAR(255),
    date VARCHAR(255)
);

SELECT * from users INNER JOIN logintimes ON users.id = logintimes.id;

select * from users inner join  logintimes on users.id = logintimes.id ;

