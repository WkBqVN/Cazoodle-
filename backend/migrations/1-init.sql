-- +migrate Up
CREATE SCHEMA IF NOT EXISTS test_survey;
-- Drop
DROP TABLE IF EXISTS test_survey.surveys;
DROP TABLE IF EXISTS test_survey.forms;
DROP TABLE IF EXISTS test_survey.client;

CREATE TABLE test_survey.surveys
(
    id     SERIAL,
    forms_id      int,
    PRIMARY KEY (id)
);

CREATE TABLE test_survey.forms
(
    id     SERIAL,
    form_data     jsonb,
    PRIMARY KEY (id)
);

CREATE TABLE test_survey.clients
(
    id SERIAL,
    PRIMARY KEY (id)
);

INSERT INTO test_survey.forms(form_data)
VALUES 
('{"email":"text","age":"int"}'),
('{"email":"text","gender":"bool"}');
INSERT INTO test_survey.surveys(forms_id)
VALUES(1);