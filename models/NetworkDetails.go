package models

/*
|-----------field---------|------type------|--------------Description---------------------|------writable------|
		id					string					16-digit network ID								no
		nwid				string				16-digit network ID (legacy)						no
		objtype				string					Always "network"								no
		name				string				A short name for this network						YES
		creationTime		integer			 Time network record was created (ms since epoch)		no
		private				boolean			 		Is access control enabled?						YES
		enableBroadcast		boolean			 	Ethernet ff:ff:ff:ff:ff:ff allowed?					YES
		v4AssignMode		object			 IPv4 management and assign options (see below)			YES
		v6AssignMode		object			 IPv6 management and assign options (see below)			YES
		mtu					integer			 		Network MTU (default: 2800)						YES
		multicastLimit		integer		 	Maximum recipients for a multicast packet				YES
		revision			integer		 	Network config revision counter							no
		routes				array[object]	 		Managed IPv4 and IPv6 routes					YES
		ipAssignmentPools	array[object]	 		IP auto-assign ranges							YES
		rules				array[object]	 			Traffic rules								YES
		capabilities		array[object]	 	Array of capability objects (see below)				YES
		tags				array[object]	 		Array of tag objects (see below)				YES
		remoteTraceTarget	string		 	10-digit ZeroTier ID of remote trace target				YES
		remoteTraceLevel	integer		 		Remote trace verbosity level						YES

https://github.com/zerotier/ZeroTierOne/tree/master/controller
*/
type NetworkDetails struct {
	AuthTokens            interface{} `json:"authTokens"`
	AuthorizationEndpoint string      `json:"authorizationEndpoint"`
	Capabilities          interface{} `json:"capabilities"`
	ClientId              string      `json:"clientId"`
	CreationTime          int         `json:"creationTime"`
	Dns                   interface{} `json:"dns"`
	EnableBroadcast       bool        `json:"enableBroadcast"`
	Id                    string      `json:"id"`
	IpAssignmentPools     interface{} `json:"ipAssignmentPools"`
	Mtu                   int         `json:"mtu"`
	MulticastLimit        int         `json:"multicastLimit"`
	Name                  string      `json:"name"`
	Nwid                  string      `json:"nwid"`
	Objtype               string      `json:"objtype"`
	Private               bool        `json:"private"`
	RemoteTraceLevel      int         `json:"remoteTraceLevel"`
	RemoteTraceTarget     interface{} `json:"remoteTraceTarget"`
	Revision              int         `json:"revision"`
	Routes                interface{} `json:"routes"`
	Rules                 interface{} `json:"rules"`
	RulesSource           string      `json:"rulesSource"`
	SsoEnabled            bool        `json:"ssoEnabled"`
	Tags                  interface{} `json:"tags"`
	V4AssignMode          interface{} `json:"v4AssignMode"`
	V6AssignMode          interface{} `json:"v6AssignMode"`
}

type MenberDetails struct {
	ActiveBridge                 bool        `json:"activeBridge"`
	Address                      string      `json:"address"`
	AuthenticationExpiryTime     int         `json:"authenticationExpiryTime"`
	Authorized                   bool        `json:"authorized"`
	Capabilities                 interface{} `json:"capabilities"`
	CreationTime                 int         `json:"creationTime"`
	Id                           string      `json:"id"`
	Identity                     string      `json:"identity"`
	IpAssignments                interface{} `json:"ipAssignments"`
	LastAuthorizedCredential     interface{} `json:"lastAuthorizedCredential"`
	LastAuthorizedCredentialType string      `json:"lastAuthorizedCredentialType"`
	LastAuthorizedTime           int         `json:"lastAuthorizedTime"`
	LastDeauthorizedTime         int         `json:"lastDeauthorizedTime"`
	NoAutoAssignIps              bool        `json:"noAutoAssignIps"`
	Nwid                         string      `json:"nwid"`
	Objtype                      string      `json:"objtype"`
	RemoteTraceLevel             int         `json:"remoteTraceLevel"`
	RemoteTraceTarget            interface{} `json:"remoteTraceTarget"`
	Revision                     int         `json:"revision"`
	SsoExempt                    bool        `json:"ssoExempt"`
	Tags                         interface{} `json:"tags"`
	VMajor                       int         `json:"vMajor"`
	VMinor                       int         `json:"vMinor"`
	VProto                       int         `json:"vProto"`
	VRev                         int         `json:"vRev"`
}
