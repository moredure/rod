// This file is generated by "./lib/proto/generate"

package proto

/*

IndexedDB

*/

// IndexedDBDatabaseWithObjectStores Database with an array of object stores.
type IndexedDBDatabaseWithObjectStores struct {

	// Name Database name.
	Name string `json:"name"`

	// Version Database version (type is not 'integer', as the standard
	// requires the version number to be 'unsigned long long')
	Version float64 `json:"version"`

	// ObjectStores Object stores in this database.
	ObjectStores []*IndexedDBObjectStore `json:"objectStores"`
}

// IndexedDBObjectStore Object store.
type IndexedDBObjectStore struct {

	// Name Object store name.
	Name string `json:"name"`

	// KeyPath Object store key path.
	KeyPath *IndexedDBKeyPath `json:"keyPath"`

	// AutoIncrement If true, object store has auto increment flag set.
	AutoIncrement bool `json:"autoIncrement"`

	// Indexes Indexes in this object store.
	Indexes []*IndexedDBObjectStoreIndex `json:"indexes"`
}

// IndexedDBObjectStoreIndex Object store index.
type IndexedDBObjectStoreIndex struct {

	// Name Index name.
	Name string `json:"name"`

	// KeyPath Index key path.
	KeyPath *IndexedDBKeyPath `json:"keyPath"`

	// Unique If true, index is unique.
	Unique bool `json:"unique"`

	// MultiEntry If true, index allows multiple entries for a key.
	MultiEntry bool `json:"multiEntry"`
}

// IndexedDBKeyType enum
type IndexedDBKeyType string

const (
	// IndexedDBKeyTypeNumber enum const
	IndexedDBKeyTypeNumber IndexedDBKeyType = "number"

	// IndexedDBKeyTypeString enum const
	IndexedDBKeyTypeString IndexedDBKeyType = "string"

	// IndexedDBKeyTypeDate enum const
	IndexedDBKeyTypeDate IndexedDBKeyType = "date"

	// IndexedDBKeyTypeArray enum const
	IndexedDBKeyTypeArray IndexedDBKeyType = "array"
)

// IndexedDBKey Key.
type IndexedDBKey struct {

	// Type Key type.
	Type IndexedDBKeyType `json:"type"`

	// Number (optional) Number value.
	Number float64 `json:"number,omitempty"`

	// String (optional) String value.
	String string `json:"string,omitempty"`

	// Date (optional) Date value.
	Date float64 `json:"date,omitempty"`

	// Array (optional) Array value.
	Array []*IndexedDBKey `json:"array,omitempty"`
}

// IndexedDBKeyRange Key range.
type IndexedDBKeyRange struct {

	// Lower (optional) Lower bound.
	Lower *IndexedDBKey `json:"lower,omitempty"`

	// Upper (optional) Upper bound.
	Upper *IndexedDBKey `json:"upper,omitempty"`

	// LowerOpen If true lower bound is open.
	LowerOpen bool `json:"lowerOpen"`

	// UpperOpen If true upper bound is open.
	UpperOpen bool `json:"upperOpen"`
}

// IndexedDBDataEntry Data entry.
type IndexedDBDataEntry struct {

	// Key Key object.
	Key *RuntimeRemoteObject `json:"key"`

	// PrimaryKey Primary key object.
	PrimaryKey *RuntimeRemoteObject `json:"primaryKey"`

	// Value Value object.
	Value *RuntimeRemoteObject `json:"value"`
}

// IndexedDBKeyPathType enum
type IndexedDBKeyPathType string

const (
	// IndexedDBKeyPathTypeNull enum const
	IndexedDBKeyPathTypeNull IndexedDBKeyPathType = "null"

	// IndexedDBKeyPathTypeString enum const
	IndexedDBKeyPathTypeString IndexedDBKeyPathType = "string"

	// IndexedDBKeyPathTypeArray enum const
	IndexedDBKeyPathTypeArray IndexedDBKeyPathType = "array"
)

// IndexedDBKeyPath Key path.
type IndexedDBKeyPath struct {

	// Type Key path type.
	Type IndexedDBKeyPathType `json:"type"`

	// String (optional) String value.
	String string `json:"string,omitempty"`

	// Array (optional) Array value.
	Array []string `json:"array,omitempty"`
}

// IndexedDBClearObjectStore Clears all entries from an object store.
type IndexedDBClearObjectStore struct {

	// SecurityOrigin Security origin.
	SecurityOrigin string `json:"securityOrigin"`

	// DatabaseName Database name.
	DatabaseName string `json:"databaseName"`

	// ObjectStoreName Object store name.
	ObjectStoreName string `json:"objectStoreName"`
}

// ProtoReq of the command
func (m IndexedDBClearObjectStore) ProtoReq() string { return "IndexedDB.clearObjectStore" }

// Call of the command, sessionID is optional.
func (m IndexedDBClearObjectStore) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// IndexedDBDeleteDatabase Deletes a database.
type IndexedDBDeleteDatabase struct {

	// SecurityOrigin Security origin.
	SecurityOrigin string `json:"securityOrigin"`

	// DatabaseName Database name.
	DatabaseName string `json:"databaseName"`
}

// ProtoReq of the command
func (m IndexedDBDeleteDatabase) ProtoReq() string { return "IndexedDB.deleteDatabase" }

// Call of the command, sessionID is optional.
func (m IndexedDBDeleteDatabase) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// IndexedDBDeleteObjectStoreEntries Delete a range of entries from an object store
type IndexedDBDeleteObjectStoreEntries struct {

	// SecurityOrigin ...
	SecurityOrigin string `json:"securityOrigin"`

	// DatabaseName ...
	DatabaseName string `json:"databaseName"`

	// ObjectStoreName ...
	ObjectStoreName string `json:"objectStoreName"`

	// KeyRange Range of entry keys to delete
	KeyRange *IndexedDBKeyRange `json:"keyRange"`
}

