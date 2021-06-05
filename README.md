# gin-bubble
A go web small program

Connect to MySQL on Mac
```bash
# step1 Start mysql container
docker run --name mysql8019 -p 3306:3306 -e MYSQL_ROOT_PASSWORD=123456 -d mysql:8.0.19

# step2 Start mysql client
docker run -it --network host --rm mysql mysql -h127.0.0.1 -P3306 --default-character-set=utf8mb4 -uroot -p
```

>[Go Web Video](https://www.bilibili.com/video/BV1gJ411p7xC?from=search&seid=6720858265709214067)

>[Golang Notes](https://gitee.com/moxi159753/LearningNotes/tree/master/Golang)
