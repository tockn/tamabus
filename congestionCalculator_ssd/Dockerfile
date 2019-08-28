FROM yoanlin/opencv-python3:jessie
LABEL maintainer="tockn <s.taku.inthebox@gmail.com>"

RUN pip install --upgrade pip
RUN pip install cython \
 && pip install six \
 && pip install Pillow \
 && pip install tensorflow \
 && pip install dlib \
 && pip install face_recognition \
 && pip install pandas \
 && rm -rf ~/.cache/pip
RUN pip install keras numpy imageio mysqlclient schedule
ENV DOCKERIZE_VERSION v0.6.0
RUN apt-get install -y wget \
 && wget https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
 && tar -C /usr/local/bin -xzvf dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
 && rm dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz
RUN pip install matplotlib
ADD ./* /home/
WORKDIR /home
CMD ["python", "analyzer_patch.py"]
