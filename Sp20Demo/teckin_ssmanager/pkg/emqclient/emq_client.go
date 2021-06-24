package emqclient

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

//相关配置信息
type MqttClient struct {
	config *EmqConfig
	client mqtt.Client
}

func NewMqttClient(config *EmqConfig) (*MqttClient, error) {
	client := &MqttClient{
		config: config,
	}
	return client, nil
}

func (mqttCli *MqttClient) checkConnect() (bool, error) {
	if mqttCli.client == nil || !mqttCli.client.IsConnected() || !mqttCli.client.IsConnectionOpen() {
		client, err := mqttCli.connect()
		if err != nil {
			return false, err
		}
		mqttCli.client = client
	} else {
		return true, nil
	}
	return true, nil
}

func NewTLSConfig() *tls.Config {
	// Import trusted certificates from CAfile.pem.
	// Alternatively, manually add CA certificates to
	// default openssl CA bundle.
	certpool := x509.NewCertPool()
	gen, _ := os.Getwd()
	pemCerts, err := ioutil.ReadFile(gen + string(os.PathSeparator) + "config" + string(os.PathSeparator) + "root-CA.crt")
	if err == nil {
		certpool.AppendCertsFromPEM(pemCerts)
	}
	//fmt.Println("0. resd pemCerts Success")

	// Import client certificate/key pair
	cert, err := tls.LoadX509KeyPair(gen+string(os.PathSeparator)+"config"+string(os.PathSeparator)+"tokyo-certificate.pem.crt", gen+string(os.PathSeparator)+"config"+string(os.PathSeparator)+"tokyo-private.pem.key")
	if err != nil {
		panic(err)
	}
	//fmt.Println("1. resd cert Success")

	// Just to print out the client certificate..
	cert.Leaf, err = x509.ParseCertificate(cert.Certificate[0])
	if err != nil {
		panic(err)
	}
	//fmt.Println("2. resd cert.Leaf Success")

	// Create tls.Config with desired tls properties
	return &tls.Config{
		// RootCAs = certs used to verify server cert.
		RootCAs: certpool,
		// ClientAuth = whether to request cert from server.
		// Since the server is set up for SSL, this happens
		// anyways.
		ClientAuth: tls.NoClientCert,
		// ClientCAs = certs used to validate client cert.
		ClientCAs: nil,
		// InsecureSkipVerify = verify that cert contents
		// match server. IP matches what is in cert etc.
		InsecureSkipVerify: true,
		// Certificates = list of certs client sends to server.
		Certificates: []tls.Certificate{cert},
	}
}

func (mqttCli *MqttClient) connect() (mqtt.Client, error) {
	tlsconfig := NewTLSConfig()
	opts := mqtt.NewClientOptions().AddBroker(mqttCli.config.url).SetClientID(mqttCli.config.clientId)
	//opts.SetProtocolVersion(3)
	opts.SetUsername(mqttCli.config.userName)
	pass := mqttCli.config.password
	opts.SetPassword(pass)
	opts.SetKeepAlive(30 * time.Second)
	opts.SetPingTimeout(10 * time.Second)
	opts.SetCleanSession(false)
	opts.SetTLSConfig(tlsconfig)
	c := mqtt.NewClient(opts)
	token := c.Connect()
	if token.Wait() && token.Error() != nil {
		fmt.Println("connect aws-iot fail")
		return nil, token.Error()
	}
	fmt.Println("connect aws-iot success")
	return c, nil
}
func (mqttCli *MqttClient) Publish(topic string, message string) (status bool, err error) {
	_, err = mqttCli.checkConnect()
	if err != nil {
		return false, err
	}
	token := mqttCli.client.Publish(topic, 1, false, message)
	if token.Wait() && token.Error() != nil {
		return false, token.Error()
	}
	return true, nil
}

func (mqttCli *MqttClient) SyncPublish(from, topic string, message string) (string, error) {
	_, err := mqttCli.checkConnect()
	if err != nil {
		return "", err
	}
	subCh := make(chan string, 1)
	timeout := make(chan bool, 1)
	go func() {
		time.Sleep(mqttCli.config.waitTimeout * time.Second)
		timeout <- true
		close(timeout)
	}()
	subToken := mqttCli.client.Subscribe(from, 0, func(client mqtt.Client, msg mqtt.Message) {
		subMessage := make(map[string]map[string]string)
		originMsg := make(map[string]map[string]string)
		_ = json.Unmarshal(msg.Payload(), &subMessage)
		_ = json.Unmarshal([]byte(message), &originMsg)
		subMsgId := subMessage["header"]["messageId"]
		originMsgId := originMsg["header"]["messageId"]
		if subMsgId == originMsgId && subMessage["header"]["method"] == originMsg["header"]["method"]+"ACK" {
			client.Unsubscribe(msg.Topic())
			subCh <- string(msg.Payload())
			close(subCh)
		}
	})
	if subToken.Wait() && subToken.Error() != nil {
		return "", subToken.Error()
	}
	pubToken := mqttCli.client.Publish(topic, 1, false, message)
	if pubToken.Wait() && pubToken.Error() != nil {
		return "", pubToken.Error()
	}
	select {
	case msg := <-subCh:
		return msg, nil
	case <-timeout:
		err = errors.New("callback get ack time out")
		return "", err
	}
}

func (mqttCli *MqttClient) Subscribe(topic string, qos byte, callback mqtt.MessageHandler) (bool, error) {
	_, err := mqttCli.checkConnect()
	if err != nil {
		return false, err
	}
	token := mqttCli.client.Subscribe(topic, qos, callback)
	if token.Wait() && token.Error() != nil {
		return false, token.Error()
	}
	return true, nil
}

func (mqttCli *MqttClient) UnSubscribe(topic []string) (bool, error) {
	_, err := mqttCli.checkConnect()
	if err != nil {
		return false, err
	}
	token := mqttCli.client.Unsubscribe(topic...)
	if token.Wait() && token.Error() != nil {
		return false, token.Error()
	}
	return true, nil
}

func (mqttCli *MqttClient) Close(t uint) {
	connectStatus := mqttCli.client.IsConnectionOpen()
	if !connectStatus {
		return
	}
	mqttCli.client.Disconnect(t)
}
