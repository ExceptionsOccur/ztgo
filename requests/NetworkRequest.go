package requests

func GetAllNetworkType() string {
	return ZtRequestGet("/network")
}

func JoinToNetwork(nwid string) string {
	return ZtRequestPost("/network/"+nwid, `{}`)
}

func LeaveNetwork(nwid string) string {
	return ZtRequestDelete("/network/"+nwid, `{}`)
}
