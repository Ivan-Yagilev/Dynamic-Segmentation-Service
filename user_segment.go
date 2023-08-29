package segmentation_service

type UserSegment struct {
	UserId      int      `json:"id" db:"id" binding:"required"`
	Segmentlist []string `json:"segmentlist" binding:"required"`
}
