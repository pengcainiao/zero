// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by grpc_generate_tools(grpcgengo) at
//

package clouddisk

import (
	"context"
	"log"

	grpctransport "github.com/go-kit/kit/transport/grpc"
	"gitlab.flyele.vip/server-side/go-zero/v2/rpcx/grpcbase"
	pb "gitlab.flyele.vip/server-side/go-zero/v2/rpcx/protos"
	"google.golang.org/grpc"
)

type serverBinding struct {
	fileDetail               grpctransport.Handler
	associatedFiles          grpctransport.Handler
	uploadFile               grpctransport.Handler
	getFilesName             grpctransport.Handler
	batchUnbindFile          grpctransport.Handler
	batchBindFile            grpctransport.Handler
	fileSync                 grpctransport.Handler
	taskFiles                grpctransport.Handler
	getBindLogWithAttachment grpctransport.Handler
	getFileAddrByFileID      grpctransport.Handler
	getBindFileSize          grpctransport.Handler
	getFileByFileID          grpctransport.Handler
	getProjectFileTotal      grpctransport.Handler
	getFilesByFileIDs        grpctransport.Handler
	getUsersFileUsedCap      grpctransport.Handler
}

func (b *serverBinding) FileDetail(ctx context.Context, req *pb.FileDetailRequest) (*pb.Response, error) {
	if ctx == nil {
		ctx = context.Background()
		log.Println("GRPC：FileDetail receive request context is nil，trace span将无法生效")
	}
	_, response, err := b.fileDetail.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return response.(*pb.Response), nil
}
func (b *serverBinding) AssociatedFiles(ctx context.Context, req *pb.AssociatedFilesRequest) (*pb.Response, error) {
	if ctx == nil {
		ctx = context.Background()
		log.Println("GRPC：AssociatedFiles receive request context is nil，trace span将无法生效")
	}
	_, response, err := b.associatedFiles.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return response.(*pb.Response), nil
}
func (b *serverBinding) UploadFile(ctx context.Context, req *pb.UploadFileRequest) (*pb.Response, error) {
	if ctx == nil {
		ctx = context.Background()
		log.Println("GRPC：UploadFile receive request context is nil，trace span将无法生效")
	}
	_, response, err := b.uploadFile.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return response.(*pb.Response), nil
}
func (b *serverBinding) GetFilesName(ctx context.Context, req *pb.GetFilesNameRequest) (*pb.Response, error) {
	if ctx == nil {
		ctx = context.Background()
		log.Println("GRPC：GetFilesName receive request context is nil，trace span将无法生效")
	}
	_, response, err := b.getFilesName.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return response.(*pb.Response), nil
}
func (b *serverBinding) BatchUnbindFile(ctx context.Context, req *pb.BatchUnbindFileRequest) (*pb.Response, error) {
	if ctx == nil {
		ctx = context.Background()
		log.Println("GRPC：BatchUnbindFile receive request context is nil，trace span将无法生效")
	}
	_, response, err := b.batchUnbindFile.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return response.(*pb.Response), nil
}
func (b *serverBinding) BatchBindFile(ctx context.Context, req *pb.BatchBindFileRequest) (*pb.Response, error) {
	if ctx == nil {
		ctx = context.Background()
		log.Println("GRPC：BatchBindFile receive request context is nil，trace span将无法生效")
	}
	_, response, err := b.batchBindFile.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return response.(*pb.Response), nil
}
func (b *serverBinding) FileSync(ctx context.Context, req *pb.FileSyncRequest) (*pb.Response, error) {
	if ctx == nil {
		ctx = context.Background()
		log.Println("GRPC：FileSync receive request context is nil，trace span将无法生效")
	}
	_, response, err := b.fileSync.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return response.(*pb.Response), nil
}
func (b *serverBinding) TaskFiles(ctx context.Context, req *pb.TaskFilesRequest) (*pb.Response, error) {
	if ctx == nil {
		ctx = context.Background()
		log.Println("GRPC：TaskFiles receive request context is nil，trace span将无法生效")
	}
	_, response, err := b.taskFiles.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return response.(*pb.Response), nil
}
func (b *serverBinding) GetBindLogWithAttachment(ctx context.Context, req *pb.GetBindLogWithAttachmentRequest) (*pb.Response, error) {
	if ctx == nil {
		ctx = context.Background()
		log.Println("GRPC：GetBindLogWithAttachment receive request context is nil，trace span将无法生效")
	}
	_, response, err := b.getBindLogWithAttachment.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return response.(*pb.Response), nil
}
func (b *serverBinding) GetFileAddrByFileID(ctx context.Context, req *pb.GetFileAddrByFileIDRequest) (*pb.Response, error) {
	if ctx == nil {
		ctx = context.Background()
		log.Println("GRPC：GetFileAddrByFileID receive request context is nil，trace span将无法生效")
	}
	_, response, err := b.getFileAddrByFileID.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return response.(*pb.Response), nil
}
func (b *serverBinding) GetBindFileSize(ctx context.Context, req *pb.GetBindFileSizeRequest) (*pb.Response, error) {
	if ctx == nil {
		ctx = context.Background()
		log.Println("GRPC：GetBindFileSize receive request context is nil，trace span将无法生效")
	}
	_, response, err := b.getBindFileSize.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return response.(*pb.Response), nil
}
func (b *serverBinding) GetFileByFileID(ctx context.Context, req *pb.GetFileByFileIDRequest) (*pb.Response, error) {
	if ctx == nil {
		ctx = context.Background()
		log.Println("GRPC：GetFileByFileID receive request context is nil，trace span将无法生效")
	}
	_, response, err := b.getFileByFileID.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return response.(*pb.Response), nil
}
func (b *serverBinding) GetProjectFileTotal(ctx context.Context, req *pb.GetProjectFileTotalRequest) (*pb.Response, error) {
	if ctx == nil {
		ctx = context.Background()
		log.Println("GRPC：GetProjectFileTotal receive request context is nil，trace span将无法生效")
	}
	_, response, err := b.getProjectFileTotal.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return response.(*pb.Response), nil
}
func (b *serverBinding) GetFilesByFileIDs(ctx context.Context, req *pb.GetFilesByFileIDsRequest) (*pb.Response, error) {
	if ctx == nil {
		ctx = context.Background()
		log.Println("GRPC：GetFilesByFileIDs receive request context is nil，trace span将无法生效")
	}
	_, response, err := b.getFilesByFileIDs.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return response.(*pb.Response), nil
}
func (b *serverBinding) GetUsersFileUsedCap(ctx context.Context, req *pb.GetUsersFileUsedCapRequest) (*pb.Response, error) {
	if ctx == nil {
		ctx = context.Background()
		log.Println("GRPC：GetUsersFileUsedCap receive request context is nil，trace span将无法生效")
	}
	_, response, err := b.getUsersFileUsedCap.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return response.(*pb.Response), nil
}

