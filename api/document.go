package api

import (
	"github.com/meilisearch/meilisearch-go"
	"meili-api/common"
)

func GetAllDocument(uid string) (resp []interface{}, err error) {
	var documents []interface{}
	err = common.Client.Index(uid).GetDocuments(&meilisearch.DocumentsRequest{
		Limit: 10,
	}, &documents)
	if err != nil {
		return nil, err
	}
	return documents, err
}

func GetDocument(index, id string) (resp interface{}, err error) {
	var document interface{}
	err = common.Client.Index(index).GetDocument(id, &document)
	if err != nil {
		return nil, err
	}
	return document, err
}

func CreateOrUpdateDocument(uid, primaryKey string, m []map[string]interface{}) (resp *meilisearch.Task, err error) {
	task, err := common.Client.Index(uid).AddDocuments(m)
	if err != nil {
		return nil, err
	}
	return task, err
}

func DeleteAllDocument(uid string) (resp *meilisearch.Task, err error) {
	_, err = GetIndex(uid)
	if err != nil {
		return nil, err
	}
	task, err := common.Client.Index(uid).DeleteAllDocuments()
	if err != nil {
		return nil, err
	}
	return task, err
}

func DeleteDocument(uid string, m []string) (resp *meilisearch.Task, err error) {
	_, err = GetIndex(uid)
	if err != nil {
		return nil, err
	}
	task, err := common.Client.Index(uid).DeleteDocuments(m)
	if err != nil {
		return nil, err
	}
	return task, err
}
