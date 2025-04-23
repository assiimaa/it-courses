CREATE TABLE payments (
                          id SERIAL PRIMARY KEY,
                          user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
                          course_id INTEGER NOT NULL REFERENCES courses(id) ON DELETE CASCADE,
                          amount NUMERIC(10, 2) NOT NULL,
                          paid_at TIMESTAMP DEFAULT now()
);
