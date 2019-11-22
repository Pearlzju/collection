package collection

import (
	"encoding/json"
	"fmt"

	"github.com/shopspring/decimal"
)

type BaseCollection struct {
	value  interface{}
	length int
	err    error
}

const (
	ErrNotImplement = "not implement %s"
)

func (c *BaseCollection) errorHandle(e string, args ...string) {
	if c.err == nil {
		c.err = fmt.Errorf(e, args)
	}
}

func (c BaseCollection) Value() interface{} {
	return c.value
}

// Length return the length of the collection.
func (c BaseCollection) Length() int {
	return c.length
}

// Select select the keys of collection and delete others.
func (c BaseCollection) Select(keys ...string) Collection {
	c.errorHandle(ErrNotImplement, "Select")
	return c
}

// Column select the values of collection by the given key.
func (c BaseCollection) Column(key string) Collection {
	c.errorHandle(ErrNotImplement, "Column")
	return c
}

// ToStruct turn the collection to the specified struct using mapstructure.
// https://github.com/mitchellh/mapstructure
func (c BaseCollection) ToStruct(dist interface{}) {
}

func (c BaseCollection) ToStructE(dist interface{}) error {
	c.errorHandle(ErrNotImplement, "ToStructE")
	return c.err
}

// All returns the underlying array represented by the collection.
func (c BaseCollection) All() []interface{} {
	c.errorHandle(ErrNotImplement, "All")
	return nil
}

func (c BaseCollection) AllE() ([]interface{}, error) {
	c.errorHandle(ErrNotImplement, "AllE")
	return nil, c.err
}

// Avg returns the average value of a given key.
func (c BaseCollection) Avg(key ...string) decimal.Decimal {
	return c.Sum(key...).Div(decimal.New(int64(c.length), 0))
}

// Sum returns the sum of all items in the collection.
func (c BaseCollection) Sum(key ...string) decimal.Decimal {
	c.errorHandle(ErrNotImplement, "Sum")
	return decimal.Decimal{}
}

// Min returns the minimum value of a given key.
func (c BaseCollection) Min(key ...string) decimal.Decimal {
	c.errorHandle(ErrNotImplement, "Min")
	return decimal.Decimal{}
}

// Max returns the maximum value of a given key.
func (c BaseCollection) Max(key ...string) decimal.Decimal {
	c.errorHandle(ErrNotImplement, "Max")
	return decimal.Decimal{}
}

// Join joins the collection's values with a string.
func (c BaseCollection) Join(delimiter string) string {
	return ""
}

func (c BaseCollection) JoinE(delimiter string) (string, error) {
	c.errorHandle(ErrNotImplement, "Join")
	return "", c.err
}

// Combine combines the values of the collection, as keys, with the values of another array or collection.
func (c BaseCollection) Combine(value []interface{}) Collection {
	c.errorHandle(ErrNotImplement, "Combine")
	return c
}

// Pluck retrieves all of the values for a given key.
func (c BaseCollection) Pluck(key string) Collection {
	c.errorHandle(ErrNotImplement, "Pluck")
	return c
}

// ToIntArray converts the collection into a plain golang slice which contains int.
func (c BaseCollection) ToIntArray() []int {
	return nil
}

func (c BaseCollection) ToIntArrayE() ([]int, error) {
	c.errorHandle(ErrNotImplement, "ToIntArrayE")
	return nil, c.err
}

// ToInt64Array converts the collection into a plain golang slice which contains int64.
func (c BaseCollection) ToInt64Array() []int64 {
	return nil
}

func (c BaseCollection) ToInt64ArrayE() ([]int64, error) {
	c.errorHandle(ErrNotImplement, "ToInt64ArrayE")
	return nil, c.err
}

// Mode returns the mode value of a given key.
func (c BaseCollection) Mode(key ...string) []interface{} {
	return nil
}

func (c BaseCollection) ModeE(key ...string) ([]interface{}, error) {
	c.errorHandle(ErrNotImplement, "ModeE")
	return nil, c.err
}

// Only returns the items in the collection with the specified keys.
func (c BaseCollection) Only(keys []string) Collection {
	c.errorHandle(ErrNotImplement, "Only")
	return c
}

