package user

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Llambi/cortito/internal/core/domain"
)

func (repo MemoryRepository) Insert(user domain.CortitoUser) (domain.CortitoUser, error) {
	value := fmt.Sprintf("%d#:#%s", user.Quota, user.RefreshAt.Format(time.RFC3339))
	repo.Store.Set(user.Ip, value)
	return user, nil
}

func (repo MemoryRepository) Update(user domain.CortitoUser) (domain.CortitoUser, error) {
	repo.Store.Set(user.Ip, fmt.Sprintf("%d#:#%s", user.Quota, user.RefreshAt.Format(time.RFC3339)))
	return user, nil
}

func (repo MemoryRepository) GetByKey(key string) (domain.CortitoUser, error) {
	if quotaAndRefreshAt, exist := repo.Store.Get(key); exist {
		quotaAndRefreshAtSplit := strings.Split(quotaAndRefreshAt, "#:#")
		quotaStr, refreshAtStr := quotaAndRefreshAtSplit[0], quotaAndRefreshAtSplit[1]
		quota, err := strconv.ParseUint(quotaStr, 10, 64)
		if err != nil {
			return domain.CortitoUser{}, errors.New("can not parse quota")
		}
		RefreshAt, err := time.Parse(time.RFC3339, refreshAtStr)
		if err != nil {
			return domain.CortitoUser{}, errors.New("can not parse RefreshAt")
		}

		return domain.CortitoUser{Ip: key, Quota: quota, RefreshAt: RefreshAt}, nil
	}
	return domain.CortitoUser{}, errors.New("user key not exists")
}
