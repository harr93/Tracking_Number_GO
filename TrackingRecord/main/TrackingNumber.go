package main

import "fmt"
import "container/list"


//import trackingNumber.Range.Relation 

type TrackingNumber struct{
	statusCode char;
	transferCode char;
	TrackingNumberRange Range
	deleted boolean;
	
	func (obj TrackingNumber) TrackingNumber (low int, high int, status_code char, transfer_code char) {
		obj.statusCode = status_code;
		obj.transferCode = transfer_code;
		TrackingNumberRange = new Range(low, high);
		obj.deleted = false;
	}
	
	func (obj TrackingNumber) findStringFromTrackingNumber () String {
		return TrackingNumberRange.getLo() + " "+TrackingNumberRange.getHi() + " "+this.statusCode + " "+this.transferCode;
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
	
	
	func (obj TrackingNumber) isDeleted() boolean {
		return deleted
	}

	func (obj TrackingNumber)  setDeleted(deleted boolean) {  //void??
		obj.deleted = deleted
	}

	func (obj TrackingNumber) getR() Range {
		return TrackingNumberRange
	}

	func (obj TrackingNumber) getStatusCode() char{
		return statusCode
	}

	func (obj TrackingNumber) setStatusCode(statusCode char) {
		obj.statusCode = statusCode
	}

	func (obj TrackingNumber) getTransferCode() char {
		return transferCode
	}

	func (obj TrackingNumnber) setTransferCode(transferCode char) {
		obj.transferCode = transferCode
	}
	
}
