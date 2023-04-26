// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by grpc_generate_tools(grpcgengo) at
//

package clouddisk

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	stdproto "github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/any"
	"gitlab.flyele.vip/server-side/go-zero/v2/rpcx/grpcbase"
	pb "gitlab.flyele.vip/server-side/go-zero/v2/rpcx/protos"
)

func makeFileDetailEndpoint(svc Repository) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(FileDetailRequest)
		v := svc.FileDetail(ctx, req)
		return v, nil
	}
}
func decodeFileDetailRequest(ctx context.Context, req interface{}) (interface{}, error) {
	r := req.(*pb.FileDetailRequest)
	var xcontext *UserContext
	if r.Context != nil {
		xcontext = &UserContext{
			UserID:        r.Context.UserID,
			Platform:      r.Context.Platform,
			ClientVersion: r.Context.ClientVersion,
			Token:         r.Context.Token,
			ClientIP:      r.Context.ClientIP,
			RequestID:     r.Context.RequestID,
		}
	}

	return FileDetailRequest{
		Files:   r.Files,
		Context: xcontext,
	}, nil
}

func encodeFileDetailResponse(ctx context.Context, req interface{}) (interface{}, error) {
	r := req.(grpcbase.Response)
	if r.Data == nil {
		return &pb.Response{
			Message: r.Message,
		}, nil
	}
	//在所有类型中匹配名称相同的消息名称
	resp := r.Data.(FileDetailResponse)
	//是复杂类型的数组
	var detailArray = make([]*pb.FileDetail, 0)

	for _, v := range resp.Detail {
		detailArray = append(detailArray, &pb.FileDetail{
			FileID:   v.FileID,
			FileName: v.FileName,
		})
	}
	pbresp := &pb.FileDetailResponse{
		Detail: detailArray,
	}
	b, err := stdproto.Marshal(pbresp)
	if err != nil {
		return nil, err
	}
	anyData := &any.Any{
		Value:   b,
		TypeUrl: stdproto.MessageName(pbresp),
	}
	return &pb.Response{
		Message: r.Message,
		Total:   int32(r.Total),
		Data:    anyData,
	}, nil
}

func makeAssociatedFilesEndpoint(svc Repository) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(AssociatedFilesRequest)
		v := svc.AssociatedFiles(ctx, req)
		return v, nil
	}
}
func decodeAssociatedFilesRequest(ctx context.Context, req interface{}) (interface{}, error) {
	r := req.(*pb.AssociatedFilesRequest)
	var xcontext *UserContext
	if r.Context != nil {
		xcontext = &UserContext{
			UserID:        r.Context.UserID,
			Platform:      r.Context.Platform,
			ClientVersion: r.Context.ClientVersion,
			Token:         r.Context.Token,
			ClientIP:      r.Context.ClientIP,
			RequestID:     r.Context.RequestID,
		}
	}

	return AssociatedFilesRequest{
		FlyeleID: r.FlyeleID,
		Context:  xcontext,
	}, nil
}

func encodeAssociatedFilesResponse(ctx context.Context, req interface{}) (interface{}, error) {
	r := req.(grpcbase.Response)
	if r.Data == nil {
		return &pb.Response{
			Message: r.Message,
		}, nil
	}
	//在所有类型中匹配名称相同的消息名称
	resp := r.Data.(AssociatedFilesResponse)
	pbresp := &pb.AssociatedFilesResponse{
		Count: resp.Count,
	}
	b, err := stdproto.Marshal(pbresp)
	if err != nil {
		return nil, err
	}
	anyData := &any.Any{
		Value:   b,
		TypeUrl: stdproto.MessageName(pbresp),
	}
	return &pb.Response{
		Message: r.Message,
		Total:   int32(r.Total),
		Data:    anyData,
	}, nil
}

func makeUploadFileEndpoint(svc Repository) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(UploadFileRequest)
		v := svc.UploadFile(ctx, req)
		return v, nil
	}
}
func decodeUploadFileRequest(ctx context.Context, req interface{}) (interface{}, error) {
	r := req.(*pb.UploadFileRequest)
	var xcontext *UserContext
	if r.Context != nil {
		xcontext = &UserContext{
			UserID:        r.Context.UserID,
			Platform:      r.Context.Platform,
			ClientVersion: r.Context.ClientVersion,
			Token:         r.Context.Token,
			ClientIP:      r.Context.ClientIP,
			RequestID:     r.Context.RequestID,
		}
	}

	return UploadFileRequest{
		FilePath:      r.FilePath,
		Content:       r.Content,
		OssBucketName: r.OssBucketName,
		Context:       xcontext,
	}, nil
}

