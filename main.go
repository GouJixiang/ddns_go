// This file is auto-generated, don't edit it. Thanks.
/**
* Author: 东南dnf
* Powered by Aliyun Darabonba Language :)
 */
package main

import (
	dns "github.com/alibabacloud-go/alidns-20150109/v2/client"
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	console "github.com/alibabacloud-go/tea-console/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
	"log"
	"os"
)

/**
* Init 初始化客户端
 * @param accessKeyId
 * @param accessKeySecret
 * @return Client
 * @throws Exception
*/
func Init(accessKeyId *string, accessKeySecret *string, regionId *string) (_result *dns.Client, _err error) {
	config := &openapi.Config{}
	// 传AccessKey ID入config
	config.AccessKeyId = accessKeyId
	// 传AccessKey Secret入config
	config.AccessKeySecret = accessKeySecret
	config.RegionId = regionId
	_result = &dns.Client{}
	_result, _err = dns.NewClient(config)
	return _result, _err
}

/**
* DescribeDomains  查询账户下域名
 * @param client      客户端
 * @throws Exception
*/
func DescribeDomains(client *dns.Client) (_err error) {
	req := &dns.DescribeDomainsRequest{}
	console.Log(tea.String("查询域名列表(json)↓"))
	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		resp, _err := client.DescribeDomains(req)
		if _err != nil {
			return _err
		}

		console.Log(util.ToJSONString(tea.ToMap(resp)))

		return nil
	}()

	if tryErr != nil {
		var error = &tea.SDKError{}
		if _t, ok := tryErr.(*tea.SDKError); ok {
			error = _t
		} else {
			error.Message = tea.String(tryErr.Error())
		}
		console.Log(error.Message)
	}
	return _err
}

/**
* AddDomain  阿里云云解析添加域名
 * @param client      客户端
 * @param domainName  域名名称
 * @throws Exception
*/
func AddDomain(client *dns.Client, domainName *string) (_err error) {
	req := &dns.AddDomainRequest{}
	req.DomainName = domainName
	console.Log(tea.String("云解析添加域名(" + tea.StringValue(domainName) + ")的结果(json)↓"))
	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		resp, _err := client.AddDomain(req)
		if _err != nil {
			return _err
		}

		console.Log(util.ToJSONString(tea.ToMap(resp)))

		return nil
	}()

	if tryErr != nil {
		var error = &tea.SDKError{}
		if _t, ok := tryErr.(*tea.SDKError); ok {
			error = _t
		} else {
			error.Message = tea.String(tryErr.Error())
		}
		console.Log(error.Message)
	}
	return _err
}

/**
* DescribeDomainRecords 查询域名解析记录
 * @param client          客户端
 * @param domainName      域名名称
 * @throws Exception
*/
func DescribeDomainRecords(client *dns.Client, domainName *string) (_err error) {
	req := &dns.DescribeDomainRecordsRequest{}
	req.DomainName = domainName
	console.Log(tea.String("查询域名(" + tea.StringValue(domainName) + ")的解析记录(json)↓"))
	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		resp, _err := client.DescribeDomainRecords(req)
		if _err != nil {
			return _err
		}

		console.Log(util.ToJSONString(tea.ToMap(resp)))

		return nil
	}()

	if tryErr != nil {
		var error = &tea.SDKError{}
		if _t, ok := tryErr.(*tea.SDKError); ok {
			error = _t
		} else {
			error.Message = tea.String(tryErr.Error())
		}
		console.Log(error.Message)
	}
	return _err
}

/**
* DescribeRecordLogs  查询域名解析记录日志
 * @param client         客户端
 * @param domainName     域名名称
 * @throws Exception
*/
func DescribeRecordLogs(client *dns.Client, domainName *string) (_err error) {
	req := &dns.DescribeRecordLogsRequest{}
	req.DomainName = domainName
	console.Log(tea.String("查询域名(" + tea.StringValue(domainName) + ")的解析记录日志(json)↓"))
	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		resp, _err := client.DescribeRecordLogs(req)
		if _err != nil {
			return _err
		}

		console.Log(util.ToJSONString(tea.ToMap(resp)))

		return nil
	}()

	if tryErr != nil {
		var error = &tea.SDKError{}
		if _t, ok := tryErr.(*tea.SDKError); ok {
			error = _t
		} else {
			error.Message = tea.String(tryErr.Error())
		}
		console.Log(error.Message)
	}
	return _err
}

