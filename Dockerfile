FROM python:3.08

RUN apt-get update && apt-get install -y libsndfile1 build-essential

WORKDIR /app
COPY . /app

RUN pip install poetry==1.7.1
RUN poetry install --no-root --only main

EXPOSE 10000
CMD ["gunicorn", "main:app"]
