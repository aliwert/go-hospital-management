# Hospital Management System

A modern hospital management system built with Go, providing a robust API for managing hospital operations including patient records, appointments, doctors, departments and more.

## Features

- 🏥 **Department Management**

  - Create and manage hospital departments
  - Assign head doctors
  - Track department capacity and schedules

- 👨‍⚕️ **Doctor Management**

  - Doctor profiles and specializations
  - Schedule management
  - Availability tracking
  - Performance metrics

- 👤 **Patient Management**

  - Patient registration and profiles
  - Medical history
  - Appointment scheduling
  - Insurance information

- 📅 **Appointment System**

  - Schedule appointments
  - Manage cancellations
  - Payment tracking
  - Automated notifications

- 📝 **Medical Records**

  - Digital record keeping
  - Test results management
  - Prescription tracking
  - Treatment history

- 🔒 **Security**
  - Role-based access control
  - JWT authentication
  - Secure password handling
  - API rate limiting

## Tech Stack

- **Backend**: Go (Fiber framework)
- **Database**: PostgreSQL
- **ORM**: GORM
- **Authentication**: JWT
- **File Storage**: AWS S3
- **Containerization**: Docker

## Prerequisites

- Go 1.21 or higher
- PostgreSQL 15 or higher
- Docker and Docker Compose
- AWS Account (for S3 storage)

## Environment Variables

Create a `.env` file in the root directory:

```env
# Application
APP_PORT=8080
JWT_SECRET=your-secret-key

# Database
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your-beatiful-password
DB_NAME=hospital_db

# AWS Configuration
AWS_REGION=your-region
AWS_ACCESS_KEY_ID=your-access-key
AWS_SECRET_ACCESS_KEY=your-secret-key
AWS_BUCKET_NAME=your-bucket-name
```

# Installation

1. Clone the repository:

```bash
git clone https://github.com/your-username/go-hospital-management.git
cd go-hospital-management
```

2. Install dependencies:

```bash
go mod download
```

3. Set up the database:

```bash
./scripts/setup-db.sh
```

4. Run locally:

```bash
go run cmd/main.go
```

5. (Optional) or run with docker:

```bash
docker-compose up --build
```

## API Documentation

- **Authentication**

  - POST /api/v1/auth/register - Register new user
  - POST /api/v1/auth/login - User login
  - POST /api/v1/auth/refresh - Refresh JWT token
