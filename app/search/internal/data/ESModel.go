package data

import (
	"context"
	"strings"

	"github.com/olivere/elastic/v7"
)

type ESModel struct {
	client *elastic.Client
	table  string
}

func NewESModel(client *elastic.Client, table string) *ESModel {
	return &ESModel{
		client: client,
		table:  table,
	}
}

func (esm *ESModel) MustQuery(ctx context.Context, field string, filter []elastic.Query, sort string, page int, limit int) (*elastic.SearchResult, error) {
	// 分页数据处理
	isAsc := true
	if sort != "" {
		sortSlice := strings.Split(sort, " ")
		sort = sortSlice[0]
		if sortSlice[1] == "desc" {
			isAsc = false
		}
	}
	// 查询位置处理
	if page <= 1 {
		page = 1
	}

	fsc := elastic.NewFetchSourceContext(true)
	// 返回字段处理
	if field != "" {
		fieldSlice := strings.Split(field, ",")
		if len(fieldSlice) > 0 {
			for _, v := range fieldSlice {
				fsc.Include(v)
			}
		}
	}

	// 开始查询位置
	fromStart := (page - 1) * limit
	query := elastic.NewBoolQuery()
	for _, q := range filter {
		query = query.Must(q)
	}
	res, err := esm.client.Search().
		Index(esm.table).
		FetchSourceContext(fsc).
		Query(query).
		Sort(sort, isAsc).
		From(fromStart).
		Size(limit).
		Pretty(true).
		Do(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}
