package requests

import (
	"encoding/json"
)

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
