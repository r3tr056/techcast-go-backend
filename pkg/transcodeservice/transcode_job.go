package transcodeservice

// import (
// 	"context"
// 	"fmt"
// 	"io"

// 	transcoder "cloud.google.com/go/video/transcoder/apiv1"
// 	transcoderpb "cloud.golang.org/genproto/googleapis/cloud/video/transcoder/v1"
// )

// func createJobFromPreset(w io.Writer, projectID string, location string, inputURI string, outputURI string, preset string) error {

// 	ctx := context.Background()
// 	client, err := transcoder.NewClient(ctx)
// 	if err != nil {
// 		return fmt.Errorf("NewClient : %v", err)
// 	}
// 	defer client.Close()

// 	req := &transcoderpb.CreateJobRequest{
// 		Parent: fmt.Sprintf("projects/%s/locations/%s", projectID, location),
// 		Job: &transcodepb.Job{
// 			InputUri: inputURI,
// 			OutputUri: outputURI,
// 			JobConfig: &transcoderpb.Job_TemplateId{
// 				TemplateId: preset,
// 			},
// 		},
// 	}

// 	// Creates the job, Jobs take a variable amount of time to run
// 	response, err := client.CreateJob(ctx, req)
// 	if err != nil {
// 		return fmt.Errorf("createJobFromPreset : %v", err)
// 	}

// 	fmt.Fprintf(w, "Job: %v", response.GetName())
// 	return nil
// }

// func pingStatus(w io.Writer, projectID , location, jobID string) error {

// 	ctx := context.Background()
// 	client, err := transcoder.NewClient(ctx)
// 	if err != nil {
// 		return fmt.Errorf("NewClient : %v", err)
// 	}

// 	defer client.Close()

// 	req := &transcoderpb.GetJobRequest{
// 		Name: fmt.Sprintf("projects/%s/locations/%s/jobs/%s", projectID, location, jobID),
// 	}

// 	response, err := client.GetJob(ctx, req)
// 	if err != nil {
// 		return fmt.Errorf("GetJob: %v", err)
// 	}

// 	fmt.Fprintf(w, "Job state : %v\n----\nJob failure reason : %v\n", response.State, response.Error)
// 	return nil
// }

// func createJobTemplate(w io.Writer, projectID, location, templateID string) error {
// 	ctx := context.Background()
// 	client, err := transcoder.NewClient(ctx)
// 	if err != nil {
// 		return fmt.Errorf("NewClient : %v", err)
// 	}

// 	defer client.Close()

// 	req := &transcoderpb.CreateJobTemplateRequest {
// 		Parent: fmt.Sprintf("projects/%s/locations/%s", projectID, location),
// 		JobTemplateId: templateID,
// 		JobTemplate: &transcoderpb.JobTemplate {
// 			Config: &transcoderpb.JobConfig {
// 				ElementaryStreams: []*transcoderpb.ElementaryStream {
// 					{
// 						Key: "video_stream0",
// 						ElementaryStream: &transcoderpb.ElementaryStream_VideoStream {
// 							VideoStream: &transcoderpb.VideoStream {
// 								CodecSettings: &transcoderpb.VideoStream_H264 {
// 									H264: &transcoderpb.VideoStream_H264CodecSettings {
// 										BitrateBps: 550000,
// 										FrameRate: 60,
// 										HeightPixels: 360,
// 										WidthPixels: 640,
// 									},
// 								},
// 							},
// 						},
// 					},
// 					{
// 						Key: "audio_stream0",
// 						ElementaryStream: &transcoderpb.ElementaryStream_AudioStream {
// 							AudioStream: &transcoderpb.AudioStream {
// 								Codec: "mp3",
// 								BitrateBps: 64000,
// 							},
// 						},
// 					},
// 				},
// 				MuxStreams: []*transcoderpb.MuxStream {
// 					{
// 						Key: "sd",
// 						Container: "mp3",
// 						ElementaryStreams: []string{"audio_stream0"},
// 					},
// 					{
// 						Key: "hd",
// 						Container: "mp3",
// 						ElementaryStreams: []string{"audio_stream0"}
// 					},
// 				},
// 			},
// 		},
// 	}

// 	response, err := client.CreateJobTemplate(ctx, req)
// 	if err != nil {
// 		return fmt.Errorf("CreatJobTemplate: %v", err)
// 	}

// 	fmt.Fprintf(w, "Job template: %v", err)
// 	return nil
// }
