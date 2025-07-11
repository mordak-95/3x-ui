 # PostgreSQL Setup Guide for 3x-ui

This guide explains how to configure PostgreSQL for use with the 3x-ui application.

## Prerequisites

- PostgreSQL server installed and running
- A database created for the application
- A database user with appropriate permissions

## Installation Steps

### 1. Install PostgreSQL

```bash
# Ubuntu/Debian
sudo apt update
sudo apt install postgresql postgresql-contrib

# CentOS/RHEL
sudo yum install postgresql postgresql-server
sudo postgresql-setup initdb
sudo systemctl start postgresql
sudo systemctl enable postgresql
```

### 2. Create Database and User

```bash
# Switch to postgres user
sudo -u postgres psql

# Create database and user
CREATE DATABASE xui;
CREATE USER xui WITH PASSWORD 'your_password';
GRANT ALL PRIVILEGES ON DATABASE xui TO xui;
\c xui
GRANT ALL ON SCHEMA public TO xui;
\q
```

## Configure 3x-ui Service

### 3. Edit Service File

After installing 3x-ui, edit the `x-ui.service` file with the following command:

```bash
sudo nano /etc/systemd/system/x-ui.service
```

Add the following environment variables in the `[Service]` section:

```ini
[Service]
# ... existing service configuration ...
Environment="XUI_POSTGRES_HOST=localhost"
Environment="XUI_POSTGRES_PORT=5432"
Environment="XUI_POSTGRES_USER=xui"
Environment="XUI_POSTGRES_PASSWORD=your_password"
Environment="XUI_POSTGRES_DB=xui"
# ... rest of service configuration ...
```

### 4. Set Environment Variables

Set the Ubuntu environment variables with the following commands:

```bash
export XUI_POSTGRES_HOST="localhost"
export XUI_POSTGRES_PORT="5432"
export XUI_POSTGRES_USER="xui"
export XUI_POSTGRES_PASSWORD="your_password"
export XUI_POSTGRES_DB="xui"
```

### 5. Update Service

```bash
# Reload systemd daemon
sudo systemctl daemon-reload

# Restart x-ui service
sudo systemctl restart x-ui

# Check service status
sudo systemctl status x-ui
```

## Backup and Restore

### Backup

```bash
# Create backup
pg_dump -h localhost -p 5432 -U xui -d xui > backup.sql
```

### Restore

```bash
# Restore from backup file
psql -h localhost -p 5432 -U xui -d xui < backup.sql
```

## Troubleshooting

### Common Issues

#### Connection Refused
```bash
# Check PostgreSQL status
sudo systemctl status postgresql

# Check port
sudo netstat -tlnp | grep 5432
```

#### Authentication Failed
```bash
# Check users
sudo -u postgres psql -c "\du"

# Change password
sudo -u postgres psql -c "ALTER USER xui PASSWORD 'new_password';"
```

#### Service Won't Start
```bash
# Check logs
sudo journalctl -u x-ui -f
```

---

**Note**: Replace `your_password` with a strong, unique password.