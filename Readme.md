Импортировать пакет в проект
RedisPull "github.com/leit040/RedisPull"

добавить в .env
PATH_TO_CONFIG="Config/redis.json"

пример json:

{
"domain.com": {
"host": "redis:6339",
"password": "password",
"db": 0
},
"domain.net":{
"host": "redis:6339",
"password": "password",
"db": 0
},
"domain.org":{
"host": "redis:6339",
"password": "password",
"db": 0
}

}

после чего можно создать пул
conn, e := RedisPull.NewConnections()

получить connection для конкретного домена

actConn, err := conn.GetConnection(domain)
if !err {
log.Fatal("Error ")
}


теперь 
actConn.Ruedis - редис клиент для ruedis
value, err := actConn.Ruedis.Do(ctx, actConn.Ruedis.B().Get().Key(key).Build()).AsBytes()


actConn.GoRedis - редис клиент go-redis
err := actConn.GoRedis.Set(ctx, "key1", "value1", 0).Err()