// ProtoReq of the command
func (m IndexedDBDeleteObjectStoreEntries) ProtoReq() string {
	return "IndexedDB.deleteObjectStoreEntries"
}

// Call of the command, sessionID is optional.
func (m IndexedDBDeleteObjectStoreEntries) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// IndexedDBDisable Disables events from backend.
type IndexedDBDisable struct {
}

// ProtoReq of the command
func (m IndexedDBDisable) ProtoReq() string { return "IndexedDB.disable" }

// Call of the command, sessionID is optional.
func (m IndexedDBDisable) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// IndexedDBEnable Enables events from backend.
type IndexedDBEnable struct {
}

// ProtoReq of the command
func (m IndexedDBEnable) ProtoReq() string { return "IndexedDB.enable" }

// Call of the command, sessionID is optional.
func (m IndexedDBEnable) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// IndexedDBRequestData Requests data from object store or index.
type IndexedDBRequestData struct {

	// SecurityOrigin Security origin.
	SecurityOrigin string `json:"securityOrigin"`

	// DatabaseName Database name.
	DatabaseName string `json:"databaseName"`

	// ObjectStoreName Object store name.
	ObjectStoreName string `json:"objectStoreName"`

	// IndexName Index name, empty string for object store data requests.
	IndexName string `json:"indexName"`

	// SkipCount Number of records to skip.
	SkipCount int `json:"skipCount"`

	// PageSize Number of records to fetch.
	PageSize int `json:"pageSize"`

	// KeyRange (optional) Key range.
	KeyRange *IndexedDBKeyRange `json:"keyRange,omitempty"`
}

// ProtoReq of the command
func (m IndexedDBRequestData) ProtoReq() string { return "IndexedDB.requestData" }

// Call of the command, sessionID is optional.
func (m IndexedDBRequestData) Call(c Client) (*IndexedDBRequestDataResult, error) {
	var res IndexedDBRequestDataResult
	return &res, call(m.ProtoReq(), m, &res, c)
}

// IndexedDBRequestDataResult Requests data from object store or index.
type IndexedDBRequestDataResult struct {

	// ObjectStoreDataEntries Array of object store data entries.
	ObjectStoreDataEntries []*IndexedDBDataEntry `json:"objectStoreDataEntries"`

	// HasMore If true, there are more entries to fetch in the given range.
	HasMore bool `json:"hasMore"`
}

// IndexedDBGetMetadata Gets metadata of an object store
type IndexedDBGetMetadata struct {

	// SecurityOrigin Security origin.
	SecurityOrigin string `json:"securityOrigin"`

	// DatabaseName Database name.
	DatabaseName string `json:"databaseName"`

	// ObjectStoreName Object store name.
	ObjectStoreName string `json:"objectStoreName"`
}

// ProtoReq of the command
func (m IndexedDBGetMetadata) ProtoReq() string { return "IndexedDB.getMetadata" }

// Call of the command, sessionID is optional.
func (m IndexedDBGetMetadata) Call(c Client) (*IndexedDBGetMetadataResult, error) {
	var res IndexedDBGetMetadataResult
	return &res, call(m.ProtoReq(), m, &res, c)
}

// IndexedDBGetMetadataResult Gets metadata of an object store
type IndexedDBGetMetadataResult struct {

	// EntriesCount the entries count
	EntriesCount float64 `json:"entriesCount"`

	// KeyGeneratorValue the current value of key generator, to become the next inserted
	// key into the object store. Valid if objectStore.autoIncrement
	// is true.
	KeyGeneratorValue float64 `json:"keyGeneratorValue"`
}

// IndexedDBRequestDatabase Requests database with given name in given frame.
type IndexedDBRequestDatabase struct {

	// SecurityOrigin Security origin.
	SecurityOrigin string `json:"securityOrigin"`

	// DatabaseName Database name.
	DatabaseName string `json:"databaseName"`
}

// ProtoReq of the command
func (m IndexedDBRequestDatabase) ProtoReq() string { return "IndexedDB.requestDatabase" }

// Call of the command, sessionID is optional.
func (m IndexedDBRequestDatabase) Call(c Client) (*IndexedDBRequestDatabaseResult, error) {
	var res IndexedDBRequestDatabaseResult
	return &res, call(m.ProtoReq(), m, &res, c)
}

// IndexedDBRequestDatabaseResult Requests database with given name in given frame.
type IndexedDBRequestDatabaseResult struct {

	// DatabaseWithObjectStores Database with an array of object stores.
	DatabaseWithObjectStores *IndexedDBDatabaseWithObjectStores `json:"databaseWithObjectStores"`
}

// IndexedDBRequestDatabaseNames Requests database names for given security origin.
type IndexedDBRequestDatabaseNames struct {

	// SecurityOrigin Security origin.
	SecurityOrigin string `json:"securityOrigin"`
}

// ProtoReq of the command
func (m IndexedDBRequestDatabaseNames) ProtoReq() string { return "IndexedDB.requestDatabaseNames" }

// Call of the command, sessionID is optional.
func (m IndexedDBRequestDatabaseNames) Call(c Client) (*IndexedDBRequestDatabaseNamesResult, error) {
	var res IndexedDBRequestDatabaseNamesResult
	return &res, call(m.ProtoReq(), m, &res, c)
}

// IndexedDBRequestDatabaseNamesResult Requests database names for given security origin.
type IndexedDBRequestDatabaseNamesResult struct {

	// DatabaseNames Database names for origin.
	DatabaseNames []string `json:"databaseNames"`
}