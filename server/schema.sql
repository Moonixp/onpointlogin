/*
CREATE TABLE userids (
    fullname VARCHAR(255),
    last_login_time VARCHAR(255),
    last_login_date VARCHAR(255)
);

ALTER TABLE userids ADD COLUMN id INT AUTO_INCREMENT PRIMARY KEY

INSERT INTO USERS (fullname) VALUES
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
    FOREIGN KEY (id) REFERENCES userids(id),
    time VARCHAR(255),
    date VARCHAR(255)
);

SELECT * from USERS INNER JOIN logintimes ON USERS.id = logintimes.id

select * from userids inner join  logintimes on userids.id = logintimes.id ;
INSERT INTO logintimes (id) VALUES
(1), (2), (3), (4), (5), (6), (7), (8), (9), (10),
(11), (12), (13), (14), (15), (16), (17), (18), (19), (20),
(21), (22), (23), (24),(25),(26);

*/


