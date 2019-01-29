// Code generated by go-swagger; DO NOT EDIT.

package access_manager

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new access manager API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for access manager API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
BindUserRole bind user role API
*/
func (a *Client) BindUserRole(params *BindUserRoleParams, authInfo runtime.ClientAuthInfoWriter) (*BindUserRoleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBindUserRoleParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "BindUserRole",
		Method:             "POST",
		PathPattern:        "/v1/user:role",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &BindUserRoleReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*BindUserRoleOK), nil

}

/*
CanDo can do API
*/
func (a *Client) CanDo(params *CanDoParams, authInfo runtime.ClientAuthInfoWriter) (*CanDoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCanDoParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CanDo",
		Method:             "POST",
		PathPattern:        "/v1/cando",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CanDoReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CanDoOK), nil

}

/*
CreateRole create role API
*/
func (a *Client) CreateRole(params *CreateRoleParams, authInfo runtime.ClientAuthInfoWriter) (*CreateRoleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateRoleParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateRole",
		Method:             "POST",
		PathPattern:        "/v1/roles",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateRoleReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateRoleOK), nil

}

/*
DeleteRoles delete roles API
*/
func (a *Client) DeleteRoles(params *DeleteRolesParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteRolesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteRolesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteRoles",
		Method:             "DELETE",
		PathPattern:        "/v1/roles",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteRolesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteRolesOK), nil

}

/*
DescribeRoles describe roles API
*/
func (a *Client) DescribeRoles(params *DescribeRolesParams, authInfo runtime.ClientAuthInfoWriter) (*DescribeRolesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDescribeRolesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DescribeRoles",
		Method:             "GET",
		PathPattern:        "/v1/roles",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DescribeRolesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DescribeRolesOK), nil

}

/*
GetRole admins permission
*/
func (a *Client) GetRole(params *GetRoleParams, authInfo runtime.ClientAuthInfoWriter) (*GetRoleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRoleParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetRole",
		Method:             "GET",
		PathPattern:        "/v1/roles/{role_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetRoleReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetRoleOK), nil

}

/*
GetRoleModule get role module API
*/
func (a *Client) GetRoleModule(params *GetRoleModuleParams, authInfo runtime.ClientAuthInfoWriter) (*GetRoleModuleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRoleModuleParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetRoleModule",
		Method:             "GET",
		PathPattern:        "/v1/roles:module",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetRoleModuleReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetRoleModuleOK), nil

}

/*
ModifyRole modify role API
*/
func (a *Client) ModifyRole(params *ModifyRoleParams, authInfo runtime.ClientAuthInfoWriter) (*ModifyRoleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewModifyRoleParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ModifyRole",
		Method:             "PATCH",
		PathPattern:        "/v1/roles",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ModifyRoleReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ModifyRoleOK), nil

}

/*
ModifyRoleModule modify role module API
*/
func (a *Client) ModifyRoleModule(params *ModifyRoleModuleParams, authInfo runtime.ClientAuthInfoWriter) (*ModifyRoleModuleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewModifyRoleModuleParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ModifyRoleModule",
		Method:             "PATCH",
		PathPattern:        "/v1/roles:module",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ModifyRoleModuleReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ModifyRoleModuleOK), nil

}

/*
UnbindUserRole groups
*/
func (a *Client) UnbindUserRole(params *UnbindUserRoleParams, authInfo runtime.ClientAuthInfoWriter) (*UnbindUserRoleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUnbindUserRoleParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UnbindUserRole",
		Method:             "DELETE",
		PathPattern:        "/v1/user:role",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UnbindUserRoleReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UnbindUserRoleOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}