package database

import (
	"context"
	"github.com/olivere/elastic/v7"
	"library_pro/config"
	"log"
	"os"
)

//连接集群
func ConnectES() *elastic.Client {
	//集群连接部分参数配置
	var options []elastic.ClientOptionFunc
	es_error_log := log.New(os.Stdout, "APP", log.LstdFlags)
	options = append(options, elastic.SetURL(config.ES_CLUSTER_HOST_ONE, config.ES_CLUSTER_HOST_TWO, config.ES_CLUSTER_HOST_THREE)) //集群地址
	options = append(options, elastic.SetHealthcheck(true))                                                                         //健康检查
	options = append(options, elastic.SetHealthcheckInterval(20))                                                                   //健康检查周期20s
	options = append(options, elastic.SetErrorLog(es_error_log))                                                                    //错误日志
	client, err := elastic.NewClient(options...)
	if err != nil {
		// Handle error
		panic(err)
	}
	return client
}

//创建文档
func CreateESDoc(client *elastic.Client, body string, index string, id string) string {
	//tweet2 := `{"user" : "olivere", "message" : "It's a Raggy Waltz"}`
	result, err := client.Index().
		Index(index).
		Id(id).
		BodyJson(body).
		Do(context.Background())
	if err != nil {
		panic(err.Error())
	}
	return result.Id
}

//删除一个文档
func DeleteESDoc(client *elastic.Client, index string, id string) string {
	res, err := client.Delete().
		Index(index).
		Id(id).
		Do(context.Background())
	if err != nil {
		panic(err.Error())
	}
	return res.Id
}

//修改
func Update(client *elastic.Client, index string, id string, update_map map[string]interface{}) string {
	res, err := client.Update().
		Index(index).
		Id(id).
		Doc(update_map).
		Do(context.Background())
	if err != nil {
		panic(err.Error())
	}
	return res.Id
}

//查找
func FindOne(client *elastic.Client, index string, id string) interface{} {
	//通过id查找
	this_doc, err := client.Get().Index(index).Id(id).Do(context.Background())
	if err != nil {
		panic(err)
	}
	return this_doc
}



//搜索
//func query() {
//	var res *elastic.SearchResult
//	var err error
//	//取所有
//	res, err = client.Search("megacorp").Type("employee").Do(context.Background())
//	printEmployee(res, err)
//
//	//字段相等
//	q := elastic.NewQueryStringQuery("last_name:Smith")
//	res, err = client.Search("megacorp").Type("employee").Query(q).Do(context.Background())
//	if err != nil {
//		println(err.Error())
//	}
//	printEmployee(res, err)
//
//	if res.Hits.TotalHits > 0 {
//		fmt.Printf("Found a total of %d Employee \n", res.Hits.TotalHits)
//
//		for _, hit := range res.Hits.Hits {
//
//			var t Employee
//			err := json.Unmarshal(*hit.Source, &t) //另外一种取数据的方法
//			if err != nil {
//				fmt.Println("Deserialization failed")
//			}
//
//			fmt.Printf("Employee name %s : %s\n", t.FirstName, t.LastName)
//		}
//	} else {
//		fmt.Printf("Found no Employee \n")
//	}
//
//	//条件查询
//	//年龄大于30岁的
//	boolQ := elastic.NewBoolQuery()
//	boolQ.Must(elastic.NewMatchQuery("last_name", "smith"))
//	boolQ.Filter(elastic.NewRangeQuery("age").Gt(30))
//	res, err = client.Search("megacorp").Type("employee").Query(q).Do(context.Background())
//	printEmployee(res, err)
//
//	//短语搜索 搜索about字段中有 rock climbing
//	matchPhraseQuery := elastic.NewMatchPhraseQuery("about", "rock climbing")
//	res, err = client.Search("megacorp").Type("employee").Query(matchPhraseQuery).Do(context.Background())
//	printEmployee(res, err)
//
//	//分析 interests
//	aggs := elastic.NewTermsAggregation().Field("interests")
//	res, err = client.Search("megacorp").Type("employee").Aggregation("all_interests", aggs).Do(context.Background())
//	printEmployee(res, err)
//
//}