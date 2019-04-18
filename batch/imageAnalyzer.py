#!C:/Software/Python/python/python.exe
# -*- coding: utf-8 -*-

# usage:
#   ./imaga.py image1 image2
#     image1: 比較元画像名
#     image2: 比較元画像名
#

from PIL import Image
import math
import argparse
import re
import os
import time

parser = argparse.ArgumentParser()
parser.add_argument("image1", help="image1 compared with image2")
parser.add_argument("image2", help="image2 compared with image1")

# 分割領域の rgb 平均の計算
def calculate_mn_rgb(im, m, n):
    #RGBに変換
    rgb_im = im.convert('RGB')
    #画像サイズを取得
    size = rgb_im.size
    # 色分布評価(x/m,y/nの剰余bitは無視される)
    m_size=math.floor(size[0]/m)
    n_size=math.floor(size[1]/n)
    bit_num =m_size * n_size
    rgb = [[[0 for i in range(3)] for j in range(n) ] for k in range(m)] # P = M * N * 3
    for x in range(m_size*m):
        for y in range(n_size*n):
            #ピクセルを取得
            r,g,b = rgb_im.getpixel((x,y))
            #平均化して(m*n*3)配列に挿入
            #print("{}, {}".format(x, y))
            rgb[math.floor(x/m_size)][math.floor(y/n_size)][0] += r / bit_num
            rgb[math.floor(x/m_size)][math.floor(y/n_size)][1] += g / bit_num
            rgb[math.floor(x/m_size)][math.floor(y/n_size)][2] += b / bit_num
    #print(rgb)
    return rgb

# rgb値の差の合計
def calculate_score(image1_rgb, image2_rgb, m, n):
    d = 0
    for x in range(m):
        for y in range(n):
            for i in range(3):
                d += (image1_rgb[x][y][i] - image2_rgb[x][y][i])**2
    return d

# m×nの反転画像のスコア
def calculate_worst_score(m, n):
    d = 0
    for x in range(m):
        for y in range(n):
            for i in range(3):
                d += 255**2
    return d

#
# # Initialize
# imageList = []
# scoreList = {}
# args = parser.parse_args() #引数解析用
# directory = os.listdir() #ディレクトリ
# M = 100 #横軸分割数
# N = 100 #縦軸分割数
# image1_path=args.image1
# image2_path=args.image2
#
# # 画像があるか判定
# if not(os.path.exists(image1_path)) or not(os.path.exists(image2_path)):
#     try:
#         raise NameError('')
#     except NameError:
#         print("The files do not exist")
#         exit()
#
# start=time.time()
# print("Processing: Calculating rgb")
# image1_rgb = calculate_mn_rgb(image1_path, M, N)
# image2_rgb = calculate_mn_rgb(image2_path, M, N)
# print("Processing: Calculating score")
# score= calculate_score(image1_rgb, image2_rgb, M, N)
# elapsed_time=time.time()-start
# worstScore = calculate_worst_score(M, N)
# print("Score: {0:8.4f}".format(100 - score / worstScore*100))
# print("Time: {}".format(elapsed_time))
