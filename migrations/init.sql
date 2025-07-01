-- Active: 1750428005859@@127.0.0.1@5432@school
-- migrations/init.sql
CREATE DATABASE school;

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(100) NOT NULL,
    role VARCHAR(20) NOT NULL, -- "admin", "teacher", "student"
    verified BOOLEAN NOT NULL DEFAULT FALSE
);

UPDATE users SET verified = true WHERE email = 'admin@example.com';

CREATE TABLE lessons (
    id SERIAL PRIMARY KEY,
    title VARCHAR(100) NOT NULL,
    description TEXT
);

CREATE TABLE schedules (
    id SERIAL PRIMARY KEY,
    lesson_id INTEGER REFERENCES lessons (id),
    teacher_id INTEGER REFERENCES users (id),
    start_time TIMESTAMP,
    end_time TIMESTAMP
);

CREATE TABLE grades (
    id SERIAL PRIMARY KEY,
    student_id INTEGER REFERENCES users (id),
    teacher_id INTEGER REFERENCES users (id),
    lesson_id INTEGER REFERENCES lessons (id),
    value VARCHAR(5),
    date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE attendance (
    id SERIAL PRIMARY KEY,
    student_id INTEGER REFERENCES users (id),
    lesson_id INTEGER REFERENCES lessons (id),
    date DATE,
    present BOOLEAN
);

CREATE TABLE homeworks (
    id SERIAL PRIMARY KEY,
    student_id INTEGER REFERENCES users (id),
    lesson_id INTEGER REFERENCES lessons (id),
    content TEXT,
    submitted_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    grade VARCHAR(5)
);

CREATE TABLE tariffs (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50),
    price DECIMAL
);

CREATE TABLE payments (
    id SERIAL PRIMARY KEY,
    student_id INTEGER REFERENCES users (id),
    tariff_id INTEGER REFERENCES tariffs (id),
    amount DECIMAL,
    date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password TEXT NOT NULL,
    role VARCHAR(50) DEFAULT 'student',
    verified BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE courses (
    id SERIAL PRIMARY KEY,
    title VARCHAR(100) NOT NULL,
    description TEXT,
    teacher_id INTEGER REFERENCES users (id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE student_course_progress (
    id SERIAL PRIMARY KEY,
    student_id INTEGER NOT NULL REFERENCES users (id),
    course_id INTEGER NOT NULL REFERENCES courses (id),
    progress INTEGER NOT NULL DEFAULT 0,
    completed BOOLEAN NOT NULL DEFAULT FALSE,
    UNIQUE (student_id, course_id)
);