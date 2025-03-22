CREATE TABLE records (
    id TEXT NOT NULL PRIMARY KEY,
    day TEXT NOT NULL,
    blood_pressure TEXT,
    weight REAL,
    urination_count INTEGER,
    sodium REAL,
    liquid REAL,
    zone INTEGER,
    created_time TEXT NOT NULL,
    last_modified_time TEXT NOT NULL
);
