import glob
import json
import urllib.request
from datetime import datetime
import time
import schedule
import cv2
import keras
from keras.applications.imagenet_utils import preprocess_input
from keras.backend.tensorflow_backend import set_session
from keras.models import Model
from keras.preprocessing import image
import numpy as np
import tensorflow as tf
from imageio import imread
import matplotlib.pyplot as plt

from ssd import SSD300
from ssd_utils import BBoxUtility

url = "http://3.113.5.185/api/bus"

def getImages ():
    fileNames = glob.glob("./images/raw/*.png")
    print("==============================")
    print("==============================")
    print("==============================")
    print("==============================")
    print(fileNames)
    print("==============================")
    print("==============================")
    print("==============================")
    imageData = []
    for name in fileNames:
        id = name.split('.')[0]
        imageData.append({'fileName': name, 'busId': id})
    return imageData


def predict(model, imageData):
    inputs = []
    images = []

    for data in imageData:
        img_path = data['fileName']
        img = image.load_img(img_path, target_size=(700, 700))
        img = image.img_to_array(img)
        images.append(imread(img_path))
        inputs.append(img.copy())
    inputs = preprocess_input(np.array(inputs))

    preds = model.predict(inputs, batch_size=1, verbose=0)

    results = bbox_util.detection_out(preds)

    for i, img in enumerate(images):
        # Parse the outputs.
        det_label = results[i][:, 0]
        det_conf = results[i][:, 1]
        det_xmin = results[i][:, 2]
        det_ymin = results[i][:, 3]
        det_xmax = results[i][:, 4]
        det_ymax = results[i][:, 5]

        # Get detections with confidence higher than 0.6.
        top_indices = [i for i, conf in enumerate(det_conf) if conf >= 0.6]

        top_conf = det_conf[top_indices]
        top_label_indices = det_label[top_indices].tolist()
        top_xmin = det_xmin[top_indices]
        top_ymin = det_ymin[top_indices]
        top_xmax = det_xmax[top_indices]
        top_ymax = det_ymax[top_indices]

        colors = plt.cm.hsv(np.linspace(0, 1, 21)).tolist()

        plt.imshow(img / 255.)
        currentAxis = plt.gca()

        count_person = 0
        for n in range(top_conf.shape[0]):
            xmin = int(round(top_xmin[n] * img.shape[1]))
            ymin = int(round(top_ymin[n] * img.shape[0]))
            xmax = int(round(top_xmax[n] * img.shape[1]))
            ymax = int(round(top_ymax[n] * img.shape[0]))
            score = top_conf[n]
            label = int(top_label_indices[n])
            label_name = voc_classes[label - 1]
            display_txt = '{:0.2f}, {}'.format(score, label_name)
            display_txt = ''
            if label_name == 'Person':
                count_person += 1
            coords = (xmin, ymin), xmax-xmin+1, ymax-ymin+1
            color = colors[label]
            currentAxis.add_patch(plt.Rectangle(*coords, fill=False, edgecolor=color, linewidth=2))
            currentAxis.text(xmin, ymin, display_txt, bbox={'facecolor':color, 'alpha':0.5})
        plt.savefig("./images/result/result.png")
        imageData[i]['congestion'] = count_person

    print("{}: {}", datetime.now().strftime("%Y/%m/%d %H:%M:%S"), imageData)
    for data in imageData:
        # sql = "update congestion_log set congestion = %s, complete = 1 where bus_id = %s and complete = 0"
        # cur.execute(sql, (data['congestion'], data['busId']))
        print(data)
        req = urllib.request.Request(url, json.dumps(data).encode(), method='PUT')
        urllib.request.urlopen(req)

np.set_printoptions(suppress=True)

config = tf.ConfigProto()
config.gpu_options.per_process_gpu_memory_fraction = 0.45
set_session(tf.Session(config=config))

voc_classes = ['Aeroplane', 'Bicycle', 'Bird', 'Boat', 'Bottle',
               'Bus', 'Car', 'Cat', 'Chair', 'Cow', 'Diningtable',
               'Dog', 'Horse','Motorbike', 'Person', 'Pottedplant',
               'Sheep', 'Sofa', 'Train', 'Tvmonitor']
NUM_CLASSES = len(voc_classes) + 1

input_shape=(700, 700, 3)
model = SSD300(input_shape, num_classes=NUM_CLASSES)

model.load_weights('weights_SSD300.hdf5', by_name=True)
bbox_util = BBoxUtility(NUM_CLASSES)

def job():
    imageData = getImages()
    print(imageData)
    predict(model, imageData)

while True:
    job()
    time.sleep(5)
