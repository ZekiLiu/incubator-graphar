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

// This file aligns with Java `org.apache.graphar.util.GeneralParams`.
// These are exported for use within the graphar Go module but not for external users.

const (
	// Column names.
	VertexIndexCol      = "_graphArVertexIndex"
	SrcIndexCol         = "_graphArSrcIndex"
	DstIndexCol         = "_graphArDstIndex"
	OffsetCol           = "_graphArOffset"
	PrimaryCol          = "_graphArPrimary"
	VertexChunkIndexCol = "_graphArVertexChunkIndex"
	EdgeIndexCol        = "_graphArEdgeIndex"

	// Keys.
	OffsetStartChunkIndexKey = "_graphar_offset_start_chunk_index"
	AggNumListOfEdgeChunkKey = "_graphar_agg_num_list_of_edge_chunk"

	// Defaults.
	DefaultVertexChunkSize int64  = 262144  // 2^18
	DefaultEdgeChunkSize   int64  = 4194304 // 2^22
	DefaultFileType        string = "parquet"
	DefaultVersion         string = "v1"
)