// Prepend adds an item to the beginning of the collection.
func (c BaseCollection) Prepend(values ...interface{}) Collection {
	c.errorHandle(ErrNotImplement, "Prepend")
	return c
}

// Pull removes and returns an item from the collection by its key.
func (c BaseCollection) Pull(key interface{}) Collection {
	c.errorHandle(ErrNotImplement, "Pull")
	return c
}

// Put sets the given key and value in the collection:.
func (c BaseCollection) Put(key string, value interface{}) Collection {
	c.errorHandle(ErrNotImplement, "Put")
	return c
}

// SortBy sorts the collection by the given key.
func (c BaseCollection) SortBy(key string) Collection {
	c.errorHandle(ErrNotImplement, "SortBy")
	return c
}

// Take returns a new collection with the specified number of items.
func (c BaseCollection) Take(num int) Collection {
	c.errorHandle(ErrNotImplement, "Take")
	return c
}

// Chunk breaks the collection into multiple, smaller collections of a given size.
func (c BaseCollection) Chunk(num int) MultiDimensionalArrayCollection {
	c.errorHandle(ErrNotImplement, "Chunk")
	return MultiDimensionalArrayCollection{}
}

// Collapse collapses a collection of arrays into a single, flat collection.
func (c BaseCollection) Collapse() Collection {
	c.errorHandle(ErrNotImplement, "Collapse")
	return c
}

// Concat appends the given array or collection values onto the end of the collection.
func (c BaseCollection) Concat(value interface{}) Collection {
	c.errorHandle(ErrNotImplement, "Concat")
	return c
}

// Contains determines whether the collection contains a given item.
func (c BaseCollection) Contains(value ...interface{}) bool {
	return false
}

func (c BaseCollection) ContainsE(value ...interface{}) (bool, error) {
	c.errorHandle(ErrNotImplement, "ContainsE")
	return false, c.err
}

// CountBy counts the occurrences of values in the collection. By default, the method counts the occurrences of every element.
func (c BaseCollection) CountBy(callback ...interface{}) map[interface{}]int {
	return nil
}

func (c BaseCollection) CountByE(callback ...interface{}) (map[interface{}]int, error) {
	c.errorHandle(ErrNotImplement, "CountByE")
	return nil, c.err
}

// CrossJoin cross joins the collection's values among the given arrays or collections, returning a Cartesian product with all possible permutations.
func (c BaseCollection) CrossJoin(array ...[]interface{}) MultiDimensionalArrayCollection {
	c.errorHandle(ErrNotImplement, "CrossJoin")
	return MultiDimensionalArrayCollection{}
}

// Dd dumps the collection's items and ends execution of the script.
func (c BaseCollection) Dd() {
}

func (c BaseCollection) DdE() error {
	c.errorHandle(ErrNotImplement, "DdE")
	return c.err
}

// Diff compares the collection against another collection or a plain PHP array based on its values.
// This method will return the values in the original collection that are not present in the given collection.
func (c BaseCollection) Diff(interface{}) Collection {
	c.errorHandle(ErrNotImplement, "Diff")
	return c
}

// DiffAssoc compares the collection against another collection or a plain PHP  array based on its keys and values.
// This method will return the key / value pairs in the original collection that are not present in the given collection.
func (c BaseCollection) DiffAssoc(map[string]interface{}) Collection {
	c.errorHandle(ErrNotImplement, "DiffAssoc")
	return c
}

// DiffKeys compares the collection against another collection or a plain PHP array based on its keys.
// This method will return the key / value pairs in the original collection that are not present in the given collection.
func (c BaseCollection) DiffKeys(map[string]interface{}) Collection {
	c.errorHandle(ErrNotImplement, "DiffKeys")
	return c
}

// Dump dumps the collection's items.
func (c BaseCollection) Dump() {
}

func (c BaseCollection) DumpE() error {
	c.errorHandle(ErrNotImplement, "DumpE")
	return c.err
}

// Each iterates over the items in the collection and passes each item to a callback.
func (c BaseCollection) Each(func(item, value interface{}) (interface{}, bool)) Collection {
	c.errorHandle(ErrNotImplement, "Each")
	return c
}

