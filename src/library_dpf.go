/// DISPLAY-PROXIMATE FORMAT
///
/// This contains utilities to parse the display-proximate format of
/// the electronic library system.

package library;

import (
	"bufio"
)

type SectioningScheme struct {
	main bool
	toplevelSection SectionDefinition
}

type SectionDefinition struct {
	name string
	break bool
	subsections []SectionDefinition
}

type Section struct {
	level string
	name string
	path string
	subsections []Section
}

/// Document definition structure.
type DocumentDefinition struct {
	languageCode string
	authorCode string
	titleCode string
	authors []string
	title string
	subtitle string
	translators []string
	year int
	publisher string
	///Sectioning schemes
	schemes []SectioningScheme
}

/// Parse the document definition, stops at the text body.
func readDocumentDefinition (reader bufio.Reader) (DocumentDefinition dd, error err) {

}

/// Skip until the body of the document (skips the document definition header).
func skipToDocumentBody (reader bufio.Reader) error {

}

/// Processes the document body and returns a table of contents according to the main division
func readDocumentTOC (reader bufio.Reader, documentDefinition DocumentDefinition) error {

}

/// Process the document body and write an html format
/// the stream must already be at the body
func processDocumentBody(reader bufio.Reader, fromSection string, toSection string, writer bufio.Writer) error {

}