func encodeUploadFileResponse(ctx context.Context, req interface{}) (interface{}, error) {
	r := req.(grpcbase.Response)
	if r.Data == nil {
		return &pb.Response{
			Message: r.Message,
		}, nil
	}
	//在所有类型中匹配名称相同的消息名称
	return &pb.Response{
		Message: r.Message,
		Total:   int32(r.Total),
	}, nil
}

func makeGetFilesNameEndpoint(svc Repository) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetFilesNameRequest)
		v := svc.GetFilesName(ctx, req)
		return v, nil
	}
}
func decodeGetFilesNameRequest(ctx context.Context, req interface{}) (interface{}, error) {
	r := req.(*pb.GetFilesNameRequest)
	var xcontext *UserContext
	if r.Context != nil {
		xcontext = &UserContext{
			UserID:        r.Context.UserID,
			Platform:      r.Context.Platform,
			ClientVersion: r.Context.ClientVersion,
			Token:         r.Context.Token,
			ClientIP:      r.Context.ClientIP,
			RequestID:     r.Context.RequestID,
		}
	}

	return GetFilesNameRequest{
		FlyeleID: r.FlyeleID,
		FilesID:  r.FilesID,
		Context:  xcontext,
	}, nil
}

func encodeGetFilesNameResponse(ctx context.Context, req interface{}) (interface{}, error) {
	r := req.(grpcbase.Response)
	if r.Data == nil {
		return &pb.Response{
			Message: r.Message,
		}, nil
	}
	//在所有类型中匹配名称相同的消息名称
	resp := r.Data.(GetFilesNameResponse)
	pbresp := &pb.GetFilesNameResponse{
		Names: resp.Names,
	}
	b, err := stdproto.Marshal(pbresp)
	if err != nil {
		return nil, err
	}
	anyData := &any.Any{
		Value:   b,
		TypeUrl: stdproto.MessageName(pbresp),
	}
	return &pb.Response{
		Message: r.Message,
		Total:   int32(r.Total),
		Data:    anyData,
	}, nil
}

func makeBatchUnbindFileEndpoint(svc Repository) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(BatchUnbindFileRequest)
		v := svc.BatchUnbindFile(ctx, req)
		return v, nil
	}
}
func decodeBatchUnbindFileRequest(ctx context.Context, req interface{}) (interface{}, error) {
	r := req.(*pb.BatchUnbindFileRequest)
	var xcontext *UserContext
	if r.Context != nil {
		xcontext = &UserContext{
			UserID:        r.Context.UserID,
			Platform:      r.Context.Platform,
			ClientVersion: r.Context.ClientVersion,
			Token:         r.Context.Token,
			ClientIP:      r.Context.ClientIP,
			RequestID:     r.Context.RequestID,
		}
	}

	return BatchUnbindFileRequest{
		FlyeleID:   r.FlyeleID,
		FlyeleType: r.FlyeleType,
		CommentID:  r.CommentID,
		FilesID:    r.FilesID,
		Context:    xcontext,
	}, nil
}

func encodeBatchUnbindFileResponse(ctx context.Context, req interface{}) (interface{}, error) {
	r := req.(grpcbase.Response)
	if r.Data == nil {
		return &pb.Response{
			Message: r.Message,
		}, nil
	}
	//在所有类型中匹配名称相同的消息名称
	return &pb.Response{
		Message: r.Message,
		Total:   int32(r.Total),
	}, nil
}