/**
* DescribeDomainRecordByRecordId 查询域名解析记录信息
 * @param client         客户端
 * @param RecordId       解析记录id
 * @throws Exception
*/
func DescribeDomainRecordByRecordId(client *dns.Client, recordId *string) (_err error) {
	req := &dns.DescribeDomainRecordInfoRequest{}
	req.RecordId = recordId
	console.Log(tea.String("查询RecordId:" + tea.StringValue(recordId) + "的域名解析记录信息(json)↓"))
	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		resp, _err := client.DescribeDomainRecordInfo(req)
		if _err != nil {
			return _err
		}

		console.Log(util.ToJSONString(tea.ToMap(resp)))

		return nil
	}()

	if tryErr != nil {
		var error = &tea.SDKError{}
		if _t, ok := tryErr.(*tea.SDKError); ok {
			error = _t
		} else {
			error.Message = tea.String(tryErr.Error())
		}
		console.Log(error.Message)
	}
	return _err
}

/**
* DescribeDomainInfo  查询域名信息
 * @param client         客户端
 * @param domainName     域名名称
 * @throws Exception
*/
func DescribeDomainInfo(client *dns.Client, domainName *string) (_err error) {
	req := &dns.DescribeDomainInfoRequest{}
	req.DomainName = domainName
	console.Log(tea.String("查询域名:" + tea.StringValue(domainName) + "的信息(json)↓"))
	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		resp, _err := client.DescribeDomainInfo(req)
		if _err != nil {
			return _err
		}

		console.Log(util.ToJSONString(tea.ToMap(resp)))

		return nil
	}()

	if tryErr != nil {
		var error = &tea.SDKError{}
		if _t, ok := tryErr.(*tea.SDKError); ok {
			error = _t
		} else {
			error.Message = tea.String(tryErr.Error())
		}
		console.Log(error.Message)
	}
	return _err
}

/**
* AddDomainRecord  添加域名解析记录
 * @param client            客户端
 * @param domainName        域名名称
 * @param RR                主机记录
 * @param recordType              记录类型(A/NS/MX/TXT/CNAME/SRV/AAAA/CAA/REDIRECT_URL/FORWARD_URL)
 * @param value             记录值
 * @throws Exception
*/
func AddDomainRecord(client *dns.Client, domainName *string, RR *string, recordType *string, Value *string) (_err error) {
	req := &dns.AddDomainRecordRequest{}
	req.DomainName = domainName
	req.RR = RR
	req.Type = recordType
	req.Value = Value
	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		resp, _err := client.AddDomainRecord(req)
		if _err != nil {
			return _err
		}

		console.Log(tea.String("添加域名解析记录的结果(json)↓"))
		console.Log(util.ToJSONString(tea.ToMap(resp)))

		return nil
	}()

	if tryErr != nil {
		var error = &tea.SDKError{}
		if _t, ok := tryErr.(*tea.SDKError); ok {
			error = _t
		} else {
			error.Message = tea.String(tryErr.Error())
		}
		console.Log(error.Message)
	}
	return _err
}

/**
* UpdateDomainRecord  更新域名解析记录
 * @param client          客户端
 * @param recordId        解析记录ID
 * @param RR              主机记录
 * @param recordType            记录类型(A/NS/MX/TXT/CNAME/SRV/AAAA/CAA/REDIRECT_URL/FORWARD_URL)
 * @param value           记录值
 * @throws Exception
*/
func UpdateDomainRecord(client *dns.Client, recordId *string, RR *string, recordType *string, Value *string) (_err error) {
	req := &dns.UpdateDomainRecordRequest{}
	req.RecordId = recordId
	req.RR = RR
	req.Type = recordType
	req.Value = Value
	console.Log(tea.String("更新域名解析记录的结果(json)↓"))
	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		resp, _err := client.UpdateDomainRecord(req)
		if _err != nil {
			return _err
		}

		console.Log(util.ToJSONString(tea.ToMap(resp)))

		return nil
	}()

	if tryErr != nil {
		var error = &tea.SDKError{}
		if _t, ok := tryErr.(*tea.SDKError); ok {
			error = _t
		} else {
			error.Message = tea.String(tryErr.Error())
		}
		console.Log(error.Message)
	}
	return _err
}

