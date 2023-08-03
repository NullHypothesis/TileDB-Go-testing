package tiledb

import (
	"reflect"
)

// genericType returns the reflect.Type for T
func genericType[T any]() reflect.Type {
	return reflect.TypeOf((*T)(nil)).Elem()
}

// datatypeOfDimensionFromIndex returns for a dimension the Datatype and whether it is of variable size
func datatypeOfDimensionFromIndex(arr *Array, dimIdx uint32) (Datatype, bool, error) {
	schema, err := arr.Schema()
	if err != nil {
		return Datatype(0), false, err
	}

	domain, err := schema.Domain()
	if err != nil {
		return Datatype(0), false, err
	}

	dimension, err := domain.DimensionFromIndex(uint(dimIdx))
	if err != nil {
		return Datatype(0), false, err
	}

	datatype, err := dimension.Type()
	if err != nil {
		return Datatype(0), false, err
	}

	cellValNum, err := dimension.CellValNum()
	if err != nil {
		return Datatype(0), false, err
	}

	return datatype, cellValNum == TILEDB_VAR_NUM, nil
}

// datatypeOfDimensionFromName returns for a dimension the Datatype and whether it is of variable size
func datatypeOfDimensionFromName(arr *Array, dimName string) (Datatype, bool, error) {
	schema, err := arr.Schema()
	if err != nil {
		return Datatype(0), false, err
	}

	domain, err := schema.Domain()
	if err != nil {
		return Datatype(0), false, err
	}

	dimension, err := domain.DimensionFromName(dimName)
	if err != nil {
		return Datatype(0), false, err
	}

	datatype, err := dimension.Type()
	if err != nil {
		return Datatype(0), false, err
	}

	cellValNum, err := dimension.CellValNum()
	if err != nil {
		return Datatype(0), false, err
	}

	return datatype, cellValNum == TILEDB_VAR_NUM, nil
}
