package redis_repo

import (
	"FindMyDosen/database"
	"FindMyDosen/repository/auth_repo"
	"context"
	"math/rand"
	"time"
)

type RefreshTokenEntity struct {
	UID          uint   `json:"uid"`
	RefreshToken string `json:"refreshToken"`
}

func GetRefreshToken(uid uint) (string, error) {
	c := context.Background()
	redis := database.GetRedisClient()
	cache := redis.Get(c, string(uid))
	return cache.Result()
}

func NewRefreshToken(uid uint) (string, error) {
	refresh, stored, err := generateRefreshKey()
	if err != nil {
		return "", err
	}
	entity := RefreshTokenEntity{
		UID:          uid,
		RefreshToken: stored,
	}
	err = updateRefreshToken(entity)
	if err != nil {
		return "", err
	}
	return refresh, nil
}

func updateRefreshToken(refreshEntity RefreshTokenEntity) error {
	c := context.Background()
	redis := database.GetRedisClient()
	return redis.Set(c, string(refreshEntity.UID), refreshEntity.RefreshToken, time.Hour*24*365).Err()
}

func generateRefreshKey() (string, string, error) {
	rand.Seed(time.Now().UnixNano()) // seed the random number generator
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()"
	length := rand.Intn(15-8) + 8 // pick a random length between 8 and 15
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	randomKey := string(b)
	hashed, err := auth_repo.HashPassword(randomKey)
	return randomKey, hashed, err
}

func DeleteRefreshToken(uid uint) {
	c := context.Background()
	redis := database.GetRedisClient()
	err := redis.Set(c, string(uid), "nil", time.Minute*0).Err()
	println("result :", err.Error())
}
