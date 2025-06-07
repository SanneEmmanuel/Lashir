FROM python:3.11-slim

RUN apt-get update && apt-get install -y libsndfile1

WORKDIR /app
COPY . /app

RUN pip install poetry
RUN poetry install --no-root --only main  # or adjust to your needs

CMD ["gunicorn", "sanne.main:app"]
