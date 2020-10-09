package spider

import (
	"encoding/json"
	"fmt"
	"gin-vue-admin/utils"
	"log"
	"strings"

	"github.com/levigross/grequests"
)

func GetJdSku(skuIDList string) ([]*SkuData, error) {
	paramJson := fmt.Sprintf("{\"skuIds\":\"%s\"}", skuIDList)
	timestamp := utils.GetCurrentDateTime()
	strToSign := fmt.Sprintf("%sapp_key%sformatjsonmethod%sparam_json%ssign_methodmd5timestamp%sv1.0%s", appsecret, appkey, skuInfoApi, paramJson, timestamp, appsecret)
	//fmt.Println(strToSign)

	sign := strings.ToUpper(utils.MD5V([]byte(strToSign)))
	//fmt.Println(sign)

	ro := &grequests.RequestOptions{
		Headers: map[string]string{
			"User-Agent": "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/67.0.3396.87 Safari/537.36",
		},
		Params: map[string]string{
			"method":       skuInfoApi,
			"access_token": "",
			"app_key":      appkey,
			"sign_method":  "md5",
			"format":       "json",
			"timestamp":    timestamp,
			"sign":         sign,
			"param_json":   paramJson,
		},
	}

	resp, err := grequests.Get("https://router.jd.com/api?v=1.0", ro)
	// You can modify the request by passing an optional RequestOptions struct

	if err != nil {
		log.Fatalln("Unable to make request: ", err)
		return nil, err
	}

	//fmt.Println(resp.String())

	if resp.StatusCode != 200 {
		fmt.Println(fmt.Sprintf("Status Code:%d Error:%v", resp.StatusCode, resp.Error))
		return nil, err
	}

	respJson := new(SkuResp)
	err = resp.JSON(respJson)
	if err != nil {
		return nil, err
	}
	//fmt.Println(respJson)

	result := new(Result)
	err = json.Unmarshal([]byte(respJson.Response.Result), result)
	if err != nil {
		return nil, err
	}

	return result.Data, nil
}