// Every may be used to verify that all elements of a collection pass a given truth test.
func (c BaseCollection) Every(CB) bool {
	return false
}

func (c BaseCollection) EveryE(CB) (bool, error) {
	c.errorHandle(ErrNotImplement, "EveryE")
	return false, c.err
}

// Except returns all items in the collection except for those with the specified keys.
func (c BaseCollection) Except([]string) Collection {
	c.errorHandle(ErrNotImplement, "Except")
	return c
}

// Filter filters the collection using the given callback, keeping only those items that pass a given truth test.
func (c BaseCollection) Filter(CB) Collection {
	c.errorHandle(ErrNotImplement, "Filter")
	return c
}

// First returns the first element in the collection that passes a given truth test.
func (c BaseCollection) First(...CB) interface{} {
	return nil
}

func (c BaseCollection) FirstE(...CB) (interface{}, error) {
	c.errorHandle(ErrNotImplement, "FirstE")
	return nil, c.err
}

// FirstWhere returns the first element in the collection with the given key / value pair.
func (c BaseCollection) FirstWhere(key string, values ...interface{}) map[string]interface{} {
	return nil
}

func (c BaseCollection) FirstWhereE(key string, values ...interface{}) (map[string]interface{}, error) {
	c.errorHandle(ErrNotImplement, "FirstWhereE")
	return nil, c.err
}

// FlatMap iterates through the collection and passes each value to the given callback.
func (c BaseCollection) FlatMap(func(value interface{}) interface{}) Collection {
	c.errorHandle(ErrNotImplement, "FlatMap")
	return c
}

// Flip swaps the collection's keys with their corresponding values.
func (c BaseCollection) Flip() Collection {
	c.errorHandle(ErrNotImplement, "Flip")
	return c
}

// Forget removes an item from the collection by its key.
func (c BaseCollection) Forget(string) Collection {
	c.errorHandle(ErrNotImplement, "Forget")
	return c
}

// ForPage returns a new collection containing the items that would be present on a given page number.
func (c BaseCollection) ForPage(int, int) Collection {
	c.errorHandle(ErrNotImplement, "ForPage")
	return c
}

// Get returns the item at a given key. If the key does not exist, null is returned.
func (c BaseCollection) Get(string, ...interface{}) interface{} {
	return nil
}

func (c BaseCollection) GetE(string, ...interface{}) (interface{}, error) {
	c.errorHandle(ErrNotImplement, "GetE")
	return nil, c.err
}

// GroupBy groups the collection's items by a given key.
func (c BaseCollection) GroupBy(string) Collection {
	c.errorHandle(ErrNotImplement, "GroupBy")
	return c
}

// Has determines if a given key exists in the collection.
func (c BaseCollection) Has(...string) bool {
	return false
}

func (c BaseCollection) HasE(...string) (bool, error) {
	c.errorHandle(ErrNotImplement, "HasE")
	return false, c.err
}

// Implode joins the items in a collection. Its arguments depend on the type of items in the collection.
func (c BaseCollection) Implode(string, string) string {
	return ""
}

func (c BaseCollection) ImplodeE(string, string) (string, error) {
	c.errorHandle(ErrNotImplement, "ImplodeE")
	return "", c.err
}

// Intersect removes any values from the original collection that are not present in the given array or collection.
func (c BaseCollection) Intersect([]string) Collection {
	c.errorHandle(ErrNotImplement, "Intersect")
	return c
}

// IntersectByKeys removes any keys from the original collection that are not present in the given array or collection.
func (c BaseCollection) IntersectByKeys(map[string]interface{}) Collection {
	c.errorHandle(ErrNotImplement, "IntersectByKeys")
	return c
}

// IsEmpty returns true if the collection is empty; otherwise, false is returned.
func (c BaseCollection) IsEmpty() bool {
	return false
}

func (c BaseCollection) IsEmptyE() (bool, error) {
	c.errorHandle(ErrNotImplement, "IsEmptyE")
	return false, c.err
}

// IsNotEmpty returns true if the collection is not empty; otherwise, false is returned.
func (c BaseCollection) IsNotEmpty() bool {
	return false
}

