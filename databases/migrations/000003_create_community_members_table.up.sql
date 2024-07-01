CREATE TABLE IF NOT EXISTS community_members (
    community_id UUID REFERENCES communities(id),
    user_id UUID NOT NULL,
    joined_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (community_id, user_id)
);