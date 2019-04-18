import MySQLdb

conn = MySQLdb.connect(
        user='root',
        passwd='password',
        host='127.0.0.1',
        db='tamabus')
cur = conn.cursor()

sql = "select * from" 
