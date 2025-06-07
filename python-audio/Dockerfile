FROM python:3.9
RUN apt-get update && apt-get install -y lilypond
RUN pip install music21
WORKDIR /app
COPY notate.py .
CMD ["python3", "notate.py"]
