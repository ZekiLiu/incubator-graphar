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

import (
	"context"
	"io"
	"net/url"
)

// Reader abstracts reading GraphAr YAML from arbitrary backends (local FS, S3, HDFS, etc.).
type Reader interface {
	Open(ctx context.Context, uri *url.URL) (io.ReadCloser, error)
}

// Writer abstracts writing GraphAr YAML to arbitrary backends.
type Writer interface {
	Create(ctx context.Context, uri *url.URL) (io.WriteCloser, error)
}

// GraphInfoLoader defines how to load GraphAr info YAML into Go objects.
// Implementations can be backed by different storage systems via Reader.
type GraphInfoLoader interface {
	LoadGraph(ctx context.Context, graphYAMLURI *url.URL) (*GraphInfo, error)
	LoadVertex(ctx context.Context, vertexYAMLURI *url.URL) (*VertexInfo, error)
	LoadEdge(ctx context.Context, edgeYAMLURI *url.URL) (*EdgeInfo, error)
}

// GraphInfoSaver defines how to save GraphAr info objects back into YAML.
// Implementations can be backed by different storage systems via Writer.
type GraphInfoSaver interface {
	SaveGraph(ctx context.Context, graphYAMLURI *url.URL, gi *GraphInfo) error
	SaveVertex(ctx context.Context, vertexYAMLURI *url.URL, vi *VertexInfo) error
	SaveEdge(ctx context.Context, edgeYAMLURI *url.URL, ei *EdgeInfo) error
}
