// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by grpc_generate_tools(grpcgengo) at
//

package clouddisk

import (
	"context"
	"log"

	"github.com/go-kit/kit/endpoint"
	"gitlab.flyele.vip/server-side/go-zero/v2/rpcx/grpcbase"
	"google.golang.org/grpc"
)

func init() {
	grpcbase.RegisterClients(grpcbase.ServerAddr(grpcbase.CloudDiskSVC), &clientBinding{})
}

type clientBinding struct {
	fileDetail               endpoint.Endpoint
	associatedFiles          endpoint.Endpoint
	uploadFile               endpoint.Endpoint
	getFilesName             endpoint.Endpoint
	batchUnbindFile          endpoint.Endpoint
	batchBindFile            endpoint.Endpoint
	fileSync                 endpoint.Endpoint
	taskFiles                endpoint.Endpoint
	getBindLogWithAttachment endpoint.Endpoint
	getFileAddrByFileID      endpoint.Endpoint
	getBindFileSize          endpoint.Endpoint
	getFileByFileID          endpoint.Endpoint
	getProjectFileTotal      endpoint.Endpoint
	getFilesByFileIDs        endpoint.Endpoint
	getUsersFileUsedCap      endpoint.Endpoint
}

func (c *clientBinding) FileDetail(ctx context.Context, params FileDetailRequest) grpcbase.Response {
	if ctx == nil {
		ctx = context.Background()
		log.Println("GRPC：FileDetail request context is nil，trace span将无法生效")
	}
	response, err := c.fileDetail(ctx, params)
	if err != nil {
		return grpcbase.Response{
			Message: err.Error(),
		}
	}
	r := response.(grpcbase.Response)
	return r
}
func (c *clientBinding) AssociatedFiles(ctx context.Context, params AssociatedFilesRequest) grpcbase.Response {
	if ctx == nil {
		ctx = context.Background()
		log.Println("GRPC：AssociatedFiles request context is nil，trace span将无法生效")
	}
	response, err := c.associatedFiles(ctx, params)
	if err != nil {
		return grpcbase.Response{
			Message: err.Error(),
		}
	}
	r := response.(grpcbase.Response)
	return r
}
func (c *clientBinding) UploadFile(ctx context.Context, params UploadFileRequest) grpcbase.Response {
	if ctx == nil {
		ctx = context.Background()
		log.Println("GRPC：UploadFile request context is nil，trace span将无法生效")
	}
	response, err := c.uploadFile(ctx, params)
	if err != nil {
		return grpcbase.Response{
			Message: err.Error(),
		}
	}
	r := response.(grpcbase.Response)
	return r
}
func (c *clientBinding) GetFilesName(ctx context.Context, params GetFilesNameRequest) grpcbase.Response {
	if ctx == nil {
		ctx = context.Background()
		log.Println("GRPC：GetFilesName request context is nil，trace span将无法生效")
	}
	response, err := c.getFilesName(ctx, params)
	if err != nil {
		return grpcbase.Response{
			Message: err.Error(),
		}
	}
	r := response.(grpcbase.Response)
	return r
}
func (c *clientBinding) BatchUnbindFile(ctx context.Context, params BatchUnbindFileRequest) grpcbase.Response {
	if ctx == nil {
		ctx = context.Background()
		log.Println("GRPC：BatchUnbindFile request context is nil，trace span将无法生效")
	}
	response, err := c.batchUnbindFile(ctx, params)
	if err != nil {
		return grpcbase.Response{
			Message: err.Error(),
		}
	}
	r := response.(grpcbase.Response)
	return r
}
func (c *clientBinding) BatchBindFile(ctx context.Context, params BatchBindFileRequest) grpcbase.Response {
	if ctx == nil {
		ctx = context.Background()
		log.Println("GRPC：BatchBindFile request context is nil，trace span将无法生效")
	}
	response, err := c.batchBindFile(ctx, params)
	if err != nil {
		return grpcbase.Response{
			Message: err.Error(),
		}
	}
	r := response.(grpcbase.Response)
	return r
}
func (c *clientBinding) FileSync(ctx context.Context, params FileSyncRequest) grpcbase.Response {
	if ctx == nil {
		ctx = context.Background()
		log.Println("GRPC：FileSync request context is nil，trace span将无法生效")
	}
	response, err := c.fileSync(ctx, params)
	if err != nil {
		return grpcbase.Response{
			Message: err.Error(),
		}
	}
	r := response.(grpcbase.Response)
	return r
}
func (c *clientBinding) TaskFiles(ctx context.Context, params TaskFilesRequest) grpcbase.Response {
	if ctx == nil {
		ctx = context.Background()
		log.Println("GRPC：TaskFiles request context is nil，trace span将无法生效")
	}
	response, err := c.taskFiles(ctx, params)
	if err != nil {
		return grpcbase.Response{
			Message: err.Error(),
		}
	}
	r := response.(grpcbase.Response)
	return r
}
func (c *clientBinding) GetBindLogWithAttachment(ctx context.Context, params GetBindLogWithAttachmentRequest) grpcbase.Response {
	if ctx == nil {
		ctx = context.Background()
		log.Println("GRPC：GetBindLogWithAttachment request context is nil，trace span将无法生效")
	}
	response, err := c.getBindLogWithAttachment(ctx, params)
	if err != nil {
		return grpcbase.Response{
			Message: err.Error(),
		}
	}
	r := response.(grpcbase.Response)
	return r
}
func (c *clientBinding) GetFileAddrByFileID(ctx context.Context, params GetFileAddrByFileIDRequest) grpcbase.Response {
	if ctx == nil {
		ctx = context.Background()
		log.Println("GRPC：GetFileAddrByFileID request context is nil，trace span将无法生效")
	}
	response, err := c.getFileAddrByFileID(ctx, params)
	if err != nil {
		return grpcbase.Response{
			Message: err.Error(),
		}
	}
	r := response.(grpcbase.Response)
	return r
}
func (c *clientBinding) GetBindFileSize(ctx context.Context, params GetBindFileSizeRequest) grpcbase.Response {
	if ctx == nil {
		ctx = context.Background()
		log.Println("GRPC：GetBindFileSize request context is nil，trace span将无法生效")
	}
	response, err := c.getBindFileSize(ctx, params)
	if err != nil {
		return grpcbase.Response{
			Message: err.Error(),
		}
	}
	r := response.(grpcbase.Response)
	return r
}
func (c *clientBinding) GetFileByFileID(ctx context.Context, params GetFileByFileIDRequest) grpcbase.Response {
	if ctx == nil {
		ctx = context.Background()
		log.Println("GRPC：GetFileByFileID request context is nil，trace span将无法生效")
	}
	response, err := c.getFileByFileID(ctx, params)
	if err != nil {
		return grpcbase.Response{
			Message: err.Error(),
		}
	}
	r := response.(grpcbase.Response)
	return r
}
func (c *clientBinding) GetProjectFileTotal(ctx context.Context, params GetProjectFileTotalRequest) grpcbase.Response {
	if ctx == nil {
		ctx = context.Background()
		log.Println("GRPC：GetProjectFileTotal request context is nil，trace span将无法生效")
	}
	response, err := c.getProjectFileTotal(ctx, params)
	if err != nil {
		return grpcbase.Response{
			Message: err.Error(),
		}
	}
	r := response.(grpcbase.Response)
	return r
}
func (c *clientBinding) GetFilesByFileIDs(ctx context.Context, params GetFilesByFileIDsRequest) grpcbase.Response {
	if ctx == nil {
		ctx = context.Background()
		log.Println("GRPC：GetFilesByFileIDs request context is nil，trace span将无法生效")
	}
	response, err := c.getFilesByFileIDs(ctx, params)
	if err != nil {
		return grpcbase.Response{
			Message: err.Error(),
		}
	}
	r := response.(grpcbase.Response)
	return r
}
func (c *clientBinding) GetUsersFileUsedCap(ctx context.Context, params GetUsersFileUsedCapRequest) grpcbase.Response {
	if ctx == nil {
		ctx = context.Background()
		log.Println("GRPC：GetUsersFileUsedCap request context is nil，trace span将无法生效")
	}
	response, err := c.getUsersFileUsedCap(ctx, params)
	if err != nil {
		return grpcbase.Response{
			Message: err.Error(),
		}
	}
	r := response.(grpcbase.Response)
	return r
}

