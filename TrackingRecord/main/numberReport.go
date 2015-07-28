package main

import "container/list"
import "sort"

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