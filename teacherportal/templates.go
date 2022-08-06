package teacherportal

import (
	"html/template"
	"path/filepath"
)

var rootTemplate *template.Template

func ImportTemplates() error {
	
	// address from cmd/teacherportal
	filePrefix, _ := filepath.Abs("../../teacherportal")
	
	var err error
	rootTemplate, err = template.ParseFiles(
		filePrefix + "/students.gohtml",
		filePrefix + "/student.gohtml")

	if err != nil {
		return err
	}

	return nil
}
