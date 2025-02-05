#!/bin/bash

# PostgreSQL connection details
DB_HOST="localhost"
DB_PORT="5432"
DB_USER="postgres"
DB_PASSWORD="your-beatiful-password"
DB_NAME="db_name"

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
NC='\033[0m' # No Color

echo "Starting database setup..."

# Check if psql is installed
if ! command -v psql &> /dev/null; then
    echo -e "${RED}Error: PostgreSQL is not installed${NC}"
    exit 1
fi

# Create database if it doesn't exist
echo "Creating database if it doesn't exist..."
psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d postgres -c "SELECT 'CREATE DATABASE $DB_NAME' WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = '$DB_NAME')\gexec"

if [ $? -ne 0 ]; then
    echo -e "${RED}Error: Failed to create database${NC}"
    exit 1
fi

# Apply schema
echo "Applying database schema..."
psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d $DB_NAME -f schema.sql

if [ $? -ne 0 ]; then
    echo -e "${RED}Error: Failed to apply schema${NC}"
    exit 1
fi

echo -e "${GREEN}Database setup completed successfully!${NC}"

# Create admin user if it doesn't exist
echo "Creating admin user..."
psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d $DB_NAME << EOF
DO \$\$
BEGIN
    IF NOT EXISTS (SELECT FROM users WHERE email = 'admin@hospital.com') THEN
        INSERT INTO users (name, email, password, role, status)
        VALUES (
            'Admin User',
            'admin@hospital.com',
            '\$2a\$10\$YourHashedPasswordHere',  -- Replace with actual hashed password
            'admin',
            true
        );
    END IF;
END
\$\$;
EOF

echo -e "${GREEN}Setup complete!${NC}"
echo "You can now connect to the database using:"
echo "Host: $DB_HOST"
echo "Port: $DB_PORT"
echo "Database: $DB_NAME"
echo "Admin email: admin@hospital.com"