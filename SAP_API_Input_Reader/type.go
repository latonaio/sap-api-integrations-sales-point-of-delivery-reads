package sap_api_input_reader

type EC_MC struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	Document      struct {
		DocumentNo     string `json:"document_no"`
		DeliverTo      string `json:"deliver_to"`
		Quantity       string `json:"quantity"`
		PickedQuantity string `json:"picked_quantity"`
		Price          string `json:"price"`
		Batch          string `json:"batch"`
	} `json:"document"`
	ProductionOrder struct {
		DocumentNo           string `json:"document_no"`
		Status               string `json:"status"`
		DeliverTo            string `json:"deliver_to"`
		Quantity             string `json:"quantity"`
		CompletedQuantity    string `json:"completed_quantity"`
		PlannedStartDate     string `json:"planned_start_date"`
		PlannedValidatedDate string `json:"planned_validated_date"`
		ActualStartDate      string `json:"actual_start_date"`
		ActualValidatedDate  string `json:"actual_validated_date"`
		Batch                string `json:"batch"`
		Work                 struct {
			WorkNo                   string `json:"work_no"`
			Quantity                 string `json:"quantity"`
			CompletedQuantity        string `json:"completed_quantity"`
			ErroredQuantity          string `json:"errored_quantity"`
			Component                string `json:"component"`
			PlannedComponentQuantity string `json:"planned_component_quantity"`
			PlannedStartDate         string `json:"planned_start_date"`
			PlannedStartTime         string `json:"planned_start_time"`
			PlannedValidatedDate     string `json:"planned_validated_date"`
			PlannedValidatedTime     string `json:"planned_validated_time"`
			ActualStartDate          string `json:"actual_start_date"`
			ActualStartTime          string `json:"actual_start_time"`
			ActualValidatedDate      string `json:"actual_validated_date"`
			ActualValidatedTime      string `json:"actual_validated_time"`
		} `json:"work"`
	} `json:"production_order"`
	APISchema      string `json:"api_schema"`
	MaterialCode   string `json:"material_code"`
	Plant_Supplier string `json:"plant/supplier"`
	Stock          string `json:"stock"`
	DocumentType   string `json:"document_type"`
	DocumentNo     string `json:"document_no"`
	PlannedDate    string `json:"planned_date"`
	ValidatedDate  string `json:"validated_date"`
	Deleted        bool   `json:"deleted"`
}

type SDC struct {
	ConnectionKey      string `json:"connection_key"`
	Result             bool   `json:"result"`
	RedisKey           string `json:"redis_key"`
	Filepath           string `json:"filepath"`
	SalesPointOfDeliveryCollection struct {
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
	} `json:"SalesPointOfDeliveryCollection"`
	APISchema                string   `json:"api_schema"`
	Accepter                 []string `json:"accepter"`
	SalesPointOfDeliveryCode string   `json:"sales_point_of_delivery_code"`
	Deleted                  bool     `json:"deleted"`
}
