FROM yoanlin/opencv-python3:jessie
LABEL maintainer="tockn <s.taku.inthebox@gmail.com>"

RUN pip install cython \
 && pip install six \
 && pip install Pillow \
 && pip install tensorflow \
 && pip install dlib \
 && pip install face_recognition \
 && pip install pandas \
 && rm -rf ~/.cache/pip
RUN pip install keras numpy imageio
ADD ./* /
CMD ["python", "analyzer.py"]
