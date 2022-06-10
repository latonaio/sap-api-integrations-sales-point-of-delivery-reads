package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-sales-point-of-delivery-reads/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToSalesPointOfDeliveryCollection(raw []byte, l *logger.Logger) ([]SalesPointOfDeliveryCollection, error) {
	pm := &responses.SalesPointOfDeliveryCollection{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to SalesPointOfDeliveryCollection. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	salesPointOfDeliveryCollection := make([]SalesPointOfDeliveryCollection, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		salesPointOfDeliveryCollection = append(salesPointOfDeliveryCollection, SalesPointOfDeliveryCollection{
			ObjectID:                    data.ObjectID,                                                       
			CreationDateTime:            data.CreationDateTime,                                               
			CreationIdentityUUID:        data.CreationIdentityUUID,                                           
			LastChangeDateTime:          data.LastChangeDateTime,                                             
			LastChangeIdentityUUID:      data.LastChangeIdentityUUID,                                         
			SalesPODDescription:         data.SalesPODDescription,                                            
			FromDate:                    data.FromDate,                                                       
			ToDate:                      data.ToDate,                                                         
			Grid:                        data.Grid,                                                           
			VoltageLevel:                data.VoltageLevel,                                                   
			PODTechRole:                 data.PODTechRole,                                                    
			PODDeregRole:                data.PODDeregRole,                                                   
			PODType:                     data.PODType,                                                        
			ServiceProvider:             data.ServiceProvider,                                                
			Distributor:                 data.Distributor,                                                    
			ServiceType:                 data.ServiceType,                                                    
			Installation:                data.Installation,                                                   
			PremiseExternalID:           data.PremiseExternalID,                                              
			HouseID:                     data.HouseID,                                                        
			BuildingID:                  data.BuildingID,                                                     
			RoomID:                      data.RoomID,                                                         
			FloorID:                     data.FloorID,                                                        
			StreetName:                  data.StreetName,                                                     
			RegionCode:                  data.RegionCode,                                                     
			CityName:                    data.CityName,                                                       
			CountryCode:                 data.CountryCode,                                                    
			StreetPostalCode:            data.StreetPostalCode,                                               
			UUID:                        data.UUID,                                                           
			SalesPointOfDeliveryID:      data.SalesPointOfDeliveryID,                                         
			EntityLastChangedOn:         data.EntityLastChangedOn,                                            
			ETag:                        data.ETag,                                                           
			Released:                    data.Released,                                                       
		})
	}

	return salesPointOfDeliveryCollection, nil
}
