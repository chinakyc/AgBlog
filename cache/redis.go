package cache

import (
    "time"
    "strconv"

    redigo "github.com/garyburd/redigo/redis"
)


type cache struct {
    Perfix string
    Pool   *redigo.Pool
    Expire time.Duration
}

func (self *cache) Get(key string) ([]byte, error) {
    data, err := self.getData(key)
    if err != nil {
        return nil, err
    }

    if data == nil {
        return nil, nil
    }

    return data, nil
}

func (self *cache) getData(key string) ([]byte, error) {
    conn := self.Pool.Get()
    defer conn.Close()

    key = self.Perfix + key

    data, err := conn.Do("GET", key)
    if err != nil {
        if err == redigo.ErrNil {
            return nil, nil
        }
        return nil, err
    }

    if data == nil {
        return nil, nil
    }

    content := data.([]byte)

    return content, nil
}

func (self *cache) Set(key string, content []byte) error {
    return self.setData(key, content)
}

func (self *cache) setData(key string, content []byte) error {
    key = self.Perfix + key

    params := []interface{}{key, content}

    conn := self.Pool.Get()
    defer conn.Close()
    if self.Expire != 0 {
        params = append(params, "EX", strconv.Itoa(int(self.Expire.Seconds())))
    }
    _, err := conn.Do("SET", params...)

    return err
}

func (self *cache) Del(key string) error {
    return self.delData(key)
}

func (self *cache) delData(key string) error {
    key = self.Perfix + key

    conn := self.Pool.Get()
    defer conn.Close()

    _, err := conn.Do("DEL", key)

    return err
}

func NewCache(perfix string, pool *redigo.Pool, ex time.Duration) *cache{
    return &cache{
        perfix,
        pool,
        ex,
    }
}

func NewCachePool(address string) *redigo.Pool{
    return &redigo.Pool{
        Dial: func() (redigo.Conn, error) {
            return redigo.Dial("tcp", address)
        },
        MaxIdle: 50,
    }
}