/**
* SetDomainRecordStatus  设置域名解析状态
 * @param client      客户端
 * @param recordId    解析记录ID
 * @param status      解析状态(ENABLE/DISABLE)
 * @throws Exception
*/
func SetDomainRecordStatus(client *dns.Client, recordId *string, status *string) (_err error) {
	req := &dns.SetDomainRecordStatusRequest{}
	req.RecordId = recordId
	req.Status = status
	console.Log(tea.String("设置域名解析状态的结果(json)↓"))
	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		resp, _err := client.SetDomainRecordStatus(req)
		if _err != nil {
			return _err
		}

		console.Log(util.ToJSONString(tea.ToMap(resp)))

		return nil
	}()

	if tryErr != nil {
		var error = &tea.SDKError{}
		if _t, ok := tryErr.(*tea.SDKError); ok {
			error = _t
		} else {
			error.Message = tea.String(tryErr.Error())
		}
		console.Log(error.Message)
	}
	return _err
}

/**
* DeleteDomainRecord  删除域名解析记录
 * @param client         客户端
 * @param recordId       解析记录ID
 * @throws Exception
*/
func DeleteDomainRecord(client *dns.Client, recordId *string) (_err error) {
	req := &dns.DeleteDomainRecordRequest{}
	req.RecordId = recordId
	console.Log(tea.String("删除域名解析记录的结果(json)↓"))
	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		resp, _err := client.DeleteDomainRecord(req)
		if _err != nil {
			return _err
		}

		console.Log(util.ToJSONString(tea.ToMap(resp)))

		return nil
	}()

	if tryErr != nil {
		var error = &tea.SDKError{}
		if _t, ok := tryErr.(*tea.SDKError); ok {
			error = _t
		} else {
			error.Message = tea.String(tryErr.Error())
		}
		console.Log(error.Message)
	}
	return _err
}

/**
* DescribeDomainGroups  查询域名组
 * @param client 客户端
 * @throws Exception
*/
func DescribeDomainGroups(client *dns.Client) (_err error) {
	req := &dns.DescribeDomainGroupsRequest{}
	console.Log(tea.String("查询域名组(json)↓"))
	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		resp, _err := client.DescribeDomainGroups(req)
		if _err != nil {
			return _err
		}

		console.Log(util.ToJSONString(tea.ToMap(resp)))

		return nil
	}()

	if tryErr != nil {
		var error = &tea.SDKError{}
		if _t, ok := tryErr.(*tea.SDKError); ok {
			error = _t
		} else {
			error.Message = tea.String(tryErr.Error())
		}
		console.Log(error.Message)
	}
	return _err
}

/**
* AddDomainGroup  添加域名组
 * @param client      客户端
 * @param groupName   域名组名
 * @throws Exception
*/
func AddDomainGroup(client *dns.Client, groupName *string) (_err error) {
	req := &dns.AddDomainGroupRequest{}
	req.GroupName = groupName
	console.Log(tea.String("添加域名组的结果(json)↓"))
	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		resp, _err := client.AddDomainGroup(req)
		if _err != nil {
			return _err
		}

		console.Log(util.ToJSONString(tea.ToMap(resp)))

		return nil
	}()

	if tryErr != nil {
		var error = &tea.SDKError{}
		if _t, ok := tryErr.(*tea.SDKError); ok {
			error = _t
		} else {
			error.Message = tea.String(tryErr.Error())
		}
		console.Log(error.Message)
	}
	return _err
}

/**
* UpdateDomainGroup  更新域名组名称
 * @param client         客户端
 * @param groupId        解析组ID
 * @param groupName      新域名组名称
 * @throws Exception
*/
func UpdateDomainGroup(client *dns.Client, groupId *string, groupName *string) (_err error) {
	req := &dns.UpdateDomainGroupRequest{}
	req.GroupId = groupId
	req.GroupName = groupName
	console.Log(tea.String("更新域名组的结果(json)↓"))
	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		resp, _err := client.UpdateDomainGroup(req)
		if _err != nil {
			return _err
		}

		console.Log(util.ToJSONString(tea.ToMap(resp)))

		return nil
	}()

	if tryErr != nil {
		var error = &tea.SDKError{}
		if _t, ok := tryErr.(*tea.SDKError); ok {
			error = _t
		} else {
			error.Message = tea.String(tryErr.Error())
		}
		console.Log(error.Message)
	}
	return _err
}