func (b *serverBinding) RegisterServer(srv *grpc.Server) error {
	pb.RegisterCloudDiskServer(srv, b)
	return nil
}

func (b *serverBinding) GRPCHandler() map[string]grpctransport.Handler {
	return map[string]grpctransport.Handler{
		"fileDetail":               b.fileDetail,
		"associatedFiles":          b.associatedFiles,
		"uploadFile":               b.uploadFile,
		"getFilesName":             b.getFilesName,
		"batchUnbindFile":          b.batchUnbindFile,
		"batchBindFile":            b.batchBindFile,
		"fileSync":                 b.fileSync,
		"taskFiles":                b.taskFiles,
		"getBindLogWithAttachment": b.getBindLogWithAttachment,
		"getFileAddrByFileID":      b.getFileAddrByFileID,
		"getBindFileSize":          b.getBindFileSize,
		"getFileByFileID":          b.getFileByFileID,
		"getProjectFileTotal":      b.getProjectFileTotal,
		"getFilesByFileIDs":        b.getFilesByFileIDs,
		"getUsersFileUsedCap":      b.getUsersFileUsedCap,
	}
}

func NewBinding(svc Repository) *serverBinding {
	return &serverBinding{
		fileDetail: grpcbase.CreateGRPCServer(
			makeFileDetailEndpoint(svc),
			decodeFileDetailRequest,
			encodeFileDetailResponse,
		),
		associatedFiles: grpcbase.CreateGRPCServer(
			makeAssociatedFilesEndpoint(svc),
			decodeAssociatedFilesRequest,
			encodeAssociatedFilesResponse,
		),
		uploadFile: grpcbase.CreateGRPCServer(
			makeUploadFileEndpoint(svc),
			decodeUploadFileRequest,
			encodeUploadFileResponse,
		),
		getFilesName: grpcbase.CreateGRPCServer(
			makeGetFilesNameEndpoint(svc),
			decodeGetFilesNameRequest,
			encodeGetFilesNameResponse,
		),
		batchUnbindFile: grpcbase.CreateGRPCServer(
			makeBatchUnbindFileEndpoint(svc),
			decodeBatchUnbindFileRequest,
			encodeBatchUnbindFileResponse,
		),
		batchBindFile: grpcbase.CreateGRPCServer(
			makeBatchBindFileEndpoint(svc),
			decodeBatchBindFileRequest,
			encodeBatchBindFileResponse,
		),
		fileSync: grpcbase.CreateGRPCServer(
			makeFileSyncEndpoint(svc),
			decodeFileSyncRequest,
			encodeFileSyncResponse,
		),
		taskFiles: grpcbase.CreateGRPCServer(
			makeTaskFilesEndpoint(svc),
			decodeTaskFilesRequest,
			encodeTaskFilesResponse,
		),
		getBindLogWithAttachment: grpcbase.CreateGRPCServer(
			makeGetBindLogWithAttachmentEndpoint(svc),
			decodeGetBindLogWithAttachmentRequest,
			encodeGetBindLogWithAttachmentResponse,
		),
		getFileAddrByFileID: grpcbase.CreateGRPCServer(
			makeGetFileAddrByFileIDEndpoint(svc),
			decodeGetFileAddrByFileIDRequest,
			encodeGetFileAddrByFileIDResponse,
		),
		getBindFileSize: grpcbase.CreateGRPCServer(
			makeGetBindFileSizeEndpoint(svc),
			decodeGetBindFileSizeRequest,
			encodeGetBindFileSizeResponse,
		),
		getFileByFileID: grpcbase.CreateGRPCServer(
			makeGetFileByFileIDEndpoint(svc),
			decodeGetFileByFileIDRequest,
			encodeGetFileByFileIDResponse,
		),
		getProjectFileTotal: grpcbase.CreateGRPCServer(
			makeGetProjectFileTotalEndpoint(svc),
			decodeGetProjectFileTotalRequest,
			encodeGetProjectFileTotalResponse,
		),
		getFilesByFileIDs: grpcbase.CreateGRPCServer(
			makeGetFilesByFileIDsEndpoint(svc),
			decodeGetFilesByFileIDsRequest,
			encodeGetFilesByFileIDsResponse,
		),
		getUsersFileUsedCap: grpcbase.CreateGRPCServer(
			makeGetUsersFileUsedCapEndpoint(svc),
			decodeGetUsersFileUsedCapRequest,
			encodeGetUsersFileUsedCapResponse,
		),
	}
}
