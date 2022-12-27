package main

import (
	"encoding/json"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth/credentials"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/cms"
	"time"
)

/*
get rds instance in current aliyun account

func GetInstance() []string {
	var instances = []string{"xxxx", "xxxxx"}
	return instances
}
*/

func (m *Metrics) MetricData(MetricsName []string) {
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

//func (d *AllData) Describe(ch <-chan *prometheus.Desc) {
//	for _, value := range d.DataPoint {
//		value["MetircName"] :
//	}
//}
