CREATE TABLE IF NOT EXISTS target_muscles(
    id bigserial PRIMARY KEY,
    muscle_name VARCHAR(255) UNIQUE NOT NULL,
    user_id bigint NOT NULL,
    created_at TIMESTAMP(0) with time zone NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_target_muscles_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);