# PostgreSQL Setup for 3x-ui

This document explains how to set up PostgreSQL for the 3x-ui application.

## Prerequisites

1. PostgreSQL server installed and running
2. A database created for the application
3. A user with appropriate permissions

## Environment Variables

The application uses the following environment variables for PostgreSQL configuration:

### Option 1: Using a complete DSN string
```bash
export XUI_POSTGRES_DSN="host=localhost port=5432 user=xui password=xui dbname=xui sslmode=disable"
```

### Option 2: Using individual environment variables
```bash
export XUI_POSTGRES_HOST="localhost"
export XUI_POSTGRES_PORT="5432"
export XUI_POSTGRES_USER="xui"
export XUI_POSTGRES_PASSWORD="xui"
export XUI_POSTGRES_DB="xui"
export XUI_POSTGRES_SSLMODE="disable"
```

## Database Setup

1. **Create the database:**
   ```sql
   CREATE DATABASE xui;
   ```

2. **Create a user (optional, you can use an existing user):**
   ```sql
   CREATE USER xui WITH PASSWORD 'xui';
   GRANT ALL PRIVILEGES ON DATABASE xui TO xui;
   ```

3. **Connect to the database and grant schema permissions:**
   ```sql
   \c xui
   GRANT ALL ON SCHEMA public TO xui;
   ```

## Application Changes

The following changes have been made to support PostgreSQL:

1. **Database Driver**: Changed from SQLite to PostgreSQL driver
2. **Connection**: Uses PostgreSQL connection string instead of file path
3. **Database Operations**: 
   - File-based backup/restore is not supported
   - Use `pg_dump` and `pg_restore` for database operations
4. **Configuration**: Added PostgreSQL-specific environment variables

## Database Backup and Restore

Since PostgreSQL doesn't use file-based databases like SQLite, you need to use PostgreSQL tools:

### Backup
```bash
pg_dump -h localhost -p 5432 -U xui -d xui > backup.sql
```

### Restore
```bash
psql -h localhost -p 5432 -U xui -d xui < backup.sql
```

## Migration from SQLite

If you're migrating from SQLite to PostgreSQL:

1. **Export data from SQLite:**
   ```bash
   sqlite3 x-ui.db .dump > sqlite_backup.sql
   ```

2. **Convert the SQLite dump to PostgreSQL format** (you may need to manually adjust some SQL syntax)

3. **Import to PostgreSQL:**
   ```bash
   psql -h localhost -p 5432 -U xui -d xui < converted_backup.sql
   ```

## Troubleshooting

### Connection Issues
- Ensure PostgreSQL is running: `sudo systemctl status postgresql`
- Check if the database exists: `psql -l`
- Verify user permissions: `psql -U xui -d xui -c "\du"`

### Permission Issues
- Make sure the user has proper permissions on the database
- Check PostgreSQL logs: `tail -f /var/log/postgresql/postgresql-*.log`

### SSL Issues
- Set `XUI_POSTGRES_SSLMODE=disable` for development
- For production, configure proper SSL certificates

## Security Considerations

1. **Use strong passwords** for the database user
2. **Enable SSL** in production environments
3. **Restrict network access** to the PostgreSQL server
4. **Use environment variables** instead of hardcoded credentials
5. **Regular backups** using `pg_dump`

## Performance Tuning

For better performance, consider:

1. **Connection pooling** (e.g., using PgBouncer)
2. **Proper indexing** on frequently queried columns
3. **Regular VACUUM** and ANALYZE operations
4. **Appropriate memory settings** in postgresql.conf 