func makeBatchBindFileEndpoint(svc Repository) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(BatchBindFileRequest)
		v := svc.BatchBindFile(ctx, req)
		return v, nil
	}
}
func decodeBatchBindFileRequest(ctx context.Context, req interface{}) (interface{}, error) {
	r := req.(*pb.BatchBindFileRequest)
	//是复杂类型的数组
	var xfileSensorsConfigs = make([]*FileSensorsConfig, 0)

	for _, v := range r.FileSensorsConfigs {
		xfileSensorsConfigs = append(xfileSensorsConfigs, &FileSensorsConfig{
			FileID: v.FileID,
			Origin: v.Origin,
			Entry:  v.Entry,
		})
	}
	var xcontext *UserContext
	if r.Context != nil {
		xcontext = &UserContext{
			UserID:        r.Context.UserID,
			Platform:      r.Context.Platform,
			ClientVersion: r.Context.ClientVersion,
			Token:         r.Context.Token,
			ClientIP:      r.Context.ClientIP,
			RequestID:     r.Context.RequestID,
		}
	}

	return BatchBindFileRequest{
		FlyeleType:         r.FlyeleType,
		FlyeleID:           r.FlyeleID,
		CommentID:          r.CommentID,
		Catalog:            r.Catalog,
		Entry:              r.Entry,
		FileUpload:         r.FileUpload,
		FileSensorsConfigs: xfileSensorsConfigs,
		Context:            xcontext,
	}, nil
}

func encodeBatchBindFileResponse(ctx context.Context, req interface{}) (interface{}, error) {
	r := req.(grpcbase.Response)
	if r.Data == nil {
		return &pb.Response{
			Message: r.Message,
		}, nil
	}
	//在所有类型中匹配名称相同的消息名称
	return &pb.Response{
		Message: r.Message,
		Total:   int32(r.Total),
	}, nil
}

func makeFileSyncEndpoint(svc Repository) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(FileSyncRequest)
		v := svc.FileSync(ctx, req)
		return v, nil
	}
}
func decodeFileSyncRequest(ctx context.Context, req interface{}) (interface{}, error) {
	r := req.(*pb.FileSyncRequest)
	var xcontext *UserContext
	if r.Context != nil {
		xcontext = &UserContext{
			UserID:        r.Context.UserID,
			Platform:      r.Context.Platform,
			ClientVersion: r.Context.ClientVersion,
			Token:         r.Context.Token,
			ClientIP:      r.Context.ClientIP,
			RequestID:     r.Context.RequestID,
		}
	}

	return FileSyncRequest{
		RequesterUID: r.RequesterUID,
		FlyeleID:     r.FlyeleID,
		FlyeleType:   r.FlyeleType,
		Context:      xcontext,
	}, nil
}

func encodeFileSyncResponse(ctx context.Context, req interface{}) (interface{}, error) {
	r := req.(grpcbase.Response)
	if r.Data == nil {
		return &pb.Response{
			Message: r.Message,
		}, nil
	}
	//在所有类型中匹配名称相同的消息名称
	return &pb.Response{
		Message: r.Message,
		Total:   int32(r.Total),
	}, nil
}

func makeTaskFilesEndpoint(svc Repository) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(TaskFilesRequest)
		v := svc.TaskFiles(ctx, req)
		return v, nil
	}
}
func decodeTaskFilesRequest(ctx context.Context, req interface{}) (interface{}, error) {
	r := req.(*pb.TaskFilesRequest)
	var xcontext *UserContext
	if r.Context != nil {
		xcontext = &UserContext{
			UserID:        r.Context.UserID,
			Platform:      r.Context.Platform,
			ClientVersion: r.Context.ClientVersion,
			Token:         r.Context.Token,
			ClientIP:      r.Context.ClientIP,
			RequestID:     r.Context.RequestID,
		}
	}

	return TaskFilesRequest{
		TaskID:  r.TaskID,
		Context: xcontext,
	}, nil
}

func encodeTaskFilesResponse(ctx context.Context, req interface{}) (interface{}, error) {
	r := req.(grpcbase.Response)
	if r.Data == nil {
		return &pb.Response{
			Message: r.Message,
		}, nil
	}
	//在所有类型中匹配名称相同的消息名称
	resp := r.Data.(TaskFilesResponse)
	//是复杂类型的数组
	var taskfilesArray = make([]*pb.TaskFiles, 0)

	for _, v := range resp.TaskFiles {
		taskfilesArray = append(taskfilesArray, &pb.TaskFiles{
			FileID:    v.FileID,
			FileName:  v.FileName,
			Resources: v.Resources,
			CreatorID: v.CreatorID,
			CommentID: v.CommentID,
		})
	}
	pbresp := &pb.TaskFilesResponse{
		TaskFiles: taskfilesArray,
	}
	b, err := stdproto.Marshal(pbresp)
	if err != nil {
		return nil, err
	}
	anyData := &any.Any{
		Value:   b,
		TypeUrl: stdproto.MessageName(pbresp),
	}
	return &pb.Response{
		Message: r.Message,
		Total:   int32(r.Total),
		Data:    anyData,
	}, nil
}

