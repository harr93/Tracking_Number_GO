package main

import "container/list"
import (
	"sort"
	"strings"
)

type TrackingNumberReport struct {
	caseName string
	trackingNumberList []TrackingNumber
}

func (obj TrackingNumberReport) returnReport() *list.List {

	finalList := *list.List{}
	for tNum := range obj.trackingNumberList {
		if(tNum.isNotDeleted()) {
			finalList.PushBack(tNum.findStringFromTrackingNumber())
		}
	}
	sort.Sort(finalList)
	finalList.PushFront(caseName)
	return finalList
}

func (obj TrackingNumberReport) generateTuple(tuple string) {

	tupleArray := strings.Split(tuple, " ")
	low := tupleArray[0]
	high := tupleArray[1]
	statusCode := tupleArray[2]
	transferCode := tupleArray[3]

	tN = TrackingNumber {low, high, statusCode, transferCode}
	checkRangeAndStatus(tN)
}

func (obj TrackingNumberReport) checkRangeAndStatus(tN TrackingNumber) {
	mutatedList = *list.List{}
	for t := range obj.trackingNumberList {
		if(tN.isDeleted()) {
			break
		}
		if(!t.isDeleted()) {
			mutatedList.addAll(tN.compare(t))
		}
	}

	if(mutatedList != null) {
		for t:= range mutatedList {
				obj.checkRangeAndStatus(t)
		}
		obj.trackingNumberList.removeAll(mutatedList)
		obj.trackingNumberList.addAll(mutatedList)
	}
	if(!tN.isDeleted()&& !obj.trackingNumberList.contains(tN)){
		obj.trackingNumberList.add(tN)
	}
}