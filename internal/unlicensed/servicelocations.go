package unlicensed

import (
	"encoding/json"
	"net/http"

	"thomascrmbz.com/smappee/types"
)

type serviceLocationsResponse []types.ServiceLocation

func (c *Client) GetServiceLocations() ([]types.ServiceLocation, error) {
	req, err := c.newRequest(http.MethodGet, "https://dashboard.smappee.net/api/v10/user/servicelocations", nil)

	slResponse := serviceLocationsResponse{}
	json.NewDecoder(req.Body).Decode(&slResponse)

	return slResponse, err
}
