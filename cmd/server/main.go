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
	"io"
	"log"
	"net"
	"sync"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/matttproud/golang_protobuf_extensions/pbutil"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	buildeventstream "github.com/bazelbuild/bazel/src/main/java/com/google/devtools/build/lib/buildeventstream/proto"
	build "google.golang.org/genproto/googleapis/devtools/build/v1"
)

const (
	port = "localhost:50051"
)

type streamState struct {
	committedSequences int64
	bazelBuildEvents   *bytes.Buffer
}

type buildEventServer struct {
	instanceName string

	lock    sync.Mutex
	streams map[string]*streamState
}

func (bes *buildEventServer) PublishLifecycleEvent(ctx context.Context, in *build.PublishLifecycleEventRequest) (*empty.Empty, error) {
	// For now, completely ignore lifecycle events.
	return &empty.Empty{}, nil
}

func (bes *buildEventServer) processBuildToolEvent(ctx context.Context, in *build.PublishBuildToolEventStreamRequest) (*build.PublishBuildToolEventStreamResponse, error) {
	bes.lock.Lock()
	defer bes.lock.Unlock()

	var streamId *build.StreamId
	streamId = in.OrderedBuildEvent.StreamId
	key := streamId.InvocationId

	state, ok := bes.streams[key]
	if ok {
		if in.OrderedBuildEvent.SequenceNumber < state.committedSequences+1 {
			// Historical event that was retransmitted. Nothing to do.
			return &build.PublishBuildToolEventStreamResponse{
				StreamId:       in.OrderedBuildEvent.StreamId,
				SequenceNumber: state.committedSequences,
			}, nil
		} else if in.OrderedBuildEvent.SequenceNumber > state.committedSequences+1 {
			// Event from the future.
			return nil, status.Error(codes.InvalidArgument, "Event has sequence number from the future")
		}
	} else {
		if in.OrderedBuildEvent.SequenceNumber != 1 {
			return nil, status.Error(codes.DataLoss, "Stream is not known by the server")
		}
		state = &streamState{
			committedSequences: 0,
			bazelBuildEvents:   bytes.NewBuffer(nil),
		}
		bes.streams[key] = state
	}

	switch buildEvent := in.OrderedBuildEvent.Event.Event.(type) {
	case *build.BuildEvent_ComponentStreamFinished:
		log.Print("BuildTool: ComponentStreamFinished: ", buildEvent.ComponentStreamFinished)
		delete(bes.streams, key)
	case *build.BuildEvent_BazelEvent:
		log.Print("DEBUG1")
		var bazelBuildEvent buildeventstream.BuildEvent
		if err := ptypes.UnmarshalAny(buildEvent.BazelEvent, &bazelBuildEvent); err != nil {
			return nil, err
		}
		if _, err := pbutil.WriteDelimited(state.bazelBuildEvents, &bazelBuildEvent); err != nil {
			return nil, err
		}
	default:
		log.Print("Received unknown BuildToolEvent")
	}

	state.committedSequences++
	return &build.PublishBuildToolEventStreamResponse{
		StreamId:       in.OrderedBuildEvent.StreamId,
		SequenceNumber: state.committedSequences,
	}, nil
}

func (bes *buildEventServer) PublishBuildToolEventStream(stream build.PublishBuildEvent_PublishBuildToolEventStreamServer) error {
	responseStream := make(chan *build.PublishBuildToolEventStreamResponse, 100)
	requestError := make(chan error, 1)

	// Handle incoming requests.
	go func() {
		defer close(responseStream)
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				requestError <- nil
				return
			}
			if err != nil {
				requestError <- err
				return
			}
			response, err := bes.processBuildToolEvent(stream.Context(), in)
			if err != nil {
				requestError <- err
				return
			}
			responseStream <- response
		}
	}()

	// Stream responses back to the client asynchronously to reduce
	// round-trips.
	for response := range responseStream {
		if err := stream.Send(response); err != nil {
			return err
		}
	}
	return <-requestError
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	build.RegisterPublishBuildEventServer(s, &buildEventServer{
		instanceName: "bazel-build-trends",
		streams: map[string]*streamState{},
	})

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
