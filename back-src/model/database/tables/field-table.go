package tables

import (
	"back-src/model/existence"
	"github.com/go-pg/pg"
)

type FieldTable struct {
	conn *pg.DB
}

func NewFieldsTable(db *pg.DB) *FieldTable {
	return &FieldTable{db}
}

func (table *FieldTable) GetFieldSkills(fieldId string) (skills []string, error error) {
	var field existence.Field
	error = table.conn.Model(&field).Column("skills").Where("id = ?", fieldId).Select()
	skills = field.Skills
	return
}

func (table *FieldTable) AddSkillToField(fieldId string, skill string) error {
	var skills []string
	field := existence.Field{Id: fieldId, Skills: skills}
	if err := table.conn.Model(&field).Column("skills").Where("id = ?", field.Id).Select(); err != nil {
		return err
	}
	field.Skills = append(field.Skills, skill)
	if _, err := table.conn.Model(&field).Column("skills").Where("id = ?", fieldId).Update(); err != nil {
		return err
	}
	return nil
}

func (table *FieldTable) IsThereFieldWithId(id string) (bool, error) {
	var fields []existence.Field
	err := table.conn.Model(&fields).Where("id = ?", id).Select()
	if err != nil {
		return false, err
	}
	return len(fields) != 0, nil
}

func (table *FieldTable) GetAllFieldsWithSkills() ([]existence.Field, error) {
	var fields []existence.Field
	err := table.conn.Model(&fields).Select()
	if err != nil {
		return []existence.Field{}, err
	}
	return fields, nil
}

func (table *FieldTable) GetField(fieldId string) (existence.Field, error) {
	field := existence.Field{}
	if err := table.conn.Model(&field).Where("id = ?", fieldId).Select(); err != nil {
		return existence.Field{}, err
	}
	return field, nil
}
