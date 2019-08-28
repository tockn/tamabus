FROM python:2.7
LABEL maintainer="tockn <s.taku.inthebox@gmail.com>"

COPY . /
WORKDIR /
RUN ["make"]
CMD ["python", "python/darknet.py"]