/**
* DeleteDomainGroup  删除域名组
 * @param client      客户端
 * @param groupId     域名组ID
 * @throws Exception
*/
func DeleteDomainGroup(client *dns.Client, groupId *string) (_err error) {
	req := &dns.DeleteDomainGroupRequest{}
	req.GroupId = groupId
	console.Log(tea.String("删除域名组的结果(json)↓"))
	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		resp, _err := client.DeleteDomainGroup(req)
		if _err != nil {
			return _err
		}

		console.Log(util.ToJSONString(tea.ToMap(resp)))

		return nil
	}()

	if tryErr != nil {
		var error = &tea.SDKError{}
		if _t, ok := tryErr.(*tea.SDKError); ok {
			error = _t
		} else {
			error.Message = tea.String(tryErr.Error())
		}
		console.Log(error.Message)
	}
	return _err
}

func _main(args []*string) (_err error) {

	// 初始化客户端
	access_key_id := "xxxx"
	access_key_secret := "xxx"
	region_id := "shanghai"
	client, _err := Init(&access_key_id, &access_key_secret, &region_id)
	if _err != nil {
		log.Fatal(_err)
	}
	// 查询账户下域名
	_err = DescribeDomains(client)
	domain := "occn.top"
	_err = DescribeDomainRecords(client, &domain)

	//regionId := args[7]   // 地域id
	//domainName := args[0] // 域名
	//RR := args[1]         // 前缀
	//recordType := args[2] // 记录类型
	//value := args[3]      // 待设置的值
	//recordId := args[4]   // 记录id
	//groupName := args[5]  // 组名
	//groupId := args[6]    // 组id
	//// 0.初始化客户端
	//client, _err := Init(env.GetEnv(tea.String("ACCESS_KEY_ID")), env.GetEnv(tea.String("ACCESS_KEY_SECRET")), regionId)
	//if _err != nil {
	//	return _err
	//}
	//
	//// 1.查询账户下域名
	//_err = DescribeDomains(client)
	//if _err != nil {
	//	return _err
	//}
	//// 2.阿里云云解析添加域名
	//_err = AddDomain(client, domainName)
	//if _err != nil {
	//	return _err
	//}
	//// 3.查询域名解析记录
	//_err = DescribeDomainRecords(client, domainName)
	//if _err != nil {
	//	return _err
	//}
	//// 4.查询域名记录日志
	//_err = DescribeRecordLogs(client, domainName)
	//if _err != nil {
	//	return _err
	//}
	//// 5.通过RecordId查询域名解析记录
	//_err = DescribeDomainRecordByRecordId(client, recordId)
	//if _err != nil {
	//	return _err
	//}
	//// 6.查询域名信息
	//_err = DescribeDomainInfo(client, domainName)
	//if _err != nil {
	//	return _err
	//}
	//// 7.添加域名解析记录
	//_err = AddDomainRecord(client, domainName, RR, recordType, value)
	//if _err != nil {
	//	return _err
	//}
	//// 8.更新域名解析记录
	//_err = UpdateDomainRecord(client, recordId, RR, recordType, value)
	//if _err != nil {
	//	return _err
	//}
	//// 9.设置域名解析状态
	//_err = SetDomainRecordStatus(client, recordId, tea.String("ENABLE"))
	//if _err != nil {
	//	return _err
	//}
	//// 10.删除域名解析记录
	//_err = DeleteDomainRecord(client, recordId)
	//if _err != nil {
	//	return _err
	//}
	//// 11.查询域名组
	//_err = DescribeDomainGroups(client)
	//if _err != nil {
	//	return _err
	//}
	//// 12.添加域名组
	//_err = AddDomainGroup(client, groupName)
	//if _err != nil {
	//	return _err
	//}
	//// 13.更新域名组名称
	//_err = UpdateDomainGroup(client, groupId, groupName)
	//if _err != nil {
	//	return _err
	//}
	//// 14.删除域名组
	//_err = DeleteDomainGroup(client, groupId)
	//if _err != nil {
	//	return _err
	//}
	//return _err
	return nil
}

func main() {
	err := _main(tea.StringSlice(os.Args[1:]))
	if err != nil {
		panic(err)
	}
}
