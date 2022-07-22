package db

import (
	"errors"
	"github.com/glebarez/sqlite"
	"github.com/ngoldack/tt/internal/model"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

type Manager interface {
	CreateFrame(frame *model.Frame) error
	UpdateFrame(frame *model.Frame) error
	DeleteFrame(frame *model.Frame) error
	FindFramesByProject(project *model.Project) (frames []model.Frame, err error)
	FindFrameByActive(active bool) (frame *model.Frame, err error)

	CreateProject(project *model.Project) error
	UpdateProject(project *model.Project) error
	DeleteProject(project *model.Project) error
	FindProjects() (projects []model.Project, err error)
	FindProjectByName(name string) (project *model.Project, err error)
	FindProjectsByActive(active bool) (projects []model.Project, err error)

	CreateTag(tag *model.Tag) error
	UpdateTag(tag *model.Tag) error
	DeleteTag(tag *model.Tag) error
	FindTags() (tags []model.Tag, err error)
	FindTagByName(name string) (tag *model.Tag, err error)
	FindTagsByProject(project *model.Project) (tags []model.Tag, err error)
}

type manager struct {
	db *gorm.DB
}

var Mgr Manager

func init() {
	newLogger := logger.New(log.New(os.Stdout, "gorm", 0), logger.Config{
		Colorful:                  false,
		IgnoreRecordNotFoundError: true,
		LogLevel:                  logger.Silent,
	})

	db, err := gorm.Open(sqlite.Open("tt.db"), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		Logger:                 newLogger,
	})
	if err != nil {
		log.Fatal("Failed to init db:", err)
	}

	if err = db.AutoMigrate(&model.Project{}); err != nil {
		log.Fatal("Error while migrating:", err)
	}
	if err = db.AutoMigrate(&model.Tag{}); err != nil {
		log.Fatal("Error while migrating:", err)
	}
	if err = db.AutoMigrate(&model.Frame{}); err != nil {
		log.Fatal("Error while migrating:", err)
	}

	Mgr = &manager{db: db}
}

func (m manager) CreateFrame(frame *model.Frame) error {
	return m.db.Create(frame).Error
}

func (m manager) UpdateFrame(frame *model.Frame) error {
	return m.db.Save(frame).Error
}

func (m manager) DeleteFrame(frame *model.Frame) error {
	return m.db.Delete(frame).Error
}

func (m manager) FindFramesByProject(project *model.Project) (frames []model.Frame, err error) {
	err = m.db.Where("project_id = ?", project.ID).Find(&frames).Error
	return
}

func (m manager) FindFrameByActive(active bool) (*model.Frame, error) {
	var frame *model.Frame
	var err error
	if active {
		err = m.db.Where("active = ?", 1).Preload("Project").Preload("Tag").First(&frame).Error
	} else {
		err = m.db.Where("active = ?", 0).Preload("Project").Preload("Tag").First(&frame).Error
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		frame = nil
		err = nil
	}
	return frame, err
}

func (m manager) CreateProject(project *model.Project) error {
	return m.db.Create(project).Error
}

func (m manager) UpdateProject(project *model.Project) error {
	return m.db.Save(project).Error
}

func (m manager) DeleteProject(project *model.Project) error {
	return m.db.Delete(project).Error
}

func (m manager) FindProjects() (projects []model.Project, err error) {
	err = m.db.Find(&projects).Error
	return
}

func (m manager) FindProjectByName(name string) (project *model.Project, err error) {
	err = m.db.Where("name = ?", name).Limit(1).Find(&project).Error
	return
}

func (m manager) FindProjectsByActive(active bool) (projects []model.Project, err error) {
	if active {
		err = m.db.Where("active = ?", 1).Find(&projects).Error
		return
	}
	err = m.db.Where("active = ?", 0).Find(&projects).Error
	return
}

func (m manager) CreateTag(tag *model.Tag) error {
	return m.db.Create(tag).Error
}

func (m manager) UpdateTag(tag *model.Tag) error {
	return m.db.Save(tag).Error
}

func (m manager) DeleteTag(tag *model.Tag) error {
	return m.db.Delete(tag).Error
}

func (m manager) FindTags() (tags []model.Tag, err error) {
	err = m.db.Find(&tags).Error
	return
}

func (m manager) FindTagByName(name string) (tag *model.Tag, err error) {
	err = m.db.Where("name = ?", name).First(&tag).Error
	return
}

func (m manager) FindTagsByProject(project *model.Project) (tags []model.Tag, err error) {
	err = m.db.Joins("JOIN frames ON frames.tag_id = tags.tag_id").Where("frames.project_id = ?", project.ID).Find(tags).Error
	return
}