func makeGetBindLogWithAttachmentEndpoint(svc Repository) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetBindLogWithAttachmentRequest)
		v := svc.GetBindLogWithAttachment(ctx, req)
		return v, nil
	}
}
func decodeGetBindLogWithAttachmentRequest(ctx context.Context, req interface{}) (interface{}, error) {
	r := req.(*pb.GetBindLogWithAttachmentRequest)
	var xcontext *UserContext
	if r.Context != nil {
		xcontext = &UserContext{
			UserID:        r.Context.UserID,
			Platform:      r.Context.Platform,
			ClientVersion: r.Context.ClientVersion,
			Token:         r.Context.Token,
			ClientIP:      r.Context.ClientIP,
			RequestID:     r.Context.RequestID,
		}
	}

	return GetBindLogWithAttachmentRequest{
		FileId:     r.FileId,
		FlyeleId:   r.FlyeleId,
		FlyeleType: r.FlyeleType,
		Context:    xcontext,
	}, nil
}

func encodeGetBindLogWithAttachmentResponse(ctx context.Context, req interface{}) (interface{}, error) {
	r := req.(grpcbase.Response)
	if r.Data == nil {
		return &pb.Response{
			Message: r.Message,
		}, nil
	}
	//在所有类型中匹配名称相同的消息名称
	resp := r.Data.(GetBindLogWithAttachmentResponse)
	pbresp := &pb.GetBindLogWithAttachmentResponse{
		BoundBy:    resp.BoundBy,
		CommentId:  resp.CommentId,
		CreatorId:  resp.CreatorId,
		FileId:     resp.FileId,
		FlyeleId:   resp.FlyeleId,
		FlyeleType: resp.FlyeleType,
	}
	b, err := stdproto.Marshal(pbresp)
	if err != nil {
		return nil, err
	}
	anyData := &any.Any{
		Value:   b,
		TypeUrl: stdproto.MessageName(pbresp),
	}
	return &pb.Response{
		Message: r.Message,
		Total:   int32(r.Total),
		Data:    anyData,
	}, nil
}

func makeGetFileAddrByFileIDEndpoint(svc Repository) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetFileAddrByFileIDRequest)
		v := svc.GetFileAddrByFileID(ctx, req)
		return v, nil
	}
}
func decodeGetFileAddrByFileIDRequest(ctx context.Context, req interface{}) (interface{}, error) {
	r := req.(*pb.GetFileAddrByFileIDRequest)
	var xcontext *UserContext
	if r.Context != nil {
		xcontext = &UserContext{
			UserID:        r.Context.UserID,
			Platform:      r.Context.Platform,
			ClientVersion: r.Context.ClientVersion,
			Token:         r.Context.Token,
			ClientIP:      r.Context.ClientIP,
			RequestID:     r.Context.RequestID,
		}
	}

	return GetFileAddrByFileIDRequest{
		FileId:  r.FileId,
		Context: xcontext,
	}, nil
}

func encodeGetFileAddrByFileIDResponse(ctx context.Context, req interface{}) (interface{}, error) {
	r := req.(grpcbase.Response)
	if r.Data == nil {
		return &pb.Response{
			Message: r.Message,
		}, nil
	}
	//在所有类型中匹配名称相同的消息名称
	resp := r.Data.(GetFileAddrByFileIDResponse)
	pbresp := &pb.GetFileAddrByFileIDResponse{
		FileAddr: resp.FileAddr,
	}
	b, err := stdproto.Marshal(pbresp)
	if err != nil {
		return nil, err
	}
	anyData := &any.Any{
		Value:   b,
		TypeUrl: stdproto.MessageName(pbresp),
	}
	return &pb.Response{
		Message: r.Message,
		Total:   int32(r.Total),
		Data:    anyData,
	}, nil
}

func makeGetBindFileSizeEndpoint(svc Repository) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetBindFileSizeRequest)
		v := svc.GetBindFileSize(ctx, req)
		return v, nil
	}
}
func decodeGetBindFileSizeRequest(ctx context.Context, req interface{}) (interface{}, error) {
	r := req.(*pb.GetBindFileSizeRequest)
	var xcontext *UserContext
	if r.Context != nil {
		xcontext = &UserContext{
			UserID:        r.Context.UserID,
			Platform:      r.Context.Platform,
			ClientVersion: r.Context.ClientVersion,
			Token:         r.Context.Token,
			ClientIP:      r.Context.ClientIP,
			RequestID:     r.Context.RequestID,
		}
	}

	return GetBindFileSizeRequest{
		FlyeleID:   r.FlyeleID,
		FlyeleType: r.FlyeleType,
		Context:    xcontext,
	}, nil
}

