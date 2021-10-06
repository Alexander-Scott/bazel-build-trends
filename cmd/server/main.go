/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package main implements a server for Greeter service.
package main

import (
	"bytes"
	"context"
	"log"
	"net"
	"sync"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"

	build "google.golang.org/genproto/googleapis/devtools/build/v1"
)

const (
	port = "localhost:50051"
)

type streamState struct {
	committedSequences int64
	bazelBuildEvents   *bytes.Buffer
}

// server is used to implement helloworld.GreeterServer.
type buildEventServer struct {
	build.UnimplementedPublishBuildEventServer
	lock    sync.Mutex
	streams map[string]*streamState
}

func (bes *buildEventServer) PublishLifecycleEvent(ctx context.Context, in *build.PublishLifecycleEventRequest) (*empty.Empty, error) {
	// For now, completely ignore lifecycle events.
	return &empty.Empty{}, nil
}

func (bes *buildEventServer) processBuildToolEvent(ctx context.Context, in *build.PublishBuildToolEventStreamRequest) (*build.PublishBuildToolEventStreamResponse, error) {
	return nil, nil
}

func (bes *buildEventServer) PublishBuildToolEventStream(stream build.PublishBuildEvent_PublishBuildToolEventStreamServer) error {
	// For now, completely ignore lifecycle events.
	return nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	build.RegisterPublishBuildEventServer(s, &buildEventServer{})

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
