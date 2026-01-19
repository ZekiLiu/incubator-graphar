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

import "net/url"

// VersionParser defines how to parse and format GraphAr info versions.
// It corresponds to Java VersionParser + VersionInfo.toString/checkType.
//
// We keep it as an interface so we can define the public contract without providing an implementation yet.
type VersionParser interface {
	Parse(version string) (*VersionInfo, error)
	Format(v *VersionInfo) (string, error)
	CheckType(v *VersionInfo, typeStr string) bool
}

// Validator defines validation behavior for info objects.
// It corresponds to Java `isValidated()` but returns an error in Go.
type Validator interface {
	Validate() error
}

// GraphInfoAPI defines GraphInfo behaviors that align with Java GraphInfo methods.
// This is an interface-only contract; implementations will be added later.
type GraphInfoAPI interface {
	Validator

	GetName() string
	GetPrefix() *url.URL
	GetVersion() *VersionInfo
	GetVertexInfos() []*VertexInfo
	GetEdgeInfos() []*EdgeInfo

	HasVertexInfo(vertexType string) bool
	HasEdgeInfo(srcType, edgeType, dstType string) bool

	GetVertexInfo(vertexType string) (*VertexInfo, error)
	GetEdgeInfo(srcType, edgeType, dstType string) (*EdgeInfo, error)

	SetStoreURIForVertex(vertexType string, storeURI *url.URL) error
	SetStoreURIForEdge(srcType, edgeType, dstType string, storeURI *url.URL) error

	GetStoreURIForVertex(vertexType string) (*url.URL, error)
	GetStoreURIForEdge(srcType, edgeType, dstType string) (*url.URL, error)

	GetTypes2StoreURI() map[string]*url.URL
}

// VertexInfoAPI defines VertexInfo behaviors.
type VertexInfoAPI interface {
	Validator

	GetType() string
	GetChunkSize() int64
	GetLabels() []string
	GetPropertyGroups() []*PropertyGroup
	GetPrefix() *url.URL
	GetVersion() *VersionInfo

	HasProperty(propertyName string) bool
	IsPrimaryKey(propertyName string) (bool, error)
	IsNullableKey(propertyName string) (bool, error)
	GetPropertyType(propertyName string) (DataType, error)
	GetPropertyCardinality(propertyName string) (Cardinality, error)
	GetPropertyGroup(propertyName string) (*PropertyGroup, error)

	GetPropertyGroupURI(pg *PropertyGroup) (*url.URL, error)
	GetPropertyGroupChunkURI(pg *PropertyGroup, chunkIndex int64) (*url.URL, error)
	GetVerticesNumFileURI() (*url.URL, error)
}

// EdgeInfoAPI defines EdgeInfo behaviors.
type EdgeInfoAPI interface {
	Validator

	GetSrcType() string
	GetEdgeType() string
	GetDstType() string
	GetConcat() string
	GetChunkSize() int64
	GetSrcChunkSize() int64
	GetDstChunkSize() int64
	IsDirected() bool
	GetPrefix() *url.URL
	GetVersion() *VersionInfo

	HasAdjListType(adjListType AdjListType) bool
	HasProperty(propertyName string) bool
	IsPrimaryKey(propertyName string) (bool, error)
	IsNullableKey(propertyName string) (bool, error)
	GetPropertyType(propertyName string) (DataType, error)
	GetPropertyGroup(propertyName string) (*PropertyGroup, error)

	GetAdjacentLists() []*AdjacentList
	GetAdjacentList(adjListType AdjListType) (*AdjacentList, error)
	GetPropertyGroups() []*PropertyGroup

	GetAdjacentListURI(adjListType AdjListType) (*url.URL, error)
	GetAdjacentListChunkURI(adjListType AdjListType, vertexChunkIndex int64) (*url.URL, error)
	GetOffsetURI(adjListType AdjListType) (*url.URL, error)
	GetOffsetChunkURI(adjListType AdjListType, vertexChunkIndex int64) (*url.URL, error)
	GetVerticesNumFileURI(adjListType AdjListType) (*url.URL, error)
	GetEdgesNumFileURI(adjListType AdjListType, vertexChunkIndex int64) (*url.URL, error)

	GetPropertyGroupURI(pg *PropertyGroup) (*url.URL, error)
	GetPropertyGroupChunkURI(pg *PropertyGroup, chunkIndex int64) (*url.URL, error)
}

// PropertyGroupAPI defines PropertyGroup behaviors.
type PropertyGroupAPI interface {
	Validator

	GetProperties() []*Property
	GetFileType() FileType
	GetPrefix() *url.URL

	HasProperty(propertyName string) bool
}

// AdjacentListAPI defines AdjacentList behaviors.
type AdjacentListAPI interface {
	Validator

	GetType() AdjListType
	GetFileType() FileType
	GetPrefix() *url.URL
}

// PropertyAPI defines Property behaviors.
type PropertyAPI interface {
	GetName() string
	GetDataType() DataType
	GetCardinality() Cardinality
	IsPrimary() bool
	IsNullable() bool
}

// ConcatEdgeTriplet joins src/edge/dst types into a stable key.
// It aligns with Java EdgeInfo.concat using GeneralParams.regularSeperator.
func ConcatEdgeTriplet(srcType, edgeType, dstType string) string {
	return srcType + RegularSeparator + edgeType + RegularSeparator + dstType
}
