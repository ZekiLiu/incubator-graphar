// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package info

import "fmt"

// RegularSeparator is used to concat edge triplets.
// It aligns with Java `GeneralParams.regularSeperator`.
const RegularSeparator = "_"

type DataType string

const (
	DataTypeBool      DataType = "bool"
	DataTypeInt32     DataType = "int32"
	DataTypeInt64     DataType = "int64"
	DataTypeFloat     DataType = "float"
	DataTypeDouble    DataType = "double"
	DataTypeString    DataType = "string"
	DataTypeList      DataType = "list"
	DataTypeDate      DataType = "date"
	DataTypeTimestamp DataType = "timestamp"
)

func ParseDataType(s string) (DataType, error) {
	switch s {
	case string(DataTypeBool), string(DataTypeInt32), string(DataTypeInt64), string(DataTypeFloat), string(DataTypeDouble), string(DataTypeString), string(DataTypeList), string(DataTypeDate), string(DataTypeTimestamp):
		return DataType(s), nil
	default:
		return "", fmt.Errorf("unknown data type: %s", s)
	}
}

type FileType string

const (
	FileTypeCSV     FileType = "csv"
	FileTypeParquet FileType = "parquet"
	FileTypeORC     FileType = "orc"
)

func ParseFileType(s string) (FileType, error) {
	switch s {
	case string(FileTypeCSV), string(FileTypeParquet), string(FileTypeORC):
		return FileType(s), nil
	default:
		return "", fmt.Errorf("unknown file type: %s", s)
	}
}

type Cardinality string

const (
	CardinalitySingle Cardinality = "single"
	CardinalityList   Cardinality = "list"
	CardinalitySet    Cardinality = "set"
)

func ParseCardinality(s string) (Cardinality, error) {
	switch s {
	case string(CardinalitySingle), string(CardinalityList), string(CardinalitySet):
		return Cardinality(s), nil
	default:
		return "", fmt.Errorf("unknown cardinality: %s", s)
	}
}

type AdjListType string

const (
	AdjListUnorderedBySource AdjListType = "unordered_by_source"
	AdjListUnorderedByDest   AdjListType = "unordered_by_dest"
	AdjListOrderedBySource   AdjListType = "ordered_by_source"
	AdjListOrderedByDest     AdjListType = "ordered_by_dest"
)

func AdjListTypeFromOrderedAndAlignedBy(ordered bool, alignedBy string) (AdjListType, error) {
	switch alignedBy {
	case "src":
		if ordered {
			return AdjListOrderedBySource, nil
		}
		return AdjListUnorderedBySource, nil
	case "dst":
		if ordered {
			return AdjListOrderedByDest, nil
		}
		return AdjListUnorderedByDest, nil
	default:
		return "", fmt.Errorf("invalid alignedBy: %s", alignedBy)
	}
}

func (t AdjListType) IsOrdered() bool {
	return t == AdjListOrderedBySource || t == AdjListOrderedByDest
}

func (t AdjListType) AlignedBy() string {
	if t == AdjListOrderedBySource || t == AdjListUnorderedBySource {
		return "src"
	}
	return "dst"
}
