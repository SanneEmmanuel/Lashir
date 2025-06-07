FROM python:3.11-slim

# Install system dependencies (including audio/FFmpeg)
RUN apt-get update && \
    apt-get install -y \
    libsndfile1 \
    ffmpeg \
    && rm -rf /var/lib/apt/lists/*

# Configure Poetry
ENV POETRY_VERSION=1.7.1
RUN pip install "poetry==$POETRY_VERSION"
RUN poetry config virtualenvs.create false

WORKDIR /app

# Copy dependency files first (optimizes caching)
COPY pyproject.toml poetry.lock* ./

# Install main dependencies (excluding dev)
RUN poetry install --only main --no-interaction --no-ansi

# Copy application
COPY . .

EXPOSE 10000
CMD ["gunicorn", "-b", "0.0.0.0:10000", "-w", "4", "-k", "uvicorn.workers.UvicornWorker", "lashir.app:app"]
