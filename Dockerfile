FROM python:3.11-slim

# Install system dependencies
RUN apt-get update && \
    apt-get install -y \
    libsndfile1 \
    ffmpeg && \
    rm -rf /var/lib/apt/lists/*

# Set up poetry
ENV PYTHONFAULTHANDLER=1 \
    PYTHONUNBUFFERED=1 \
    PYTHONHASHSEED=random \
    PIP_NO_CACHE_DIR=off \
    PIP_DISABLE_PIP_VERSION_CHECK=on \
    PIP_DEFAULT_TIMEOUT=100 \
    POETRY_VERSION=1.7.1

RUN pip install poetry==$POETRY_VERSION

WORKDIR /app

# Copy only pyproject.toml first
COPY pyproject.toml ./

# Generate poetry.lock (temporary workaround)
RUN poetry lock --no-update

# Now install dependencies
RUN poetry config virtualenvs.create false && \
    poetry install --only main --no-interaction --no-ansi

# Copy the rest of the application
COPY . .

EXPOSE 10000

CMD ["gunicorn", "-b", "0.0.0.0:10000", "-w", "4", "-k", "uvicorn.workers.UvicornWorker", "lashir.app:app"]