func encodeGetBindFileSizeResponse(ctx context.Context, req interface{}) (interface{}, error) {
	r := req.(grpcbase.Response)
	if r.Data == nil {
		return &pb.Response{
			Message: r.Message,
		}, nil
	}
	//在所有类型中匹配名称相同的消息名称
	resp := r.Data.(GetBindFileSizeResponse)
	pbresp := &pb.GetBindFileSizeResponse{
		TotalSize: resp.TotalSize,
	}
	b, err := stdproto.Marshal(pbresp)
	if err != nil {
		return nil, err
	}
	anyData := &any.Any{
		Value:   b,
		TypeUrl: stdproto.MessageName(pbresp),
	}
	return &pb.Response{
		Message: r.Message,
		Total:   int32(r.Total),
		Data:    anyData,
	}, nil
}

func makeGetFileByFileIDEndpoint(svc Repository) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetFileByFileIDRequest)
		v := svc.GetFileByFileID(ctx, req)
		return v, nil
	}
}
func decodeGetFileByFileIDRequest(ctx context.Context, req interface{}) (interface{}, error) {
	r := req.(*pb.GetFileByFileIDRequest)
	var xcontext *UserContext
	if r.Context != nil {
		xcontext = &UserContext{
			UserID:        r.Context.UserID,
			Platform:      r.Context.Platform,
			ClientVersion: r.Context.ClientVersion,
			Token:         r.Context.Token,
			ClientIP:      r.Context.ClientIP,
			RequestID:     r.Context.RequestID,
		}
	}

	return GetFileByFileIDRequest{
		FileId:  r.FileId,
		Context: xcontext,
	}, nil
}

func encodeGetFileByFileIDResponse(ctx context.Context, req interface{}) (interface{}, error) {
	r := req.(grpcbase.Response)
	if r.Data == nil {
		return &pb.Response{
			Message: r.Message,
		}, nil
	}
	//在所有类型中匹配名称相同的消息名称
	resp := r.Data.(GetFileByFileIDResponse)
	pbresp := &pb.GetFileByFileIDResponse{
		CreatorId:  resp.CreatorId,
		Extension:  resp.Extension,
		FileName:   resp.FileName,
		FlyeleId:   resp.FlyeleId,
		FlyeleType: resp.FlyeleType,
		Id:         resp.Id,
		Length:     resp.Length,
		Md5:        resp.Md5,
		Resource:   resp.Resource,
	}
	b, err := stdproto.Marshal(pbresp)
	if err != nil {
		return nil, err
	}
	anyData := &any.Any{
		Value:   b,
		TypeUrl: stdproto.MessageName(pbresp),
	}
	return &pb.Response{
		Message: r.Message,
		Total:   int32(r.Total),
		Data:    anyData,
	}, nil
}

func makeGetProjectFileTotalEndpoint(svc Repository) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetProjectFileTotalRequest)
		v := svc.GetProjectFileTotal(ctx, req)
		return v, nil
	}
}
func decodeGetProjectFileTotalRequest(ctx context.Context, req interface{}) (interface{}, error) {
	r := req.(*pb.GetProjectFileTotalRequest)
	var xcontext *UserContext
	if r.Context != nil {
		xcontext = &UserContext{
			UserID:        r.Context.UserID,
			Platform:      r.Context.Platform,
			ClientVersion: r.Context.ClientVersion,
			Token:         r.Context.Token,
			ClientIP:      r.Context.ClientIP,
			RequestID:     r.Context.RequestID,
		}
	}

	return GetProjectFileTotalRequest{
		ProjectID: r.ProjectID,
		Context:   xcontext,
	}, nil
}

func encodeGetProjectFileTotalResponse(ctx context.Context, req interface{}) (interface{}, error) {
	r := req.(grpcbase.Response)
	if r.Data == nil {
		return &pb.Response{
			Message: r.Message,
		}, nil
	}
	//在所有类型中匹配名称相同的消息名称
	resp := r.Data.(GetProjectFileTotalResponse)
	pbresp := &pb.GetProjectFileTotalResponse{
		Total: resp.Total,
	}
	b, err := stdproto.Marshal(pbresp)
	if err != nil {
		return nil, err
	}
	anyData := &any.Any{
		Value:   b,
		TypeUrl: stdproto.MessageName(pbresp),
	}
	return &pb.Response{
		Message: r.Message,
		Total:   int32(r.Total),
		Data:    anyData,
	}, nil
}

