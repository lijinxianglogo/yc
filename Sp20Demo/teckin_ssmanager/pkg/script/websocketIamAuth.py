'''
aws api gateway添加iam信息校验，get方法
'''
import websocket
import sys, datetime, hashlib, hmac

access_key = "AKIAWED5N4XQSN2BHBXT"         # aws 访问密钥
secret_key = "wCQ0MQLv+dAK5hhbUU8nOLDAqCTH/LQDt+NUhYhj"         # aws 密钥
wss_url = "wss://l3u1zeammb.execute-api.ap-northeast-1.amazonaws.com/production"    # apigateway生成的wss的url
wss_host = "l3u1zeammb.execute-api.ap-northeast-1.amazonaws.com"                    # apigateway生成的wss的host
region = "ap-northeast-1"            # apigateway所在region
canonical_uri = "/production"   # 请求参数的路径


def on_message(ws, message):
    print(ws)
    print(message)


def on_error(ws, error):
    print(ws)
    print(error)


def on_close(ws):
    print(ws)
    print("### closed ###")


websocket.enableTrace(True)


def sign(key, msg):
    return hmac.new(key, msg.encode('utf-8'), hashlib.sha256).digest()


def getSignatureKey(key, dateStamp, regionName, serviceName):
    kDate = sign(('AWS4' + key).encode('utf-8'), dateStamp)
    kRegion = sign(kDate, regionName)
    kService = sign(kRegion, serviceName)
    kSigning = sign(kService, 'aws4_request')
    return kSigning


def get_websocket_header(access_key, secret_key, host, method='GET', service='execute-api', region="us-east-1", canonical_uri="/", request_parameters=""):
    if access_key is None or secret_key is None:
        print('No access key is available.')
        sys.exit()
    # Create a date for headers and the credential string
    t = datetime.datetime.utcnow()
    amzdate = t.strftime('%Y%m%dT%H%M%SZ')
    datestamp = t.strftime('%Y%m%d')
    canonical_querystring = request_parameters
    canonical_headers = 'host:' + host + '\n' + 'x-amz-date:' + amzdate + '\n'
    signed_headers = 'host;x-amz-date'
    payload_hash = hashlib.sha256(('').encode('utf-8')).hexdigest()
    canonical_request = method + '\n' + \
                        canonical_uri + '\n' + \
                        canonical_querystring + '\n' + \
                        canonical_headers + '\n' + \
                        signed_headers + '\n' + \
                        payload_hash
    algorithm = 'AWS4-HMAC-SHA256'
    credential_scope = datestamp + '/' + region + '/' + service + '/' + 'aws4_request'
    string_to_sign = algorithm + '\n' + amzdate + '\n' + credential_scope + '\n' + hashlib.sha256(
        canonical_request.encode('utf-8')).hexdigest()
    signing_key = getSignatureKey(secret_key, datestamp, region, service)
    signature = hmac.new(signing_key, (string_to_sign).encode('utf-8'), hashlib.sha256).hexdigest()
    authorization_header = algorithm + ' ' + 'Credential=' + access_key + '/' + credential_scope + ', ' + 'SignedHeaders=' + signed_headers + ', ' + 'Signature=' + signature
    return {'x-amz-date': amzdate, 'Authorization': authorization_header, "AuthInfo": "this is a key"}


ws = websocket.WebSocketApp(wss_url,
                            header=get_websocket_header(access_key=access_key,
                                                        secret_key=secret_key,
                                                        host=wss_host,
                                                        region=region,
                                                        canonical_uri=canonical_uri),
                            on_message=on_message,
                            on_error=on_error,
                            on_close=on_close)
ws.run_forever()

