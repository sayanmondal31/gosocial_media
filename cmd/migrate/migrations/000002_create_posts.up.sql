

CREATE TABLE IF NOT EXISTS posts(
     id bigserial PRIMARY KEY,
     content TEXT,
     title TEXT NOT NULL, 
     user_id BIGINT NOT NULL , 
     create_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
     update_at timestamp(0) with time zone NOT NULL DEFAULT NOW() 
);