func (c *clientBinding) GRPCClient(cc *grpc.ClientConn) interface{} {
	c.newClient(cc)
	return c
}

func (c *clientBinding) newClient(cc *grpc.ClientConn) {

	c.fileDetail = grpcbase.CreateGRPCClientEndpoint(cc, "pb.CloudDisk",
		"FileDetail",
		encodeFileDetailRequest,
		decodeFileDetailResponse)
	c.associatedFiles = grpcbase.CreateGRPCClientEndpoint(cc, "pb.CloudDisk",
		"AssociatedFiles",
		encodeAssociatedFilesRequest,
		decodeAssociatedFilesResponse)
	c.uploadFile = grpcbase.CreateGRPCClientEndpoint(cc, "pb.CloudDisk",
		"UploadFile",
		encodeUploadFileRequest,
		decodeUploadFileResponse)
	c.getFilesName = grpcbase.CreateGRPCClientEndpoint(cc, "pb.CloudDisk",
		"GetFilesName",
		encodeGetFilesNameRequest,
		decodeGetFilesNameResponse)
	c.batchUnbindFile = grpcbase.CreateGRPCClientEndpoint(cc, "pb.CloudDisk",
		"BatchUnbindFile",
		encodeBatchUnbindFileRequest,
		decodeBatchUnbindFileResponse)
	c.batchBindFile = grpcbase.CreateGRPCClientEndpoint(cc, "pb.CloudDisk",
		"BatchBindFile",
		encodeBatchBindFileRequest,
		decodeBatchBindFileResponse)
	c.fileSync = grpcbase.CreateGRPCClientEndpoint(cc, "pb.CloudDisk",
		"FileSync",
		encodeFileSyncRequest,
		decodeFileSyncResponse)
	c.taskFiles = grpcbase.CreateGRPCClientEndpoint(cc, "pb.CloudDisk",
		"TaskFiles",
		encodeTaskFilesRequest,
		decodeTaskFilesResponse)
	c.getBindLogWithAttachment = grpcbase.CreateGRPCClientEndpoint(cc, "pb.CloudDisk",
		"GetBindLogWithAttachment",
		encodeGetBindLogWithAttachmentRequest,
		decodeGetBindLogWithAttachmentResponse)
	c.getFileAddrByFileID = grpcbase.CreateGRPCClientEndpoint(cc, "pb.CloudDisk",
		"GetFileAddrByFileID",
		encodeGetFileAddrByFileIDRequest,
		decodeGetFileAddrByFileIDResponse)
	c.getBindFileSize = grpcbase.CreateGRPCClientEndpoint(cc, "pb.CloudDisk",
		"GetBindFileSize",
		encodeGetBindFileSizeRequest,
		decodeGetBindFileSizeResponse)
	c.getFileByFileID = grpcbase.CreateGRPCClientEndpoint(cc, "pb.CloudDisk",
		"GetFileByFileID",
		encodeGetFileByFileIDRequest,
		decodeGetFileByFileIDResponse)
	c.getProjectFileTotal = grpcbase.CreateGRPCClientEndpoint(cc, "pb.CloudDisk",
		"GetProjectFileTotal",
		encodeGetProjectFileTotalRequest,
		decodeGetProjectFileTotalResponse)
	c.getFilesByFileIDs = grpcbase.CreateGRPCClientEndpoint(cc, "pb.CloudDisk",
		"GetFilesByFileIDs",
		encodeGetFilesByFileIDsRequest,
		decodeGetFilesByFileIDsResponse)
	c.getUsersFileUsedCap = grpcbase.CreateGRPCClientEndpoint(cc, "pb.CloudDisk",
		"GetUsersFileUsedCap",
		encodeGetUsersFileUsedCapRequest,
		decodeGetUsersFileUsedCapResponse)
}