func (c BaseCollection) IsNotEmptyE() (bool, error) {
	c.errorHandle(ErrNotImplement, "IsNotEmptyE")
	return false, c.err
}

// KeyBy keys the collection by the given key. If multiple items have the same key, only the last one will
// appear in the new collection.
func (c BaseCollection) KeyBy(interface{}) Collection {
	c.errorHandle(ErrNotImplement, "KeyBy")
	return c
}

// Keys returns all of the collection's keys.
func (c BaseCollection) Keys() Collection {
	c.errorHandle(ErrNotImplement, "Keys")
	return c
}

// Last returns the last element in the collection that passes a given truth test.
func (c BaseCollection) Last(...CB) interface{} {
	return nil
}

func (c BaseCollection) LastE(...CB) (interface{}, error) {
	c.errorHandle(ErrNotImplement, "LastE")
	return nil, c.err
}

// MapToGroups groups the collection's items by the given callback.
func (c BaseCollection) MapToGroups(MapCB) Collection {
	c.errorHandle(ErrNotImplement, "MapToGroups")
	return c
}

// MapWithKeys iterates through the collection and passes each value to the given callback.
func (c BaseCollection) MapWithKeys(MapCB) Collection {
	c.errorHandle(ErrNotImplement, "MapWithKeys")
	return c
}

// Median returns the median value of a given key.
func (c BaseCollection) Median(key ...string) decimal.Decimal {
	c.errorHandle(ErrNotImplement, "Median")
	return decimal.Decimal{}
}

// Merge merges the given array or collection with the original collection. If a string key in the given items
// matches a string key in the original collection, the given items's value will overwrite the value in the
// original collection.
func (c BaseCollection) Merge(interface{}) Collection {
	c.errorHandle(ErrNotImplement, "Merge")
	return c
}

func (c BaseCollection) Nth(...int) Collection {
	c.errorHandle(ErrNotImplement, "Nth")
	return c
}

// Pad will fill the array with the given value until the array reaches the specified size.
func (c BaseCollection) Pad(int, interface{}) Collection {
	c.errorHandle(ErrNotImplement, "Pad")
	return c
}

// Partition separate elements that pass a given truth test from those that do not.
func (c BaseCollection) Partition(PartCB) (Collection, Collection) {
	c.errorHandle(ErrNotImplement, "Partition")
	return c, c
}

// Pop removes and returns the last item from the collection.
func (c BaseCollection) Pop() interface{} {
	return nil
}

func (c BaseCollection) PopE() (interface{}, error) {
	c.errorHandle(ErrNotImplement, "PopE")
	return nil, c.err
}

// Push appends an item to the end of the collection.
func (c BaseCollection) Push(interface{}) Collection {
	c.errorHandle(ErrNotImplement, "Push")
	return c
}

// Random returns a random item from the collection.
func (c BaseCollection) Random(...int) Collection {
	c.errorHandle(ErrNotImplement, "Random")
	return c
}

// Reduce reduces the collection to a single value, passing the result of each iteration into the subsequent iteration.
func (c BaseCollection) Reduce(ReduceCB) interface{} {
	return nil
}

func (c BaseCollection) ReduceE(ReduceCB) (interface{}, error) {
	c.errorHandle(ErrNotImplement, "ReduceE")
	return nil, c.err
}

// Reject filters the collection using the given callback.
func (c BaseCollection) Reject(CB) Collection {
	c.errorHandle(ErrNotImplement, "Reject")
	return c
}

// Reverse reverses the order of the collection's items, preserving the original keys.
func (c BaseCollection) Reverse() Collection {
	c.errorHandle(ErrNotImplement, "Reverse")
	return c
}

// Search searches the collection for the given value and returns its key if found. If the item is not found,
// -1 is returned.
func (c BaseCollection) Search(interface{}) int {
	return -1
}

func (c BaseCollection) SearchE(interface{}) (int, error) {
	c.errorHandle(ErrNotImplement, "SearchE")
	return -1, c.err
}

// Shift removes and returns the first item from the collection.
func (c BaseCollection) Shift() Collection {
	c.errorHandle(ErrNotImplement, "Shift")
	return c
}

