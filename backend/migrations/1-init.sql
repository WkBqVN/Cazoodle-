-- +migrate Up
CREATE SCHEMA IF NOT EXISTS test_survey;
-- Drop
DROP TABLE IF EXISTS test_survey.survey;
-- stock
CREATE TABLE test_survey.surve
(
    stock_id     SERIAL,
    stock_name   Char(20) not null, --limit at 20 char
    stock_price  int, -- small calculate at test
    last_update  timestamp,
    PRIMARY KEY (stock_id)
);

INSERT INTO testdata.stocks(stock_name, stock_price, last_update)
VALUES ('stock_1', 111, current_timestamp);

INSERT INTO testdata.stocks(stock_name, stock