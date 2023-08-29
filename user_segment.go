package segmentation_service

type UserSegment struct {
	UserId      int    `json:"id" db:"id" binding:"required"`
	Segmentname string `json:"segmentname" db:"segmentname" binding:"required"`
}
