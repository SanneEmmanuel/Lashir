FROM python:3.11-slim

# Install runtime system dependencies
RUN apt-get update && \
    apt-get install -y --no-install-recommends \
    libsndfile1 \
    ffmpeg \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app

# Copy only requirement files first for caching
COPY pyproject.toml ./

# Install Poetry and project dependencies (using pyproject.toml's built-in)
RUN pip install --no-cache-dir poetry && \
    poetry install --only main --no-interaction --no-ansi --no-root

# Copy application code
COPY . .

# Security best practices
RUN useradd -m appuser && chown -R appuser:appuser /app
USER appuser

# Runtime config
ENV PYTHONUNBUFFERED=1 \
    PYTHONPATH=/app

EXPOSE 10000
CMD ["gunicorn", "-b", "0.0.0.0:10000", "-w", "4", "-k", "uvicorn.workers.UvicornWorker", "main:app"]
