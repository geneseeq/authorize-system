// Package control contains the heart of the domain model.
package control

import (
	"github.com/geneseeq/authorize-system/task/baseJob/model"
	"github.com/geneseeq/authorize-system/task/baseJob/service"
	"gopkg.in/mgo.v2/bson"
)

func FindAllLabCode() (error, []string) {
	var dropInstance = service.NewBaseDBRepository("GENESEEQ", "SAMPLE_RECEIVE")
	var sampleInstace = service.NewBaseDBRepository("GENESEEQ", "SAMPLE_RECEIVE_ITEM")
	dropList, err := dropInstance.Distinct("ID", bson.M{"STATE": "0"})
	if err != nil {
		println("find drop failed.")
		return err, nil
	}
	var labeCodeList []string
	tempMap := map[string][]string{"$nin": dropList}
	condition := bson.M{}
	condition["SAMPLE_RECEIVE"] = tempMap
	labeCodeList, err = sampleInstace.Distinct("LAB_CODE", condition)
	if err != nil {
		println("find lab code failed.")
		return err, nil
	}
	return nil, labeCodeList
}

func FindSampleBaseInfo(labeCodeList []string) ([]model.BaseInfoModel, error) {
	// var pipeline []bson.M
	tmpMatch := bson.M{"LAB_CODE": bson.M{"$in": labeCodeList}}
	tmpLookup := bson.M{"from": "DIC_SAMPLE_TYPE", "localField": "DIC_SAMPLE_TYPE", "foreignField": "ID", "as": "sample_advanced_info"}
	pipeline := []bson.M{
		bson.M{"$match": tmpMatch},
		bson.M{"$lookup": tmpLookup}}
	var sampleInstace = service.NewBaseDBRepository("GENESEEQ", "SAMPLE_RECEIVE_ITEM")
	result, err := sampleInstace.Aggregate(&pipeline)
	if err == nil {
		return result, nil
	}
	return nil, nil
}

func FindSampleOrderInfo(OrderID string) []model.BaseInfoModel {
	var sampleOrderList = []string{OrderID}
	tmpMatch := bson.M{"ID": bson.M{"$in": sampleOrderList}}
	tmpMedicaLlookup := bson.M{"from": "CRM_PATIENT", "localField": "MEDICAL_NUMBER", "foreignField": "ID", "as": "patient_info"}
	tmpConsumerLookup := bson.M{"from": "CRM_CONSUMER_MARKET", "localField": "ID", "foreignField": "SAMPLE_ORDER_ID", "as": "consumer_info"}
	pipeline := []bson.M{
		bson.M{"$match": tmpMatch},
		bson.M{"$lookup": tmpMedicaLlookup},
		bson.M{"$lookup": tmpConsumerLookup}}
	var sampleInstace = service.NewBaseDBRepository("GENESEEQ", "SAMPLE_ORDER")
	result, err := sampleInstace.Aggregate(&pipeline)
	if err == nil {
		return result
	}
	return nil

}
