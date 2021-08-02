DROP DATABASE IF EXISTS meteor_assessment;

CREATE DATABASE meteor_assessment;

USE meteor_assessment;

DROP TABLE IF EXISTS family_members;

CREATE TABLE family_members(
id INT UNSIGNED NOT NULL AUTO_INCREMENT,
PRIMARY KEY (id),
household_id INT UNSIGNED NOT NULL,
name VARCHAR(30) NOT NULL,
gender VARCHAR(10) NOT NULL,
marital_status VARCHAR(20) NOT NULL,
spouse VARCHAR(30) NOT NULL,
occupation_type VARCHAR(30) NOT NULL,
annual_income FLOAT NOT NULL,
dob VARCHAR(10));

DROP TABLE IF EXISTS households;

CREATE TABLE households(
id INT UNSIGNED NOT NULL AUTO_INCREMENT,
PRIMARY KEY (id),
housing_type VARCHAR(30) NOT NULL);

INSERT INTO households VALUES(1, 'HDB');
INSERT INTO households VALUES(2, 'HDB');
INSERT INTO households VALUES(3, 'CONDOMINIUM');
INSERT INTO households VALUES(4, 'LANDED');
INSERT INTO households VALUES(5, 'LANDED');
INSERT INTO households VALUES(6, 'HDB');

INSERT INTO family_members VALUES (1, 1, 'Mark Tan', 'MALE', 'SINGLE', '', 'STUDENT', 0.0, '24-12-2011');
INSERT INTO family_members VALUES (2, 1, 'Mavis Tan', 'FEMALE', 'SINGLE', '', 'STUDENT', 0.0, '21-01-2017');
INSERT INTO family_members VALUES (3, 1, 'Tan Jun Han', 'MALE', 'MARRIED', 'Rebecca Chew', 'EMPLOYED', 75000.0, '04-09-1969');
INSERT INTO family_members VALUES (4, 1, 'Rebecca Chew', 'FEMALE', 'MARRIED', 'Tan Jun Han', 'EMPLOYED', 82000.0, '14-11-1971');

INSERT INTO family_members VALUES (5, 2, 'Reuben Wong', 'MALE', 'SINGLE', '', 'STUDENT', 0.0, '17-08-2004');
INSERT INTO family_members VALUES (6, 2, 'Rachael Wong', 'FEMALE', 'SINGLE', '', 'STUDENT', 0.0, '22-05-2006');
INSERT INTO family_members VALUES (7, 2, 'John Wong', 'MALE', 'SINGLE', '', 'EMPLOYED', 70000.0, '02-01-1964');

INSERT INTO family_members VALUES (8, 3, 'Trevor Teo', 'MALE', 'SINGLE', '', 'STUDENT', 0.0, '09-04-1997');
INSERT INTO family_members VALUES (9, 3, 'Therese Teo', 'FEMALE', 'SINGLE', '', 'STUDENT', 0.0, '09-04-1997');
INSERT INTO family_members VALUES (10, 3, 'Russell Teo', 'MALE', 'MARRIED', 'Felicia Tang', 'EMPLOYED', 72000.0, '22-01-1965');
INSERT INTO family_members VALUES (11, 3, 'Felicia Tang', 'FEMALE', 'MARRIED', 'Russell Teo', 'EMPLOYED', 70000.0, '29-03-1967');
INSERT INTO family_members VALUES (12, 3, 'Jade Leow', 'FEMALE', 'SINGLE', '', 'UNEMPLOYED', 0.0, '30-07-1939');

INSERT INTO family_members VALUES (13, 4, 'Amanda Lee', 'FEMALE', 'SINGLE', '', 'UNEMPLOYED', 0.0, '29-10-2000');
INSERT INTO family_members VALUES (14, 4, 'Delilah Wong', 'FEMALE', 'SINGLE', '', 'EMPLOYED', 50000.0, '20-05-1969');

INSERT INTO family_members VALUES (15, 5, 'Molly Baker', 'FEMALE', 'SINGLE', '', 'STUDENT', 0.0, '12-08-2004');
INSERT INTO family_members VALUES (16, 5, 'Christopher Baker', 'MALE', 'MARRIED', 'Germaine Chua', 'EMPLOYED', 48000.0, '22-04-1970');
INSERT INTO family_members VALUES (17, 5, 'Germaine Chua', 'FEMALE', 'MARRIED', 'Christopher Baker', 'EMPLOYED', 70000.0, '02-06-1970');

INSERT INTO family_members VALUES (18, 6, 'Eric Foo', 'MALE', 'SINGLE', '', 'STUDENT', 0.0, '14-02-2006');
INSERT INTO family_members VALUES (19, 6, 'Nicole Foo', 'FEMALE', 'SINGLE', '', 'STUDENT', 0.0, '21-02-2017');
INSERT INTO family_members VALUES (20, 6, 'Nate Foo', 'MALE', 'SINGLE', '', 'EMPLOYED', 47000.0, '04-09-1969');
INSERT INTO family_members VALUES (21, 6, 'Leila Foo', 'FEMALE', 'SINGLE', '', 'EMPLOYED', 60000.0, '14-11-1971');