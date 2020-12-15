create database if not exists  se_test;
use se_test;

-- Make the database
create table if not exists  negativations
(
    id bigint auto_increment
    primary key,
    company_document varchar(55) not null,
    company_name varchar(255) not null,
    customer_document varchar(55) not null,
    value decimal(15,3) not null,
    contract varchar(55) not null,
    debtDate datetime null,
    inclusion_date datetime null
    );

SET FOREIGN_KEY_CHECKS = 0;
INSERT INTO se_test.negativations (id, company_document, company_name, customer_document, value, contract, debtDate, inclusion_date) VALUES (1, '59291534000167', 'ABC S.A.', '51537476467', 1235.230, 'bc063153-fb9e-4334-9a6c-0d069a42065b', '2015-11-13 20:32:51', '2020-11-13 20:32:51');
INSERT INTO se_test.negativations (id, company_document, company_name, customer_document, value, contract, debtDate, inclusion_date) VALUES (2, '77723018000146', '123 S.A.', '51537476467', 400.000, '5f206825-3cfe-412f-8302-cc1b24a179b0', '2015-10-12 20:32:51', '2020-10-12 20:32:51');
INSERT INTO se_test.negativations (id, company_document, company_name, customer_document, value, contract, debtDate, inclusion_date) VALUES (3, '4843574000182', 'DBZ S.A.', '26658236674', 59.990, '3132f136-3889-4efb-bf92-e1efbb3fe15e', '2015-09-11 20:32:51', '2020-09-11 20:32:51');
INSERT INTO se_test.negativations (id, company_document, company_name, customer_document, value, contract, debtDate, inclusion_date) VALUES (4, '23993551000107', 'XPTO S.A.', '62824334010', 230.500, '8b441dbb-3bb4-4fc9-9b46-bdaad00a7a98', '2015-08-10 20:32:51', '2020-08-10 20:32:51');
INSERT INTO se_test.negativations (id, company_document, company_name, customer_document, value, contract, debtDate, inclusion_date) VALUES (5, '70170935000100', 'ASD S.A.', '25124543043', 10340.670, 'd6628a0e-d4dd-4f14-8591-2ddc7f1bbeff', '2015-07-09 20:32:51', '2020-07-09 20:32:51');
SET FOREIGN_KEY_CHECKS = 1;