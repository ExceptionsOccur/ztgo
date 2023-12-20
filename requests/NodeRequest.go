package requests

func GetNodeType() string {
	return ZtRequestGet("/status")
}
