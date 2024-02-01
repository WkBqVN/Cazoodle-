-- +migrate Up
CREATE SCHEMA IF NOT EXISTS test_survey;
-- Drop
DROP TABLE IF EXISTS test_survey.surveys;
DROP TABLE IF EXISTS test_survey.form_templates;
DROP TABLE IF EXISTS test_survey.client;

DROP TABLE IF EXISTS test_survey.datas;

CREATE TABLE test_survey.form_templates
(
    id     SERIAL,
    form_template jsonb,
    PRIMARY KEY (id)
);

CREATE TABLE test_survey.surveys
(
    id     SERIAL,
    form_template_id int REFERENCES test_survey.form_templates(id) ON DELETE CASCADE,
    PRIMARY KEY (id)
);

CREATE TABLE test_survey.clients
(
    id SERIAL,
    survey_id int REFERENCES test_survey.surveys(id) ON DELETE CASCADE,
    PRIMARY KEY (id)
);

CREATE TABLE test_survey.datas
(
    id     SERIAL,
    form_template_id int REFERENCES test_survey.form_templates(id) ON DELETE CASCADE,
    form_data jsonb,
    PRIMARY KEY (id)
);

INSERT INTO test_survey.form_templates(form_template)
VALUES('[{"email":"test"}]');
INSERT INTO test_survey.surveys(form_template_id)
VALUES(1);
INSERT INTO test_survey.clients(survey_id)
VALUES(1);