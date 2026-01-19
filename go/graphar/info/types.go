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

// VersionInfo aligns with Java `org.apache.graphar.info.VersionInfo`.
// String form is like: "gar/v1" or "gar/v1 (type1,type2)".
type VersionInfo struct {
	Version          int
	UserDefinedTypes []string
}

// GraphInfo is the top-level metadata object of a GraphAr graph.
// This is a type definition only (no YAML parsing/serialization logic here).
type GraphInfo struct {
	Name string

	// Base URI of the graph, corresponding to `prefix` in YAML.
	Prefix *url.URL

	Version *VersionInfo

	VertexInfos []*VertexInfo
	EdgeInfos   []*EdgeInfo

	// Store URIs for vertex/edge info YAMLs (useful for round-tripping).
	// Key conventions align with Java GraphInfo:
	// - vertex: "<type>.vertex"
	// - edge: "<src>_<edge>_<dst>.edge"
	Types2StoreURI map[string]*url.URL
}

// VertexInfo describes a vertex type in GraphAr.
type VertexInfo struct {
	Type      string
	ChunkSize int64

	PropertyGroups []*PropertyGroup
	Labels         []string

	// Base URI of the vertex type, corresponding to `prefix` in YAML.
	Prefix *url.URL

	Version *VersionInfo
}

// EdgeInfo describes an edge type in GraphAr.
type EdgeInfo struct {
	SrcType  string
	EdgeType string
	DstType  string

	ChunkSize    int64
	SrcChunkSize int64
	DstChunkSize int64
	Directed     bool

	AdjacentLists  []*AdjacentList
	PropertyGroups []*PropertyGroup

	// Base URI of the edge type, corresponding to `prefix` in YAML.
	Prefix *url.URL

	Version *VersionInfo
}

// AdjacentList describes an adjacency list variant for an edge type.
type AdjacentList struct {
	Type     AdjListType
	FileType FileType
	Prefix   *url.URL
}

// PropertyGroup describes how a set of properties are stored together.
type PropertyGroup struct {
	Properties []*Property
	FileType   FileType
	Prefix     *url.URL
}

// Property describes a single vertex/edge property.
type Property struct {
	Name        string
	DataType    DataType
	Cardinality Cardinality
	Primary     bool
	Nullable    bool
}