func makeGetFilesByFileIDsEndpoint(svc Repository) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetFilesByFileIDsRequest)
		v := svc.GetFilesByFileIDs(ctx, req)
		return v, nil
	}
}
func decodeGetFilesByFileIDsRequest(ctx context.Context, req interface{}) (interface{}, error) {
	r := req.(*pb.GetFilesByFileIDsRequest)
	var xcontext *UserContext
	if r.Context != nil {
		xcontext = &UserContext{
			UserID:        r.Context.UserID,
			Platform:      r.Context.Platform,
			ClientVersion: r.Context.ClientVersion,
			Token:         r.Context.Token,
			ClientIP:      r.Context.ClientIP,
			RequestID:     r.Context.RequestID,
		}
	}

	return GetFilesByFileIDsRequest{
		FileId:  r.FileId,
		Context: xcontext,
	}, nil
}

func encodeGetFilesByFileIDsResponse(ctx context.Context, req interface{}) (interface{}, error) {
	r := req.(grpcbase.Response)
	if r.Data == nil {
		return &pb.Response{
			Message: r.Message,
		}, nil
	}
	//在所有类型中匹配名称相同的消息名称
	resp := r.Data.(GetFilesByFileIDsResponse)
	//是复杂类型的数组
	var filesArray = make([]*pb.GetFileByFileIDResponse, 0)

	for _, v := range resp.Files {
		filesArray = append(filesArray, &pb.GetFileByFileIDResponse{
			Id:         v.Id,
			FileName:   v.FileName,
			Resource:   v.Resource,
			Length:     v.Length,
			Md5:        v.Md5,
			CreatorId:  v.CreatorId,
			Extension:  v.Extension,
			FlyeleType: v.FlyeleType,
			FlyeleId:   v.FlyeleId,
		})
	}
	pbresp := &pb.GetFilesByFileIDsResponse{
		Files: filesArray,
	}
	b, err := stdproto.Marshal(pbresp)
	if err != nil {
		return nil, err
	}
	anyData := &any.Any{
		Value:   b,
		TypeUrl: stdproto.MessageName(pbresp),
	}
	return &pb.Response{
		Message: r.Message,
		Total:   int32(r.Total),
		Data:    anyData,
	}, nil
}

func makeGetUsersFileUsedCapEndpoint(svc Repository) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetUsersFileUsedCapRequest)
		v := svc.GetUsersFileUsedCap(ctx, req)
		return v, nil
	}
}
func decodeGetUsersFileUsedCapRequest(ctx context.Context, req interface{}) (interface{}, error) {
	r := req.(*pb.GetUsersFileUsedCapRequest)
	var xcontext *UserContext
	if r.Context != nil {
		xcontext = &UserContext{
			UserID:        r.Context.UserID,
			Platform:      r.Context.Platform,
			ClientVersion: r.Context.ClientVersion,
			Token:         r.Context.Token,
			ClientIP:      r.Context.ClientIP,
			RequestID:     r.Context.RequestID,
		}
	}

	return GetUsersFileUsedCapRequest{
		UserID:      r.UserID,
		WorkspaceID: r.WorkspaceID,
		Context:     xcontext,
	}, nil
}

func encodeGetUsersFileUsedCapResponse(ctx context.Context, req interface{}) (interface{}, error) {
	r := req.(grpcbase.Response)
	if r.Data == nil {
		return &pb.Response{
			Message: r.Message,
		}, nil
	}
	//在所有类型中匹配名称相同的消息名称
	resp := r.Data.(GetUsersFileUsedCapResponse)
	pbresp := &pb.GetUsersFileUsedCapResponse{
		AllCap: resp.AllCap,
	}
	b, err := stdproto.Marshal(pbresp)
	if err != nil {
		return nil, err
	}
	anyData := &any.Any{
		Value:   b,
		TypeUrl: stdproto.MessageName(pbresp),
	}
	return &pb.Response{
		Message: r.Message,
		Total:   int32(r.Total),
		Data:    anyData,
	}, nil
}
