### 配置
- 监听端口: 27999

### 前言
- 所有的查询都需要带时间查询
- 例：查询几月x号到几月y号的某人的弹幕
- 所有的查询都要带offset，也就页数
- 由于是保留的7天的弹幕，所以不用加密了

### queryBarrage_id
- /douyu/db/barrage/queryBarrage_id
- 根据用户id查询弹幕
```
参数:
rid: 直播间号（若为空则查全部）
id: 用户id
minTime: 下限时间
maxTime: 上限时间
page: 偏移
```

### queryBarrage_uid
- /douyu/db/barrage/queryBarrage_uid
- 根据用户uid查询弹幕
```
参数:
rid: 直播间号（若为空则查全部）
uid: 用户id
minTime: 下限时间
maxTime: 上限时间
page: 偏移
```

### queryBarrageNum_time
- /douyu/db/barrage/queryBarrageNum_time
- 返回时间段弹幕数量
```
参数：
rid: 直播间号（若为空则查全部）
minTime: 下限时间
maxTime: 上限时间
```