// Shuffle randomly shuffles the items in the collection.
func (c BaseCollection) Shuffle() Collection {
	c.errorHandle(ErrNotImplement, "Shuffle")
	return c
}

// Slice returns a slice of the collection starting at the given index.
func (c BaseCollection) Slice(...int) Collection {
	c.errorHandle(ErrNotImplement, "Slice")
	return c
}

// Sort sorts the collection.
func (c BaseCollection) Sort() Collection {
	c.errorHandle(ErrNotImplement, "Sort")
	return c
}

// SortByDesc has the same signature as the sortBy method, but will sort the collection in the opposite order.
func (c BaseCollection) SortByDesc() Collection {
	c.errorHandle(ErrNotImplement, "SortByDesc")
	return c
}

// Splice removes and returns a slice of items starting at the specified index.
func (c BaseCollection) Split(int) Collection {
	c.errorHandle(ErrNotImplement, "Split")
	return c
}

// Split breaks a collection into the given number of groups.
func (c BaseCollection) Splice(index ...int) Collection {
	c.errorHandle(ErrNotImplement, "Splice")
	return c
}

// Unique returns all of the unique items in the collection.
func (c BaseCollection) Unique() Collection {
	c.errorHandle(ErrNotImplement, "Unique")
	return c
}

// WhereIn filters the collection by a given key / value contained within the given array.
func (c BaseCollection) WhereIn(string, []interface{}) Collection {
	c.errorHandle(ErrNotImplement, "WhereIn")
	return c
}

// WhereNotIn filters the collection by a given key / value not contained within the given array.
func (c BaseCollection) WhereNotIn(string, []interface{}) Collection {
	c.errorHandle(ErrNotImplement, "WhereNotIn")
	return c
}

// ToJson converts the collection into a json string.
func (c BaseCollection) ToJson() string {
	s, err := json.Marshal(c.value)
	if err != nil {
		return ""
	}
	return string(s)
}

func (c BaseCollection) ToJsonE() (string, error) {
	s, err := json.Marshal(c.value)
	if err != nil {
		return "", err
	}
	return string(s), nil
}

// ToNumberArray converts the collection into a plain golang slice which contains decimal.Decimal.
func (c BaseCollection) ToNumberArray() []decimal.Decimal {
	return nil
}

func (c BaseCollection) ToNumberArrayE() ([]decimal.Decimal, error) {
	c.errorHandle(ErrNotImplement, "ToNumberArrayE")
	return nil, c.err
}

// ToStringArray converts the collection into a plain golang slice which contains string.
func (c BaseCollection) ToMultiDimensionalArray() [][]interface{} {
	return nil
}

func (c BaseCollection) ToMultiDimensionalArrayE() ([][]interface{}, error) {
	c.errorHandle(ErrNotImplement, "ToMultiDimensionalArrayE")
	return nil, c.err
}

// ToStringArray converts the collection into a plain golang slice which contains string.
func (c BaseCollection) ToStringArray() []string {
	return nil
}

func (c BaseCollection) ToStringArrayE() ([]string, error) {
	c.errorHandle(ErrNotImplement, "ToStringArray")
	return nil, c.err
}

// ToMap converts the collection into a plain golang map.
func (c BaseCollection) ToMap() map[string]interface{} {
	return nil
}

func (c BaseCollection) ToMapE() (map[string]interface{}, error) {
	c.errorHandle(ErrNotImplement, "ToMapE")
	return nil, c.err
}

// ToMapArray converts the collection into a plain golang slice which contains map.
func (c BaseCollection) ToMapArray() []map[string]interface{} {
	return nil
}

func (c BaseCollection) ToMapArrayE() ([]map[string]interface{}, error) {
	c.errorHandle(ErrNotImplement, "ToMapArrayE")
	return nil, c.err
}

// Where filters the collection by a given key / value pair.
func (c BaseCollection) Where(key string, values ...interface{}) Collection {
	c.errorHandle(ErrNotImplement, "Where")
	return c
}

// Count returns the total number of items in the collection.
func (c BaseCollection) Count() int {
	return c.length
}
