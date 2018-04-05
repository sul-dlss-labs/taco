package datautils

import (
	"fmt"
)

// FileType is the type URI for files
const FileType = "http://sdr.sul.stanford.edu/models/sdr3-file.jsonld"

// FilesetType is the type URI for filesets
const FilesetType = "http://sdr.sul.stanford.edu/models/sdr3-fileset.jsonld"

// ObjectTypes is the list of object subtype URIs
var ObjectTypes = []string{
	"http://sdr.sul.stanford.edu/models/sdr3-object.jsonld",
	"http://sdr.sul.stanford.edu/models/sdr3-3d.jsonld",
	"http://sdr.sul.stanford.edu/models/sdr3-agreement.jsonld",
	"http://sdr.sul.stanford.edu/models/sdr3-book.jsonld",
	"http://sdr.sul.stanford.edu/models/sdr3-document.jsonld",
	"http://sdr.sul.stanford.edu/models/sdr3-geo.jsonld",
	"http://sdr.sul.stanford.edu/models/sdr3-image.jsonld",
	"http://sdr.sul.stanford.edu/models/sdr3-page.jsonld",
	"http://sdr.sul.stanford.edu/models/sdr3-photograph.jsonld",
	"http://sdr.sul.stanford.edu/models/sdr3-manuscript.jsonld",
	"http://sdr.sul.stanford.edu/models/sdr3-map.jsonld",
	"http://sdr.sul.stanford.edu/models/sdr3-media.jsonld",
	"http://sdr.sul.stanford.edu/models/sdr3-track.jsonld",
	"http://sdr.sul.stanford.edu/models/sdr3-webarchive-binary.jsonld",
	"http://sdr.sul.stanford.edu/models/sdr3-webarchive-seed.jsonld",
}

// CollectionTypes is the list of object subtype URIs
var CollectionTypes = []string{
	"http://sdr.sul.stanford.edu/models/sdr3-collection.jsonld",
	"http://sdr.sul.stanford.edu/models/sdr3-curated-collection.jsonld",
	"http://sdr.sul.stanford.edu/models/sdr3-user-collection.jsonld",
	"http://sdr.sul.stanford.edu/models/sdr3-exhibit.jsonld",
	"http://sdr.sul.stanford.edu/models/sdr3-series.jsonld",
}

// Resource represents the resource as it exists in the persistence layer
// this is very similar to models.Resource, but COULD vary, so we should
// keep them separated
type Resource struct {
	JSON JSONObject
}

// NewResource creates a new resource instance
func NewResource(data map[string]interface{}) *Resource {
	return &Resource{JSON: JSONObject(data)}
}

// ID returns the document's identifier
func (d *Resource) ID() string {
	return d.JSON.GetS("id")
}

// Version returns the document's version
func (d *Resource) Version() int {
	return d.JSON.GetI("version")
}

// Type returns the document's type
func (d *Resource) Type() string {
	return d.JSON.GetS("@type")
}

// IsFile returns true if the resource has the file type assertion
func (d *Resource) IsFile() bool {
	return d.Type() == FileType
}

// IsFileset returns true if the resource has the fileset type assertion
func (d *Resource) IsFileset() bool {
	return d.Type() == FilesetType
}

// IsObject returns true if the resource has an object type assertion
func (d *Resource) IsObject() bool {
	return contains(ObjectTypes, d.Type())
}

// IsCollection returns true if the resource has an object type assertion
func (d *Resource) IsCollection() bool {
	return contains(CollectionTypes, d.Type())
}

// CurrentVersion returns true if this is the current version
func (d *Resource) CurrentVersion() bool {
	return d.JSON.GetB("currentVersion")
}

// WithID sets the document's identifier
func (d *Resource) WithID(id string) *Resource {
	d.JSON["id"] = id
	return d
}

// WithCurrentVersion sets the currentVersion flag
func (d *Resource) WithCurrentVersion(flag bool) *Resource {
	d.JSON["currentVersion"] = flag
	return d
}

// WithVersion sets the version
func (d *Resource) WithVersion(version int) *Resource {
	d.JSON["version"] = version
	return d
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// Structural returns structural.isContainedBy
func (d *Resource) Structural() *JSONObject {
	return d.JSON.GetObj("structural")
}

// Identification returns the identification subschema
func (d *Resource) Identification() *JSONObject {
	return d.JSON.GetObj("identification")
}

func (d *Resource) String() string {
	return fmt.Sprintf("<Resource id: '%s'>", d.ID())
}
