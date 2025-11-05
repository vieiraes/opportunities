-- Database initialization script

CREATE TABLE IF NOT EXISTS opportunities (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    company VARCHAR(255),
    location VARCHAR(255),
    salary DECIMAL(10, 2),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Indexes
CREATE INDEX idx_opportunities_company ON opportunities(company);
CREATE INDEX idx_opportunities_created_at ON opportunities(created_at);

-- Insert sample data
INSERT INTO opportunities (title, description, company, location, salary) VALUES
('Backend Developer', 'Desenvolvedor Go Sênior', 'Tech Corp', 'São Paulo, SP', 12000.00),
('DevOps Engineer', 'Especialista em Docker e Kubernetes', 'Cloud Solutions', 'Remote', 15000.00);