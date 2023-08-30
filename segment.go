package segmentation_service

import "errors"

type Segment struct {
	Id          int    `json:"id" db:"id"`
	Segmentname string `json:"segmentname" db:"segmentname" binding:"required"`
}

type SegmentWOId struct {
	Segmentname string `json:"segmentname" db:"segmentname" binding:"required"`
}

type UpdateSegmentInput struct {
	Segmentname *string `json:"segmentname"`
}

func (i UpdateSegmentInput) Validate() error {
	if i.Segmentname == nil {
		return errors.New("no values to update")
	}
	return nil
}
