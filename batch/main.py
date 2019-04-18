# -*- coding: utf-8 -*-

from flask import Flask, request, jsonify
import imageAnalyzer as ia
import json
import re
import io
from PIL import Image
import asyncio

app = Flask(__name__)

m = 100 # 横軸分割数
n = 100 # 縦軸分割数

default_image = Image.open('./batch/default_image.jpg')
default_rgb = ia.calculate_mn_rgb(default_image, m, n)


async def worker(queue):
    while True:
        data = await queue.get()
        base64 = data['path']
        id = data['id']
        score = analyze(base64)
        print('score={}, id={}', score, id)
        queue.task_done()


def analyze(base64):
    # image_data = re.sub('^data:image/.+;base64,', '', base64).decode('base64')
    # image = Image.open(cStringIO.StringIO(image_data))

    image = Image.open(base64)

    rgb = ia.calculate_mn_rgb(image, m, n)
    score= ia.calculate_score(rgb, default_rgb, m, n)
    worstscore = ia.calculate_worst_score(m, n)
    score = 100 - score / worstscore*100
    return score


# 非同期処理のQueue
queue = asyncio.Queue()


@app.route('/analyze')
def post():
    # data = json.loads(request.data)
    queue.put_nowait({'path': './batch/image.jpg', 'id': 1})


if __name__ == '__main__':
    tasks = []
    for i in range(3):
        task = asyncio.create_task(worker(queue))
        tasks.append(task)
    app.run(debug=True)
