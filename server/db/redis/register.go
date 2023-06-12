package redis

import (
	"context"
	"fmt"
	"time"
)

func (DB *DB) GetRegisterCode(phoneNumber string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), DB.timeout)
	defer cancel()
	val, err := DB.Get(ctx, "register_code_"+phoneNumber).Result()
	if err != nil {
		return "", err
	}
	return fmt.Sprint(val), nil
}

func (DB *DB) SetRegisterCode(phoneNumber string, code string) error {
	ctx, cancel := context.WithTimeout(context.Background(), DB.timeout)
	defer cancel()
	return DB.Set(ctx, "register_code_"+phoneNumber, code, time.Minute).Err()
}
