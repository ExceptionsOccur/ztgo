package requests

import (
	"encoding/json"
)

// 初始化为预设值，则可以判断是否是第一次启动，如果是第一次启动则不推送
var obsoleteMids map[string][]string = map[string][]string{"init": {"data"}}

func UpdateController(data string) string {
	var dataMap map[string]interface{}
	err := json.Unmarshal([]byte(data), &dataMap)
	if err != nil {
		return ""
	}
	url := "/controller/network/" + dataMap["nwid"].(string)
	return ZtRequestPost(url, data)
}

func CreateController(name string) string {
	nodeType := GetNodeType()
	nodeTypeJson := make(map[string]interface{})
	json.Unmarshal([]byte(nodeType), &nodeTypeJson)
	nodeAddr := nodeTypeJson["address"].(string)

	return ZtRequestPost("/controller/network/"+nodeAddr+"______", "{\"name\":\""+name+"\"}")
}

func DeleteController(nwid string) string {
	return ZtRequestDelete("/controller/network/"+nwid, `{}`)
}

func getControllerList() string {
	return ZtRequestGet("/controller/network")
}

func getControllerType(nwid string) string {
	return ZtRequestGet("/controller/network/" + nwid)
}

func GetAllControllerType() []string {
	var controllerArray []string
	var controllerTypeArray []string
	controllerListStr := getControllerList()
	err := json.Unmarshal([]byte(controllerListStr), &controllerArray)
	if err != nil {
		return nil
	}
	for _, nwid := range controllerArray {
		controllerTypeArray = append(controllerTypeArray, getControllerType(nwid))
	}
	return controllerTypeArray
}

func getMemberList(nwid string) []string {
	var mids []string
	resMap := make(map[string]int)
	resStr := ZtRequestGet("/controller/network/" + nwid + "/member")
	json.Unmarshal([]byte(resStr), &resMap)
	for k := range resMap {
		mids = append(mids, k)
	}

	return mids
}

func CountMembers() map[string]int {
	ctsMap := make(map[string]interface{})
	ctsArray := GetAllControllerType()
	resMap := make(map[string]int)
	for _, cts := range ctsArray {
		json.Unmarshal([]byte(cts), &ctsMap)
		resMap[ctsMap["nwid"].(string)] = len(GetAllMembersTypeByNwid(ctsMap["nwid"].(string)))
	}
	return resMap
}

func GetMemberType(nwid string, mid string) string {
	return ZtRequestGet("/controller/network/" + nwid + "/member/" + mid)
}

func GetAllMembersTypeByNwid(nwid string) []string {
	mids := getMemberList(nwid)
	var memberTypeArray []string
	for _, v := range mids {
		memberTypeArray = append(memberTypeArray, GetMemberType(nwid, v))
	}
	return memberTypeArray
}

func UpdateMember(data string) string {
	var dataMap map[string]interface{}
	err := json.Unmarshal([]byte(data), &dataMap)
	if err != nil {
		return ""
	}
	url := "/controller/network/" + dataMap["nwid"].(string) + "/member/" + dataMap["id"].(string)
	return ZtRequestPost(url, data)
}

func DeleteMember(nwid string, mid string) string {
	url := "/controller/network/" + nwid + "/member/" + mid
	return ZtRequestDelete(url, `{}`)
}

func find(slice []string, item string) (int, bool) {
	for i, value := range slice {
		if value == item {
			return i, true
		}
	}
	return -1, false
}

func CheckMembers() (map[string][]string, map[string][]string) {

	mids := make(map[string][]string)

	// 获取控制器列表
	var controllerArray []string
	controllerListStr := getControllerList()
	err := json.Unmarshal([]byte(controllerListStr), &controllerArray)
	if err != nil {
		return nil, nil
	}
	for _, nwid := range controllerArray {
		mids[nwid] = getMemberList(nwid)
	}
	// 程序第一次启动，不做操作，仅更新作为对比的旧数组
	if len(obsoleteMids["init"]) != 0 {
		obsoleteMids = mids
		return nil, nil
	}

	// 新旧列表都为空，不做操作
	if len(mids) == 0 && len(obsoleteMids) == 0 {
		return nil, nil
	}
	diffJoin := make(map[string][]string)
	diffLeave := make(map[string][]string)

	// 在实际存在的控制器上做对比
	for key, value := range mids {
		if len(value) == 0 && len(obsoleteMids[key]) == 0 {
			continue
		}
		if len(value) == 0 && len(obsoleteMids[key]) > 0 {
			diffLeave[key] = obsoleteMids[key]
			continue
		}
		if len(obsoleteMids[key]) == 0 && len(value) > 0 {
			diffJoin[key] = value
			continue
		}
		// 遍历新的成员列表
		for _, item := range value {
			// 查找旧列表
			_, exist := find(obsoleteMids[key], item)
			// 如果不存在，则该项是新增的
			if !exist {
				diffJoin[key] = append(diffJoin[key], item)
			}
		}

		// 遍历旧的成员列表
		for _, item := range obsoleteMids[key] {
			// 查找新列表
			_, exist := find(value, item)
			// 如果不存在，则该项是离开的
			if !exist {
				diffLeave[key] = append(diffLeave[key], item)
			}
		}

	}
	// 处理完后本次数据变为旧数据
	obsoleteMids = mids
	return diffJoin, diffLeave
}
