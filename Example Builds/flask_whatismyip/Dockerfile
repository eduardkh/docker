# Use a minimal Python image for building
FROM python:3.14.0a3-alpine3.21 AS builder

WORKDIR /app

# Install build dependencies
RUN apk add --no-cache gcc musl-dev libffi-dev

# Install Python dependencies
COPY requirements.txt .
RUN pip install --no-cache-dir --prefix=/install -r requirements.txt
RUN pip install --no-cache-dir --prefix=/install gunicorn

# Use a smaller runtime image
FROM python:3.14.0a3-alpine3.21

WORKDIR /app

# Install runtime dependencies
RUN apk add --no-cache libffi

# Copy installed Python dependencies
COPY --from=builder /install /usr/local

# Copy the application code
COPY . .

# Create a non-root user
RUN addgroup -S appgroup && adduser -S appuser -G appgroup
USER appuser

# Expose the port (use an ENV variable for flexibility)
ENV PORT=8000
EXPOSE $PORT

# Command to run the application
CMD ["gunicorn", "-b", "0.0.0.0:8000", "app:app"]
