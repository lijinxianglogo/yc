package awsiot
//
//type Config struct {
//	akId      string // aws秘钥key
//	secretKey string // aws秘钥secret
//	region    string // 服务区域
//	endpoint  string // 地址
//}
//
//func NewConfig() *Config {
//	c := &Config{
//		akId:      "AKIAWED5N4XQVXOQCAPN ", // aws秘钥key
//		secretKey: "itmUAmdeBOGAqKa0Cj7CPicnGO2AhThvhi5p/8aR", // aws秘钥secret
//		region:    "ap-northeast-1", // 服务区域
//		endpoint:  "a1on17nd3yssr4-ats.iot.ap-northeast-1.amazonaws.com", // 服务地址
//	}
//	return c
//}
//
///*
//设置服务地址,必填项
//*/
//func (c *Config) SetEndpoint(endpoint string) *Config {
//	c.endpoint = endpoint
//	return c
//}
//
///*
//设置秘钥id，必填项
//*/
//func (c *Config) SetAkId(akId string) *Config {
//	c.akId = akId
//	return c
//}
//
///*
//设置秘钥，必填项
//*/
//func (c *Config) SetSecretKey(secretKey string) *Config {
//	c.secretKey = secretKey
//	return c
//}
//
///*
//设置区域，必填项
//*/
//func (c *Config) SetRegion(region string) *Config {
//	c.region = region
//	return c
//}
