package main

import (
	"encoding/json"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth/credentials"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/cms"
	"time"
)

func (m *Metrics) MetricData(MetricsName []string) {
	// 每次刷新数据前将m.datapoint数据清空，否则新数据会被append到m.datapoint切片后
	m.DataPoint = []map[string]interface{}{}
	endTime := time.Now().UTC().Format(time.RFC3339)
	startTime := time.Now().UTC().Add(-5 * time.Minute).Format(time.RFC3339)
	config := sdk.NewConfig()
	credential := credentials.NewAccessKeyCredential(AccessKey, SecretKey)
	client, err := cms.NewClientWithOptions("cn-hangzhou", config, credential)
	if err != nil {
		panic(err)
	}
	request := cms.CreateDescribeMetricListRequest()
	request.Scheme = "https"
	request.Namespace = "acs_rds_dashboard"
	request.Period = "60"
	request.Dimensions = "{\"instanceId\":\"rm-uf64if27mi7e9p63l\"}"
	request.Length = "100"
	request.StartTime = startTime
	request.EndTime = endTime
	for _, metric := range MetricsName {
		var datapoint = []map[string]interface{}{}
		request.MetricName = metric
		response, err := client.DescribeMetricList(request)
		if err != nil {
			fmt.Print(err.Error())
		}
		err = json.Unmarshal([]byte(response.Datapoints), &datapoint)
		if err != nil {
			fmt.Println(err)
		}
		datapoint[2]["MetircName"] = metric
		m.DataPoint = append(m.DataPoint, datapoint[2])
	}

}
