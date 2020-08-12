package main

func init() {

}

////////////////////////////////////////
// M E T A D A T A
////////////////////////////////////////
type Metadata struct {
	TypeInfo struct {
		TypeConstraints            map[string]map[string]bool   `json:"typeConstraints"`
		TypeConstraintDisplayNames map[string]map[string]string `json:"typeConstraintDisplayNames"`
	} `json:"typeInfo"`
}

type GetMetaDataRequest struct {
	Username    string `json:"username"`
	Token       string `json:"token"`
	RequestType string `json:"type"`
}

type TDGMetadataResponse struct {
	responseType string   `json:"type"`
	metadata     Metadata `json:"metadata"`
	Error        string   `json:"error"`
}

func CreateGetMetadataRequest(token string) GetMetaDataRequest {
	return GetMetaDataRequest{Username: "javaconductor@yahoo.com", RequestType: "getUserDataSets", Token: token}
}

////////////////////////////////////////
// U S E R   D A T A   S E T S
////////////////////////////////////////
type GetUserDataSetsResponse struct {
	Type         string        `json:"type"`
	Username     string        `json:"username"`
	DataSetSpecs []DataSetSpec `json:"dataSetSpecs"`
	Error        string        `json:"error"`
}

type TDGUserDataSetsRequest struct {
	Username    string `json:"username"`
	Token       string `json:"token"`
	RequestType string `json:"type"`
}

func CreateUserDataSetsRequest(username string, token string) TDGUserDataSetsRequest {
	return TDGUserDataSetsRequest{
		Username:    username,
		RequestType: "getUserDataSets",
		Token:       token,
	}
}

////////////////////////////////////////
////////////////////////////////////////
