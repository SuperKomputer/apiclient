// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new operations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for operations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreateAccount creates payable account of a user on cluster

creates payable account of a user on cluster
*/
func (a *Client) CreateAccount(params *CreateAccountParams) (*CreateAccountAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateAccountParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createAccount",
		Method:             "POST",
		PathPattern:        "/clusters/{clusterId}/account/{username}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateAccountReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateAccountAccepted), nil

}

/*
CreateCluster creates a cluster

creates a cluster
*/
func (a *Client) CreateCluster(params *CreateClusterParams) (*CreateClusterAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateClusterParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createCluster",
		Method:             "POST",
		PathPattern:        "/clusters",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateClusterReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateClusterAccepted), nil

}

/*
CreateUser creates a user

creates a user
*/
func (a *Client) CreateUser(params *CreateUserParams) (*CreateUserAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateUserParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createUser",
		Method:             "POST",
		PathPattern:        "/users",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateUserReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateUserAccepted), nil

}

/*
DeleteAccount deletes payable account of a user on cluster

deletes payable account of a user on cluster
*/
func (a *Client) DeleteAccount(params *DeleteAccountParams) (*DeleteAccountAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteAccountParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteAccount",
		Method:             "DELETE",
		PathPattern:        "/clusters/{clusterId}/account/{username}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteAccountReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteAccountAccepted), nil

}

/*
DeleteCluster deletes a cluster

deletes a cluster with the given name
*/
func (a *Client) DeleteCluster(params *DeleteClusterParams) (*DeleteClusterAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteClusterParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteCluster",
		Method:             "DELETE",
		PathPattern:        "/clusters/{clusterId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteClusterReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteClusterAccepted), nil

}

/*
DeleteUser deletes a user

deletes a user with the given name
*/
func (a *Client) DeleteUser(params *DeleteUserParams) (*DeleteUserAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteUserParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteUser",
		Method:             "DELETE",
		PathPattern:        "/users/{username}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteUserReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteUserAccepted), nil

}

/*
FetchUserClusters fetches the list of using clusters of user with given username
*/
func (a *Client) FetchUserClusters(params *FetchUserClustersParams) (*FetchUserClustersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFetchUserClustersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "fetchUserClusters",
		Method:             "GET",
		PathPattern:        "/users/{username}/clusters",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &FetchUserClustersReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*FetchUserClustersOK), nil

}

/*
GetAccount gets payable account of a user on cluster

get payable account of a user on cluster
*/
func (a *Client) GetAccount(params *GetAccountParams) (*GetAccountOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAccountParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getAccount",
		Method:             "GET",
		PathPattern:        "/clusters/{clusterId}/account/{username}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAccountReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetAccountOK), nil

}

/*
GetCluster gets a cluster

get a cluster details with clusterId
*/
func (a *Client) GetCluster(params *GetClusterParams) (*GetClusterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetClusterParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getCluster",
		Method:             "GET",
		PathPattern:        "/clusters/{clusterId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetClusterReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetClusterOK), nil

}

/*
GetUser gets details information of a user with given username

get details information of a user with given username
*/
func (a *Client) GetUser(params *GetUserParams) (*GetUserOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetUserParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getUser",
		Method:             "GET",
		PathPattern:        "/users/{username}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetUserReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetUserOK), nil

}

/*
ListClusters gets a list of all clusters

get a list of all clusters managed by the SuperKomputer
*/
func (a *Client) ListClusters(params *ListClustersParams) (*ListClustersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListClustersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listClusters",
		Method:             "GET",
		PathPattern:        "/clusters",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListClustersReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListClustersOK), nil

}

/*
ListUsers gets a list of known users
*/
func (a *Client) ListUsers(params *ListUsersParams) (*ListUsersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListUsersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listUsers",
		Method:             "GET",
		PathPattern:        "/users",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListUsersReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListUsersOK), nil

}

/*
Login logs in to the super komputer service

## create a token which is used to provide identity

*/
func (a *Client) Login(params *LoginParams) (*LoginOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLoginParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "login",
		Method:             "POST",
		PathPattern:        "/login",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &LoginReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*LoginOK), nil

}

/*
UpdateCluster updates a cluster

updates a cluster with the given update config
*/
func (a *Client) UpdateCluster(params *UpdateClusterParams) (*UpdateClusterAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateClusterParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateCluster",
		Method:             "PUT",
		PathPattern:        "/clusters/{clusterId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateClusterReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateClusterAccepted), nil

}

/*
UpdateUser updates a user

updates a user with the given update config
*/
func (a *Client) UpdateUser(params *UpdateUserParams) (*UpdateUserAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateUserParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateUser",
		Method:             "PUT",
		PathPattern:        "/users/{username}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateUserReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateUserAccepted), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
