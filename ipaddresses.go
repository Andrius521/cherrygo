package cherrygo

import (
	"fmt"
)

const baseIPSPath = "/v1/projects"

// IpAddressesService is an interface for interfacing with the the Server endpoints of the CherryServers API
// See: https://api.cherryservers.com/doc/#tag/Ip-Addresses
type IpAddressesService interface {
	List(projectID int, opts *GetOptions) ([]IPAddress, *Response, error)
	Get(projectID int, ipID string, opts *GetOptions) (IPAddress, *Response, error)
	Create(projectID int, request *CreateIPAddress) (IPAddress, *Response, error)
	Remove(projectID int, ipID string) (IPAddress, *Response, error)
	Update(projectID int, ipID string, request *UpdateIPAddress) (IPAddress, *Response, error)
}

// IPAddresses fields
type IPAddress struct {
	ID            string            `json:"id,omitempty"`
	Address       string            `json:"address,omitempty"`
	AddressFamily int               `json:"address_family,omitempty"`
	Cidr          string            `json:"cidr,omitempty"`
	Gateway       string            `json:"gateway,omitempty"`
	Type          string            `json:"type,omitempty"`
	Region        Region            `json:"region,omitempty"`
	RoutedTo      RoutedTo          `json:"routed_to,omitempty"`
	AssignedTo    AssignedTo        `json:"assigned_to,omitempty"`
	TargetedTo    AssignedTo        `json:"targeted_to,omitempty"`
	Project       Project           `json:"project,omitempty"`
	PtrRecord     string            `json:"ptr_record,omitempty"`
	ARecord       string            `json:"a_record,omitempty"`
	Tags          map[string]string `json:"tags,omitempty"`
	Href          string            `json:"href,omitempty"`
}

// RoutedTo fields
type RoutedTo struct {
	ID            string `json:"id,omitempty"`
	Address       string `json:"address,omitempty"`
	AddressFamily int    `json:"address_family,omitempty"`
	Cidr          string `json:"cidr,omitempty"`
	Gateway       string `json:"gateway,omitempty"`
	Type          string `json:"type,omitempty"`
	Region        Region `json:"region,omitempty"`
}

// AssignedTo fields
type AssignedTo struct {
	ID       int     `json:"id,omitempty"`
	Name     string  `json:"name,omitempty"`
	Href     string  `json:"href,omitempty"`
	Hostname string  `json:"hostname,omitempty"`
	Image    string  `json:"image,omitempty"`
	Region   Region  `json:"region,omitempty"`
	State    string  `json:"state,omitempty"`
	Pricing  Pricing `json:"pricing,omitempty"`
}

// IPClient paveldi client
type IPsClient struct {
	client *Client
}

// CreateIPAddress fields for adding addition IP address
type CreateIPAddress struct {
	Type       string             `json:"type,omitempty"`
	Region     string             `json:"region,omitempty"`
	PtrRecord  string             `json:"ptr_record,omitempty"`
	ARecord    string             `json:"a_record,omitempty"`
	RoutedTo   string             `json:"routed_to,omitempty"`
	AssignedTo string             `json:"assigned_to,omitempty"`
	TargetedTo string             `json:"targeted_to,omitempty"`
	Tags       *map[string]string `json:"tags,omitempty"`
}

// UpdateIPAddress fields for updating IP address
type UpdateIPAddress struct {
	PtrRecord  string             `json:"ptr_record,omitempty"`
	ARecord    string             `json:"a_record,omitempty"`
	RoutedTo   string             `json:"routed_to,omitempty"`
	AssignedTo string             `json:"assigned_to,omitempty"`
	TargetedTo string             `json:"targeted_to,omitempty"`
	Tags       *map[string]string `json:"tags,omitempty"`
}

// List func lists ip addresses
func (i *IPsClient) List(projectID int, opts *GetOptions) ([]IPAddress, *Response, error) {
	path := opts.WithQuery(fmt.Sprintf("%s/%d/ips", baseIPSPath, projectID))

	var trans []IPAddress

	resp, err := i.client.MakeRequest("GET", path, nil, &trans)
	if err != nil {
		err = fmt.Errorf("Error: %v", err)
	}

	return trans, resp, err
}

// List func lists teams
func (i *IPsClient) Get(projectID int, ipID string, opts *GetOptions) (IPAddress, *Response, error) {
	path := opts.WithQuery(fmt.Sprintf("%s/%d/ips/%s", baseIPSPath, projectID, ipID))

	var trans IPAddress

	resp, err := i.client.MakeRequest("GET", path, nil, &trans)
	if err != nil {
		err = fmt.Errorf("Error: %v", err)
	}

	return trans, resp, err
}

// Create function orders new floating IP address
func (i *IPsClient) Create(projectID int, request *CreateIPAddress) (IPAddress, *Response, error) {
	var trans IPAddress

	path := fmt.Sprintf("%s/%d/ips", baseIPSPath, projectID)

	resp, err := i.client.MakeRequest("POST", path, request, &trans)
	if err != nil {
		err = fmt.Errorf("Error: %v", err)
	}

	return trans, resp, err
}

// Update function updates existing IP address
func (i *IPsClient) Update(projectID int, ipID string, request *UpdateIPAddress) (IPAddress, *Response, error) {
	var trans IPAddress

	path := fmt.Sprintf("%s/%d/ips/%s", baseIPSPath, projectID, ipID)

	resp, err := i.client.MakeRequest("PUT", path, request, &trans)
	if err != nil {
		err = fmt.Errorf("Error: %v", err)
	}

	return trans, resp, err
}

// Remove function removes existing project IP address
func (i *IPsClient) Remove(projectID int, ipID string) (IPAddress, *Response, error) {
	var trans IPAddress

	path := fmt.Sprintf("%s/%d/ips/%s", baseIPSPath, projectID, ipID)

	resp, err := i.client.MakeRequest("DELETE", path, nil, &trans)
	if err != nil {
		err = fmt.Errorf("Error: %v", err)
	}

	return trans, resp, err
}
