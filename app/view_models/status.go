package viewmodels

import (
	statusModel "gin_weibo/app/models/status"
	"gin_weibo/pkg/time"
)

// StatusViewModel 微博
type StatusViewModel struct {
	ID        int
	Content   string
	UserID    int
	CreatedAt string
}

// NewStatusViewModelSerializer 微博数据展示
func NewStatusViewModelSerializer(s *statusModel.Status) *StatusViewModel {
	return &StatusViewModel{
		ID:        int(s.ID),
		Content:   s.Content,
		UserID:    int(s.UserID),
		CreatedAt: time.SinceForHuman(s.CreatedAt),
	}
}
