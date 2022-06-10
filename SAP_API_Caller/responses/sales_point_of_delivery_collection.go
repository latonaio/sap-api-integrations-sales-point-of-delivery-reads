package responses

type SalesPointOfDeliveryCollection struct {
	D struct {
		Results []struct {
			ObjectID                   string `json:"ObjectID"`                              
			CreationDateTime           string `json:"CreationDateTime"`                              
			CreationIdentityUUID       string `json:"CreationIdentityUUID"`                              
			LastChangeDateTime         string `json:"LastChangeDateTime"`                              
			LastChangeIdentityUUID     string `json:"LastChangeIdentityUUID"`                              
			SalesPODDescription        string `json:"SalesPODDescription"`                              
			FromDate                   string `json:"FromDate"`                              
			ToDate                     string `json:"ToDate"`                              
			Grid                       string `json:"Grid"`                              
			VoltageLevel               string `json:"VoltageLevel"`                              
			PODTechRole                string `json:"PODTechRole"`                              
			PODDeregRole               string `json:"PODDeregRole"`                              
			PODType                    string `json:"PODType"`                              
			ServiceProvider            string `json:"ServiceProvider"`                              
			Distributor                string `json:"Distributor"`                              
			ServiceType                string `json:"ServiceType"`                              
			Installation               string `json:"Installation"`                              
			PremiseExternalID          string `json:"PremiseExternalID"`                              
			HouseID                    string `json:"HouseID"`                              
			BuildingID                 string `json:"BuildingID"`                              
			RoomID                     string `json:"RoomID"`                              
			FloorID                    string `json:"FloorID"`                              
			StreetName                 string `json:"StreetName"`                              
			RegionCode                 string `json:"RegionCode"`                              
			CityName                   string `json:"CityName"`                              
			CountryCode                string `json:"CountryCode"`                              
			StreetPostalCode           string `json:"StreetPostalCode"`                              
			UUID                       string `json:"UUID"`                              
			SalesPointOfDeliveryID     string `json:"SalesPointOfDeliveryID"`                              
			EntityLastChangedOn        string `json:"EntityLastChangedOn"`                              
			ETag                       string `json:"ETag"`                              
			Released                   string `json:"Released"`                              
		} `json:results"`
	} `json:"d"`
}