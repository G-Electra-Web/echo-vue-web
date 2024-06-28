-- Table: users
CREATE TABLE users (
    user_id SERIAL PRIMARY KEY,
    username VARCHAR(100) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    full_name VARCHAR(255),
    join_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Table: members
CREATE TABLE members (
    member_id SERIAL PRIMARY KEY,
    user_id INT NOT NULL UNIQUE REFERENCES users(user_id),
    reg_num VARCHAR(225) NOT NULL UNIQUE,
    membership_level VARCHAR(50),
    join_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Table: core_members
CREATE TABLE core_members (
    core_member_id SERIAL PRIMARY KEY,
    user_id INT NOT NULL UNIQUE REFERENCES users(user_id),
    role VARCHAR(100),
    join_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Table: staff
CREATE TABLE staff (
    staff_id SERIAL PRIMARY KEY,
    user_id INT NOT NULL UNIQUE REFERENCES users(user_id),
    role VARCHAR(100),
    department VARCHAR(100),
    hire_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Table: events
CREATE TABLE events (
    event_id SERIAL PRIMARY KEY,
    event_name VARCHAR(255) NOT NULL,
    event_date DATE,
    description TEXT,
    location VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Table: attendees
CREATE TABLE attendees (
    user_id INT NOT NULL REFERENCES users(user_id),
    event_id INT NOT NULL REFERENCES events(event_id),
    PRIMARY KEY (user_id, event_id)
);

-- Table: gallery
CREATE TABLE gallery (
    image_id SERIAL PRIMARY KEY,
    event_id INT REFERENCES events(event_id),
    user_id INT REFERENCES users(user_id),
    image_url VARCHAR(255) NOT NULL,
    description TEXT,
    uploaded_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Table: comments
CREATE TABLE comments (
    comment_id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(user_id),
    event_id INT REFERENCES events(event_id),
    comment_text TEXT,
    commented_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Table: projects
CREATE TABLE projects (
    project_id SERIAL PRIMARY KEY,
    project_name VARCHAR(255) NOT NULL,
    description TEXT,
    status VARCHAR(50) DEFAULT 'Pending', -- Pending, Approved, Rejected
    created_by INT NOT NULL REFERENCES users(user_id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Table: project_requests
CREATE TABLE project_requests (
    request_id SERIAL PRIMARY KEY,
    project_id INT NOT NULL REFERENCES projects(project_id),
    user_id INT NOT NULL REFERENCES users(user_id),
    status VARCHAR(50) DEFAULT 'Pending', -- Pending, Approved, Rejected
    requested_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Table: project_members
CREATE TABLE project_members (
    project_id INT NOT NULL REFERENCES projects(project_id),
    user_id INT NOT NULL REFERENCES users(user_id),
    PRIMARY KEY (project_id, user_id)
);

-- Table: notices
CREATE TABLE notices (
    notice_id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    content TEXT,
    created_by INT NOT NULL REFERENCES users(user_id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

