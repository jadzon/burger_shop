FROM postgres:15-alpine

# Set environment variables for PostgreSQL
ENV POSTGRES_DB=myapp
ENV POSTGRES_USER=postgres
ENV POSTGRES_PASSWORD=password

# Copy initialization scripts (optional)
# COPY ./init.sql /docker-entrypoint-initdb.d/

# Expose the PostgreSQL port
EXPOSE 5432

# The postgres image already has a default CMD, so we don't need to specify one