package jsonTest

import (
	"encoding/json"
	"fmt"
)

type VersionRangeConfig struct {
	// 起始版本操作 等于 大于等于 大于 不限  不选
	StartOption string `json:"s_op"`
	StartVersion  string `json:"s_val"`
	// 结束版本操作 小于等于 小于 不限 不选
	EndOption   string `json:"e_op"`
}

var input = `[
    {
        "s_op":"eq",
        "s_val":"8.5.90",
        "e_op":"",
        "e_val":""
    },
    {
        "s_op":"gt",
        "s_val":"1.1.1",
        "e_op":"nil",
        "e_val":""
    },
    {
        "s_op":"nil",
        "s_val":"",
        "e_val":"111",
        "e_op":"lt"
    },
    {
        "s_op":"nil",
        "s_val":"",
        "e_val":"nil",
        "e_op":""
    }
]`

type Person struct {
	Name string `json:"name"`
}

var temp = `{
	"name": "wfc",
	"age": 10
}`

type RegionInfo struct {
	CityCode     string `json:"city_code"`
	ProvinceCode string `json:"province_code"`
	CountryCode  string `json:"country_code"`
}

// Response ip库响应
type Response struct {
	Country      string `json:"country"`
	Province     string `json:"province"`
	CityName     string `json:"city"`
	AreaCode     string `json:"area_code"`
	Standard     RegionInfo `json:"standard"` // 国际标准码
	Private		 RegionInfo `json:"private"` // 美图五位码
	Latitude     string `json:"latitude"`
	Longitude    string `json:"longitude"`
}

var input1 = `{"country":"中国","province":"台湾","city":"花莲县","private":{"country_code":"10184","province_code":"10248","city_code":"10779"},"latitude":"23.987159","longitude":"121.601571"}`

var  input2 = `{
        "1":{
            "meitu":{
                "location_type":1,
                "coordinate":"待定",
                "word":"跳过",
                "is_read_second":1
            }
        }
}`
func Parse() {
	//temp := make(map[string]string)
	//temp["10086"] = input
	//
	//var result = make(map[string][]VersionRangeConfig)
	//for k, v := range temp {
	//	var tmp []VersionRangeConfig
	//	if err1 := json.Unmarshal([]byte(v), &tmp); err1 == nil {
	//		result[k] = tmp
	//	} else {
	//		fmt.Println("err: ", err1.Error())
	//	}
	//}
	//
	////fmt.Println(result)
	//var value Response
	//err := json.Unmarshal([]byte(input1), &value)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("-======" + value.Standard.CountryCode)
	//fmt.Println(value)

	var SkipConfig SkipAdvertConfig
	//var temp = make(map[string]map[string]*SkipInfo)

	err := json.Unmarshal([]byte(input2), &SkipConfig)
	if err != nil {
		panic(err)
	}

	fmt.Println(SkipConfig["1"]["meitu"].Coordinate)
}


type SkipAdvertConfig map[string]map[string]*SkipInfo

type SkipInfo struct {
	LocationType int32 `json:"location_type"`
	Coordinate  string `json:"coordinate"`
	Word string `json:"word"`
	IsReadSecond int32 `json:"is_read_second"`
}
