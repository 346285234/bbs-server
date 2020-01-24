package gorm

import (
	"github.com/346285234/bbs-server/pkg/bbs"
	"github.com/jinzhu/gorm"
)

type CommentService struct {
	op commentOperation
}

func NewCommentService(db *gorm.DB) CommentService {
	op := newCommentOperation(db)
	return CommentService{op}
}

func (c *CommentService) List(topicID uint) (comments []*bbs.Comment, err error) {
	return c.op.list(topicID)
}

func (c *CommentService) Reply(comment bbs.Comment, parentID uint) (*bbs.Comment, error) {
	return c.op.add(comment, parentID)
}

func (c *CommentService) Revoke(topicID uint, id uint) (err error) {
	return c.op.remove(topicID, id)
}

type commentOperation struct {
	db *gorm.DB
}

func newCommentOperation(db *gorm.DB) commentOperation {
	return commentOperation{db}
}

func (c *commentOperation) list(topicID uint) (comments []*bbs.Comment, err error) {
	if err := c.db.Preload("Subs").Find(&comments).Error; err != nil {
		return nil, err
	}

	var subs = make(map[uint]bool)
	for i := 0; i < len(comments); {
		v := comments[i]
		if subs[v.ID] {
			// remove.
			comments = append(comments[:i], comments[i+1:]...)
			continue
		}
		if len(v.Subs) != 0 {
			for _, subV := range v.Subs {
				subs[subV.ID] = true
			}
		}
		i++
	}
	return comments, nil
}

func (c *commentOperation) get(id uint) (*bbs.Comment, error) {
	var result bbs.Comment
	if err := c.db.Where("id = ?", id).First(&result).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *commentOperation) add(comment bbs.Comment, parentID uint) (*bbs.Comment, error) {
	if err := c.db.Create(&comment).Error; err != nil {
		return nil, err
	}

	if parentID == 0 {
		return &comment, nil
	}

	var parent bbs.Comment
	if err := c.db.First(&parent, parentID).Error; err != nil {
		return nil, err
	}
	if err := c.db.Model(&parent).Association("Subs").Append(&comment).Error; err != nil {
		return nil, err
	}

	return &comment, nil

}

func (c *commentOperation) remove(topicID uint, id uint) (err error) {
	//if err := data.Db.Where("user_id = ? AND id = ?", userID, topicID).Delete(&bbs.Topic{}).Error; err != nil {
	//	return err
	//}

	return nil
}

func (c *commentOperation) update(comment *bbs.Comment) error {
	return c.db.Save(comment).Error
}
