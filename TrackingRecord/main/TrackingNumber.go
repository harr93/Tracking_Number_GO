package main

import "fmt"
import "container/list"

func main() {
	r1 := Range{1, 100000}
	fmt.Print(r1.getHi())
}

//import trackingNumber.Range.Relation 

type TrackingNumber struct {
	statusCode   string
	transferCode string
	TrackingNumberRange Range
	deleted bool
}
	func (obj TrackingNumber) TrackingNumber (low int, high int, status_code string, transfer_code string) {
		obj.statusCode = status_code;
		obj.transferCode = transfer_code;
		obj.TrackingNumberRange = Range(low, high);
		obj.deleted = false;
	}
	
	func (obj TrackingNumber) findStringFromTrackingNumber () string {
		return obj.TrackingNumberRange.getLo() + " "+obj.TrackingNumberRange.getHi() + " "+obj.statusCode + " "+obj.transferCode;
	}

	
	//type List struct{
	//TrackingNumber List;
	//anotherTrackingNumber List }
	
	func (obj TrackingNumber) List TrackingNumber compare(anotherTrackingNumber TrackingNumber ) 
	{
		newTrackingNumberRows List<TrackingNumber>  = new ArrayList<TrackingNumber>();
		
		Relation relation  = obj.TrackingNumberRange.classify(anotherTrackingNumber.getR());
		if(relation==Relation.SAME)
		{
			anotherTrackingNumber.setDeleted(true);
			//newTrackingNumberRows.add(this);
		}
		else if(relation==Relation.SUPERSET)
		{
			anotherTrackingNumber.setDeleted(true);
			//newTrackingNumberRows.add(new TrackingNumber(this.TrackingNumberRange.getLo(), anotherTrackingNumber.getR().getLo()-1, this.statusCode, this.transferCode));
			//newTrackingNumberRows.add(anotherTrackingNumber);
			//newTrackingNumberRows.add(new TrackingNumber(anotherTrackingNumber.getR().getHi()+1, this.TrackingNumberRange.getHi(), this.statusCode, this.transferCode));
			
		}
		else if (relation==Relation.SUBSET)
		{
			if (obj.statusCode == anotherTrackingNumber.statusCode && obj.transferCode == anotherTrackingNumber.transferCode) {
				obj.setDeleted(true);
			} else {
				obj.setDeleted(true);
				anotherTrackingNumber.setDeleted(true);
				newTrackingNumberRows.add(new TrackingNumber(anotherTrackingNumber.getR().getLo(), obj.getR().getLo()-1, anotherTrackingNumber.statusCode, anotherTrackingNumber.transferCode));
				newTrackingNumberRows.add(new TrackingNumber(obj.getR().getLo(), obj.getR().getHi(), obj.statusCode, obj.transferCode));
				newTrackingNumberRows.add(new TrackingNumber(obj.getR().getHi()+1, anotherTrackingNumber.getR().getHi(), anotherTrackingNumber.statusCode, anotherTrackingNumber.transferCode));
			}
			//newTrackingNumberRows.add(anotherTrackingNumber);
		}
		else if(relation == Relation.LESSOVERLAP)
		{
			obj.setDeleted(true);
			anotherTrackingNumber.setDeleted(true);
			//newTrackingNumberRows.add(anotherTrackingNumber);
			newTrackingNumberRows.add(new TrackingNumber(obj.getR().getLo(), obj.getR().getHi(), obj.statusCode, obj.transferCode));
			newTrackingNumberRows.add(new TrackingNumber(obj.getR().getHi()+1, anotherTrackingNumber.getR().getHi(), anotherTrackingNumber.statusCode, anotherTrackingNumber.transferCode));
			
		}
		else if(relation==Relation.MOREOVERLAP)
		{
			obj.setDeleted(true);
			anotherTrackingNumber.setDeleted(true);
			//newTrackingNumberRows.add(anotherTrackingNumber);
			newTrackingNumberRows.add(new TrackingNumber(anotherTrackingNumber.getR().getLo(), obj.getR().getLo()-1, anotherTrackingNumber.statusCode, anotherTrackingNumber.transferCode));
			newTrackingNumberRows.add(new TrackingNumber(obj.getR().getLo(), obj.getR().getHi(), obj.statusCode, obj.transferCode));
		}
		return newTrackingNumberRows;
	}
	
	
	func (obj TrackingNumber) isDeleted() bool {
		return obj.deleted
	}

	func (obj TrackingNumber)  setDeleted(deleted bool) {  //void??
		obj.deleted = deleted
	}

	func (obj TrackingNumber) getR() Range {
		return obj.TrackingNumberRange
	}

	func (obj TrackingNumber) getStatusCode() string{
		return obj.statusCode
	}

	func (obj TrackingNumber) setStatusCode(statusCode string) {
		obj.statusCode = statusCode
	}

	func (obj TrackingNumber) getTransferCode() string {
		return obj.transferCode
	}

	func (obj TrackingNumber) setTransferCode(transferCode string) {
		obj.transferCode = transferCode
	}
	
}