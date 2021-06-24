package awsiot

import (
	"teckin_ssmanager/config"

	//"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iotdataplane"
)

type MqttClient struct {
	config *config.Config
}

func New(config *config.Config) *MqttClient {
	return &MqttClient{config}
}
func (mqttCli *MqttClient) Publish(topic string, message []byte) (err error) {

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(mqttCli.config.Awsiot.Region),
		Endpoint:    aws.String(mqttCli.config.Awsiot.Endpoint),
		Credentials: credentials.NewStaticCredentials(mqttCli.config.Awsiot.AkId, mqttCli.config.Awsiot.SecretKey, ""),
	})
	mySession := session.Must(sess, err)
	//mjson, _ := json.Marshal(message)
	// Create a IoT client with additional configuration
	svc := iotdataplane.New(mySession)
	params := &iotdataplane.PublishInput{}
	params.SetPayload(message)
	params.SetQos(int64(1))
	params.SetTopic(topic)
	req, _ := svc.PublishRequest(params)
	err = req.Send()
	return
}
