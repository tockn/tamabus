import MySQLdb
import schedule
from datetime import datetime
import time
import cv2
import keras
from keras.applications.imagenet_utils import preprocess_input
from keras.backend.tensorflow_backend import set_session
from keras.models import Model
from keras.preprocessing import image
import numpy as np
import tensorflow as tf
from imageio import imread

from ssd import SSD300
from ssd_utils import BBoxUtility

def getImages (cur):
    sql = "select id from buses"
    cur.execute(sql)
    ids = []
    for row in cur.fetchall():
        ids.append(row[0])

    imageData = []
    for id in ids:
        sql = "select path, bus_id from images where bus_id = %s order by created_at desc limit 1"
        cur.execute(sql, (id,))
        res = cur.fetchall()
        if len(res) != 1: continue
        imageData.append({'fileName': res[0][0], 'busId': res[0][1]})
    return imageData

def predict(cur, model, imageData):
    inputs = []
    images = []

    for data in imageData:
        img_path = '/home/images/' + data['fileName']
        img = image.load_img(img_path, target_size=(700, 700))
        img = image.img_to_array(img)
        images.append(imread(img_path))
        inputs.append(img.copy())
    inputs = preprocess_input(np.array(inputs))

    preds = model.predict(inputs, batch_size=1, verbose=0)

    results = bbox_util.detection_out(preds)

    for n, img in enumerate(images):
        # Parse the outputs.
        det_label = results[n][:, 0]
        det_conf = results[n][:, 1]
        det_xmin = results[n][:, 2]
        det_ymin = results[n][:, 3]
        det_xmax = results[n][:, 4]
        det_ymax = results[n][:, 5]

        # Get detections with confidence higher than 0.6.
        top_indices = [n for i, conf in enumerate(det_conf) if conf >= 0.6]

        top_conf = det_conf[top_indices]
        top_label_indices = det_label[top_indices].tolist()
        top_xmin = det_xmin[top_indices]
        top_ymin = det_ymin[top_indices]
        top_xmax = det_xmax[top_indices]
        top_ymax = det_ymax[top_indices]

        count_person = 0
        for i in range(top_conf.shape[0]):
            xmin = int(round(top_xmin[i] * img.shape[1]))
            ymin = int(round(top_ymin[i] * img.shape[0]))
            xmax = int(round(top_xmax[i] * img.shape[1]))
            ymax = int(round(top_ymax[i] * img.shape[0]))
            score = top_conf[i]
            label = int(top_label_indices[i])
            label_name = voc_classes[label - 1]
            display_txt = '{:0.2f}, {}'.format(score, label_name)
            display_txt = ''
            if label_name == 'Person' and score > 0.6:
                count_person += 1
        imageData[n]['congestion'] = count_person

    print("{}: {}", datetime.now().strftime("%Y/%m/%d %H:%M:%S"), imageData)
    for data in imageData:
        sql = "update congestion_log set congestion = %s, complete = 1 where bus_id = %s and complete = 0"
        cur.execute(sql, (data['congestion'], data['busId']))

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

conn = MySQLdb.connect(
    user='root',
    passwd='password',
    host='db',
    db='tamabus',
    autocommit=True
)

def job():
    cur = conn.cursor()
    imageData = getImages(cur)
    predict(cur, model, imageData)

schedule.every(9).seconds.do(job)

while True:
    schedule.run_pending()
    time.sleep(1)
