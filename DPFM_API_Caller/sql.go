package dpfm_api_caller

import (
	dpfm_api_input_reader "data-platform-api-storage-bin-deletes-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-storage-bin-deletes-rmq-kube/DPFM_API_Output_Formatter"
	"fmt"
	"strings"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func (c *DPFMAPICaller) GeneralRead(
	input *dpfm_api_input_reader.SDC,
	log *logger.Logger,
) *dpfm_api_output_formatter.General {

	where := strings.Join([]string{
		fmt.Sprintf("WHERE general.BusinessPartner = %d ", input.General.BusinessPartner),
		fmt.Sprintf("AND general.Plant = \"%s\" ", input.General.Plant),
		fmt.Sprintf("AND general.StorageLocation = \"%s\" ", input.General.StorageLocation),
		fmt.Sprintf("AND general.StorageBin = \"%s\" ", input.General.StorageBin),
	}, "")

	rows, err := c.db.Query(
		`SELECT 
    	general.BusinessPartner,
    	general.Plant,
		general.StorageLocation,
    	general.StorageBin,
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_storage_bin_general_data as general 
		` + where + ` ;`)
	if err != nil {
		log.Error("%+v", err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToGeneral(rows)
	if err != nil {
		log.Error("%+v", err)
		return nil
	}

	return data
}
