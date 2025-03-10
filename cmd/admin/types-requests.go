package main

// LoginRequest to receive login credentials
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// LogoutRequest to receive logout requests
type LogoutRequest struct {
	CSRFToken string `json:"csrftoken"`
}

// DistributedQueryRequest to receive query requests
type DistributedQueryRequest struct {
	CSRFToken    string   `json:"csrftoken"`
	Environments []string `json:"environment_list"`
	Platforms    []string `json:"platform_list"`
	UUIDs        []string `json:"uuid_list"`
	Hosts        []string `json:"host_list"`
	Query        string   `json:"query"`
	Repeat       int      `json:"repeat"`
}

// DistributedCarveRequest to receive carve requests
type DistributedCarveRequest struct {
	CSRFToken    string   `json:"csrftoken"`
	Environments []string `json:"environment_list"`
	Platforms    []string `json:"platform_list"`
	UUIDs        []string `json:"uuid_list"`
	Hosts        []string `json:"host_list"`
	Path         string   `json:"path"`
	Repeat       int      `json:"repeat"`
}

// DistributedQueryActionRequest to receive query requests
type DistributedQueryActionRequest struct {
	CSRFToken string   `json:"csrftoken"`
	Names     []string `json:"names"`
	Action    string   `json:"action"`
}

// DistributedCarvesActionRequest to receive carves requests
type DistributedCarvesActionRequest struct {
	CSRFToken string   `json:"csrftoken"`
	IDs       []string `json:"ids"`
	Action    string   `json:"action"`
}

// NodeMultiActionRequest to receive node action requests
type NodeMultiActionRequest struct {
	CSRFToken string   `json:"csrftoken"`
	Action    string   `json:"action"`
	UUIDs     []string `json:"uuids"`
}

// SettingsRequest to receive changes to settings
type SettingsRequest struct {
	CSRFToken string `json:"csrftoken"`
	Action    string `json:"action"`
	Boolean   bool   `json:"boolean"`
	Type      string `json:"type"`
	Name      string `json:"name"`
	Value     string `json:"value"`
}

// ConfigurationRequest to receive changes to configuration
type ConfigurationRequest struct {
	CSRFToken        string `json:"csrftoken"`
	ConfigurationB64 string `json:"configuration"`
}

// IntervalsRequest to receive changes to intervals
type IntervalsRequest struct {
	CSRFToken      string `json:"csrftoken"`
	ConfigInterval int    `json:"config"`
	LogInterval    int    `json:"log"`
	QueryInterval  int    `json:"query"`
}

// ExpirationRequest to receive expiration changes to enroll/remove nodes
type ExpirationRequest struct {
	CSRFToken string `json:"csrftoken"`
	Action    string `json:"action"`
	Type      string `json:"type"`
}

// EnvironmentsRequest to receive changes to environments
type EnvironmentsRequest struct {
	CSRFToken string `json:"csrftoken"`
	Action    string `json:"action"`
	Name      string `json:"name"`
	Hostname  string `json:"hostname"`
	Type      string `json:"type"`
	Icon      string `json:"icon"`
	DebugHTTP bool   `json:"debughttp"`
}

// UsersRequest to receive user action requests
type UsersRequest struct {
	CSRFToken string `json:"csrftoken"`
	Action    string `json:"action"`
	Username  string `json:"username"`
	Fullname  string `json:"fullname"`
	Password  string `json:"password"`
	Admin     bool   `json:"admin"`
}

// AdminResponse to be returned to requests
type AdminResponse struct {
	Message string `json:"message"`
}
