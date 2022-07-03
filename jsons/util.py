import urllib3

http = urllib3.PoolManager()  # 创建PoolManager对象生成请求, 由该实例对象处理与线程池的连接以及线程安全的所有细节
with open("productparamimgs.json", 'a+') as f:
    f.write('[')
for i in range(34):
    response = http.request('GET', 'http://cmall.congz.top/api/v1/param-imgs/' + str(i + 1))  # get方式请求
    reStr = response.data.decode('utf-8')
    reStr = reStr[reStr.find("[") + 1:reStr.rfind("]")]
    with open("productparamimgs.json", 'a+') as f:
        f.write(reStr + ',')
with open("productparamimgs.json", 'a+') as f:
    f.write(']')
