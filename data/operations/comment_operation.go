package operations

import (
	"github.com/346285234/bbs-server/data"
	"github.com/346285234/bbs-server/data/models"
)

type commentOperation struct {
}

var CoO = commentOperation{}

func (_ *commentOperation) List(topicID uint) (comments []*models.Comment, err error) {
	if err := data.Db.Preload("Subs").Find(&comments).Error; err != nil {
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

func (_ *commentOperation) Add(comment models.Comment, parentID uint) (*models.Comment, error) {
	if err := data.Db.Create(&comment).Error; err != nil {
		return nil, err
	}

	if parentID == 0 {
		return &comment, nil
	}

	var parent models.Comment
	if err := data.Db.First(&parent, parentID).Error; err != nil {
		return nil, err
	}
	if err := data.Db.Model(&parent).Association("Subs").Append(&comment).Error; err != nil {
		return nil, err
	}

	return &comment, nil

}

func (_ *commentOperation) Remove(topicID uint, id uint) (err error) {
	//if err := data.Db.Where("user_id = ? AND id = ?", userID, topicID).Delete(&models.Topic{}).Error; err != nil {
	//	return err
	//}

	return nil
}
