package library;

import (
	"encoding/xml"
	"net/http"
	"strings"
)

var languages = map[string]map[string]string {
	"en" : {
		"en" : "English",
		"fr" : "Anglais"
	},
	"fr" : {
		"en" : "French",
		"fr" : "Français"
	}
}

type AuthorName struct {

}

type AuthorNames struct {

}

type Author struct {

}

type Edition struct {

}

type Editions struct {

}

type WorkName struct {

}

type WorkNames struct {

}

type Work struct {

}

///Name of the library
var libraryName = map[string]string {
	"en" : "Electronic Library",
	"fr" : "Bibliothèque Électronique",
	"la" : "Bibliotheca Electronica"
}

///Name of the table of contents
var tocName = map[string]string {
	"en" : "Table of Contents",
	"fr" : "Table des matières",
	"la" : "Continentur in hoc"
}

//Name for the previous page/section
var previousName = map[string]string {
	"en" : "Previous",
	"fr" : "Précédente",
	"la" : ""
}

//Name for the next page/section
var nextName = map[string]string {
	"en" : "Next",
	"fr" : "Suivante",
	"la" : "",
}

/// /bibliotheque/descartes/meditations?view=page?edition=gonthier?section=VI, 2

func (lib *Library) WriteLibraryPageHeader(w http.ResponseWriter, language string) {

}

func (lib *Library) WriteLibraryPageFooter(w http.ResponseWriter, language string) {

}

///Library manager.
type Library struct {
	///Directory where the library texts are located.
	libraryWorkDirectory string
	///Directory where the library database xml files are located.
	libraryDatabaseDirectory string

	///List of custom pages
	customPages map[string]CustomPage
	///List of works without authors
	unauthoredWorks map[string]Work
	///List of works which have an author
	authoredWorks map[string]map[string]Work
}

///Returns if the library has been loaded
func (lib *Library) Loaded() bool {
	return (customPages!=nil && unauthoredWorks != nil && authoredWorks !=nil)
}

///Load from database (or reloads).
func (lib *Library) Load() error {

}

///Execute the page for an author.
///
///        interfaceLanguage:
///        author:
///        viewBy: either 'language' or 'work' (language will list the interface language first)
///
func (lib *Library) ExecuteAuthorPage(w http.ResponseWriter,  interfaceLanguage string, author string, viewBy string) {

}

///Execute the page for a work.
///
///        interfaceLanguage: the language the interface must be displayed in
///        author: the url name of the author
///        title: the url title of the work
///        view: either 'complete' or the name of a particular sectionning scheme.
///        edition: the edition of the work, or nil for the default
///        language: the language to print the work in (will be ignored if it contradicts the edition)
///        sectionFrom: the section requested for the work
///        sectionTo ...
///
func (lib *Library) ExecuteWorkPage(w http.ResponseWriter, interfaceLanguage string, author string, title string, view string, edition string, language string, sectionFrom string, sectionTo string) {

}

///Execute on a custom page.
func (lib *Library) ExecuteCustomPage(w http.ResponseWriter, interfaceLanguage string, page string) {

}

///Execute from the url query
func (lib *Library) Execute(w http.ResponseWriter, moniker string, path string, map[string]string parameters) error {
	//Load if not loaded
	if !lib.Loaded() {
		err := lib.Load()
		if err != nil {
			return err
		}
	}
}
