-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users
(
    user_id serial not null unique,
    passport_number VARCHAR(255) NOT NULL,
    surname VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    patronymic VARCHAR(255),
    address VARCHAR(255) NOT NULL
);


CREATE TABLE IF NOT EXISTS tasks
(
    task_id serial not null unique,
    is_completed bool not null,
    title varchar(255) not null,
    user_id integer,
    started_at timestamp with time zone NOT NULL,
    finished_at timestamp with time zone,
    duration interval,
    FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
);


CREATE TABLE IF NOT EXISTS info
(
    info_id serial not null unique,
    passport_series integer NOT NULL,
    passport_number integer NOT NULL,
    surname VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    patronymic VARCHAR(255),
    address VARCHAR(255) NOT NULL
);


INSERT INTO info (passport_series, passport_number, surname, name, patronymic, address)
VALUES
(1234, 123456, 'Иванов','Иван', 'Иванович', 'Москва, ул.Пушкина'),
(2222, 333444, 'Сидоров','Василий', 'Петрович', 'Нижний Новгород, уп.Победы'),
(5588, 111222, 'Михайлов','Джек', 'Юрьевич', 'Новосибирск, ул.Спортивная');

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS tasks;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS info;
-- +goose StatementEnd
