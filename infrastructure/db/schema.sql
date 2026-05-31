-- Custom Types
CREATE TYPE support_group_type AS ENUM ('standard', 'chronic');
CREATE TYPE organizer_status_type AS ENUM ('active', 'probation', 'forbidden');
CREATE TYPE meeting_status_type AS ENUM ('scheduled', 'active', 'grace_period', 'closed');

-- Users Table
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    username TEXT UNIQUE NOT NULL,
    email TEXT UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    organizer_status organizer_status_type DEFAULT 'active',
    probation_counter INTEGER DEFAULT 0,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Chronic Cohorts (Private groups)
CREATE TABLE chronic_cohorts (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    description TEXT,
    creator_id UUID REFERENCES users(id),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Cohort Members (For chronic cohorts)
CREATE TABLE cohort_members (
    cohort_id UUID REFERENCES chronic_cohorts(id) ON DELETE CASCADE,
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    role TEXT DEFAULT 'member', -- 'permanent' status mentioned in plan
    PRIMARY KEY (cohort_id, user_id)
);

-- Meetings Table
CREATE TABLE meetings (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    short_id CHAR(8) UNIQUE NOT NULL, -- 8-character alphanumeric unique index
    title TEXT NOT NULL,
    cohort_id UUID REFERENCES chronic_cohorts(id), -- NULL for standard groups
    creator_id UUID REFERENCES users(id),
    status meeting_status_type DEFAULT 'scheduled',
    meeting_type support_group_type DEFAULT 'standard',
    scheduled_at TIMESTAMP WITH TIME ZONE,
    started_at TIMESTAMP WITH TIME ZONE,
    ended_at TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Meeting Participants (Ephemeral aliases)
CREATE TABLE meeting_participants (
    meeting_id UUID REFERENCES meetings(id) ON DELETE CASCADE,
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    alias TEXT NOT NULL,
    is_organizer BOOLEAN DEFAULT FALSE,
    joined_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (meeting_id, user_id)
);

-- Safety Cases
CREATE TABLE safety_cases (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    meeting_id UUID REFERENCES meetings(id),
    reporter_id UUID REFERENCES users(id),
    accused_id UUID REFERENCES users(id),
    round_index INTEGER NOT NULL,
    s3_escrow_path TEXT, -- Path to 10s chunks if escalated
    status TEXT DEFAULT 'pending', -- 'pending', 'resolved', 'dismissed'
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Organizer Ratings Tally (Anonymized)
CREATE TABLE organizer_ratings_tally (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    organizer_id UUID REFERENCES users(id),
    meeting_id UUID REFERENCES meetings(id),
    rating INTEGER CHECK (rating >= 1 AND rating <= 5),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
