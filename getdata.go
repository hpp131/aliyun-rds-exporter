package main

import (
	"encoding/json"
	"fmt"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth/credentials"
	cms "github.com/aliyun/alibaba-cloud-sdk-go/services/cms"
)

func main() {
	//type dataPoint struct {
	//	InstanceId string
	//	Average    float64
	//}
	var resultpoint = []map[string]interface{}{}
	var datapoint = []map[string]interface{}{}
	config := sdk.NewConfig()
	credential := credentials.NewAccessKeyCredential("LTAI5tPwoRHDHjZDCY85QPN8", "WBEOSMwaqljhBQlqoi5WfTFqVe45wg")
	/* use STS Token
	credential := credentials.NewStsTokenCredential("<your-access-key-id>", "<your-access-key-secret>", "<your-sts-token>")
	*/
	client, err := cms.NewClientWithOptions("cn-shanghai", config, credential)
	if err != nil {
		panic(err)
	}
	request := cms.CreateDescribeMetricListRequest()
	request.Scheme = "https"
	request.Namespace = "acs_rds_dashboard"
	request.MetricName = "ConnectionUsage"
	request.Dimensions = "{\"instanceId\":\"rm-uf64if27mi7e9p63l\"}"
	request.Period = "60"
	request.StartTime = "2022-12-26T03:13:43Z"
	request.EndTime = "2022-12-26T03:18:43Z"
	response, err := client.DescribeMetricList(request)
	if err != nil {
		fmt.Print(err.Error())
	}
	err = json.Unmarshal([]byte(response.Datapoints), &datapoint)
	if err != nil {
		fmt.Println(err)
	}
	datapoint[2]["metricname"] = "ConnectionUsage"
	//fmt.Printf("the transpoints type is %T\n", datapoint)
	fmt.Printf("the value is ", datapoint[2])
	resultpoint = append(resultpoint, datapoint[2])
}
