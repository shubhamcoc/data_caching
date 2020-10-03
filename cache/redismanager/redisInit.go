package dbmanager

import (
	"fmt"
	"os/exec"
	"errors"
	"time"
	"github.com/go-redis/redis"
)

type data struct {
        key   string
        value string
}


func Init() error {
	cmd := exec.Command("/usr/local/go/src/cache/script.sh")
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Failed to start redis Server, Error: %s\n", err)
		return err
	}
	return nil
}

var client *redis.Client

// RedisConnect : This Struct used to have default variables used for redis. Also to comprise methods of redis to it's scope
type RedisConnect struct {}

// NewRedisConnect : This is a constructor function to connect redis database
func NewRedisConnect() (*RedisConnect, error) {

	client = redis.NewClient(&redis.Options{
		Addr:     ":6379",
		Password: "",
		DB:       0,
	})
	_, err := client.Ping().Result()
	return &RedisConnect{}, err
}

// Read : This helps to read the data from Redis, It Accepts keyname as input
func (pRedisConnect *RedisConnect) Read(keyname string) (string, error) {
	data, err := client.Get(keyname).Result()
	if err == redis.Nil {
		return "", errors.New("Key Not Found")
	} else if err != nil {
		return "", err
	}
	return data, err
}

// Remove : This helps to remove the data from Redis, It Accepts keyname as input
func (pRedisConnect *RedisConnect) Remove(keyname string) error {
	_, err := client.Del(keyname).Result()
	if err != nil {
		return err
	}
	return nil
}

// Store : This helps to store the data in redis, It Accepts value as input
func (pRedisConnect *RedisConnect) Store(key string, value []byte) (string, error) {

	err := client.Set(key, value, 10*time.Minute).Err()

	if err != nil {
		fmt.Println("error in saving data:", err)
		return "", err
	} else {
		return key, nil
        }
}

// Rpush : This helps to add the data in tail of redis list, It Accepts value as input
func (pRedisConnect *RedisConnect) Rpush(key string, value string)  error {

	err := client.RPush(key, value).Err()

	if err != nil {
		fmt.Println("error in performing rpush data:", err)
		return err
	}
	return nil
}

// Lpush : This helps to add the data in head of redis list, It Accepts value as input
func (pRedisConnect *RedisConnect) Lpush(key string, value string)  error {

	err := client.LPush(key, value).Err()

	if err != nil {
		fmt.Println("error in performing rpush data:", err)
		return err
	}
	return nil
}

// Lrange : This helps to list the data from Redis, It Accepts keyname as input, start and stop offset
func (pRedisConnect *RedisConnect) Lrange(keyname string, start, stop int64) ([]string, error) {
	data, err := client.LRange(keyname, start, stop).Result()
	if err == redis.Nil {
		return nil, errors.New("Key Not Found")
	} else if err != nil {
		return nil, err
	}
	return data, err
}

