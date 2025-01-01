FROM python:3.9-slim as builder

WORKDIR /app

RUN apt-get update && apt-get install -y gcc

COPY requirements.txt .

RUN pip install --prefix=/install -r requirements.txt

FROM python:3.9-slim

WORKDIR /app

COPY . .

COPY --from=builder /install /usr/local

EXPOSE 5000

CMD ["python", "app.py"]