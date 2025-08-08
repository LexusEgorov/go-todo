package client

type Client struct{}

func New() *Client {
	return &Client{}
}

//TODO: send tokens to auth for verification
//TODO: send user data to auth for register
//TODO: send login+pass to auth for authorization
