CREATE TABLE schedules (
                           id SERIAL PRIMARY KEY,
                           course_id INTEGER NOT NULL REFERENCES courses(id) ON DELETE CASCADE,
                           instructor_id INTEGER NOT NULL REFERENCES instructors(id) ON DELETE CASCADE, -- связь с инструкторами
                           start_time TIMESTAMP NOT NULL,
                           end_time TIMESTAMP NOT NULL,
                           location TEXT,
                           day_of_week TEXT NOT NULL -- день недели (например, "Monday", "Tuesday", и т.д.)
);
