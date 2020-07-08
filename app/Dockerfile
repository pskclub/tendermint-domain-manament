FROM python:3.7.3-alpine
ENV PYTHONUNBUFFERED 1
ENV WORKDIR /usr/app

RUN apk add --update --upgrade --no-cache bash mariadb-connector-c gmp
RUN apk add --no-cache --virtual .build-deps python3-dev mariadb-dev build-base gmp-dev

WORKDIR ${WORKDIR}
ADD ./requirements.txt .
RUN pip install --upgrade pip
RUN pip install -r requirements.txt
RUN apk del .build-deps
