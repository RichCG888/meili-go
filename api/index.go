package api

import (
	"github.com/meilisearch/meilisearch-go"
	"meili-api/common"
)

func GetAllIndexes() (resp []*meilisearch.Index, err error) {
	indexes, err := common.Client.GetAllIndexes()
	if err != nil {
		return nil, err
	}
	return indexes, err
}

func GetIndex(index string) (resp *meilisearch.Index, err error) {
	i, err := common.Client.GetIndex(index)
	if err != nil {
		return nil, err
	}
	return i, err
}

func CreateIndex(uid, primaryKey string) (resp *meilisearch.Task, err error) {
	indexConfig := new(meilisearch.IndexConfig)
	indexConfig.Uid = uid
	if len(primaryKey) != 0 {
		indexConfig.PrimaryKey = primaryKey
	}
	i, err := common.Client.CreateIndex(indexConfig)
	if err != nil {
		return nil, err
	}
	return i, err
}

func UpdateIndex(uid, primaryKey string) (resp *meilisearch.Task, err error) {
	_, err = GetIndex(uid)
	if err != nil {
		return nil, err
	}
	task, err := common.Client.Index(uid).UpdateIndex(primaryKey)
	if err != nil {
		return nil, err
	}
	return task, err
}

func DeleteIndex(uid string) (resp *meilisearch.Task, err error) {
	_, err = GetIndex(uid)
	if err != nil {
		return nil, err
	}
	task, err := common.Client.DeleteIndex(uid)
	if err != nil {
		return nil, err
	}
	return task, err
}
