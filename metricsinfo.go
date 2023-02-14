package main

import (
	"encoding/json"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth/credentials"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/cms"
	"sync"
	"time"
)

func (m *Metrics) MetricData(instanceid []string) {
	// 每次刷新数据前将m.datapoint数据清空，否则新数据会被append到m.datapoint切片后
	m.DataPoint = []map[string]interface{}{}
	var demensions []map[string]string
	for _, value := range instanceid {
		element := make(map[string]string)
		element["instanceId"] = value
		demensions = append(demensions, element)
	}
	demension, err := json.Marshal(demensions)
	if err != nil {
		fmt.Println(err)
	}
	configfile.AliYunApiRequest(string(demension), m.DataPoint)

}

func (c *ConfigFile) AliYunApiRequest(demensions string, dp []map[string]interface{}) {
	endTime := time.Now().UTC().Format(time.RFC3339)
	startTime := time.Now().UTC().Add(-1 * time.Minute).Format(time.RFC3339)
	config := sdk.NewConfig()
	credential := credentials.NewAccessKeyCredential(c.Accesskey, c.Secretkey)
	client, err := cms.NewClientWithOptions(configfile.Region, config, credential)
	if err != nil {
		panic(err)
	}
	request := cms.CreateDescribeMetricListRequest()
	request.Scheme = "https"
	request.Namespace = "acs_rds_dashboard"
	request.Period = "60"
	request.Length = "100"
	request.StartTime = startTime
	request.EndTime = endTime
	request.Dimensions = demensions
	var wg sync.WaitGroup
	var mt sync.Mutex
	for _, metric := range MetricsName {
		metric := metric
		var datapoint []map[string]interface{}
		wg.Add(1)
		go func() {
			defer wg.Done()
			request.MetricName = metric
			response, err := client.DescribeMetricList(request)
			if err != nil {
				fmt.Print(err.Error())
			}
			err = json.Unmarshal([]byte(response.Datapoints), &datapoint)
			if err != nil {
				fmt.Println(err)
			}
			//对返回的response中的map元素手动添加metrics_name字段
			for _, value := range datapoint {
				mt.Lock()
				value["metrics_name"] = metric
				mt.Unlock()
			}
			//mt.Lock()
			dp = append(dp, datapoint...)
			//mt.Unlock()
		}()

	}
	wg.Wait()
}
