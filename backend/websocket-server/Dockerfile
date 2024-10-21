FROM python:3.13-alpine

WORKDIR /app

RUN pip install poetry

COPY . /app
RUN poetry install

CMD [ "poetry", "run", "python", "main.py" ]
