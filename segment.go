package segmentation_service

type Segment struct {
	Id          int    `json:"id" db:"id"`
	Segmentname string `json:"segmentname" db:"segmentname" binding:"required"`
}

type UserSegment struct {
	Id        int
	UserId    int
	SegmentId int
}
