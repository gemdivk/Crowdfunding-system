CREATE TABLE "UserPoints" (
                             user_id VARCHAR PRIMARY KEY,
                             points INTEGER DEFAULT 0,
                             created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                             updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "Achievements" (
                             id SERIAL PRIMARY KEY,
                             user_id VARCHAR,
                             achievement_name VARCHAR,
                             achieved_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                             FOREIGN KEY (user_id) REFERENCES "UserPoints" (user_id)
);