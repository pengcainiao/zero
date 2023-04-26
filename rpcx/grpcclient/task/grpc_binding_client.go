// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by grpc_generate_tools(grpcgengo) at
//

package task

import (
	"context"
	"log"

	"github.com/go-kit/kit/endpoint"
	"gitlab.flyele.vip/server-side/go-zero/v2/rpcx/grpcbase"
	"google.golang.org/grpc"
)

func init() {
	grpcbase.RegisterClients(grpcbase.ServerAddr(grpcbase.TaskSVC), &clientBinding{})
}

type clientBinding struct {
	getTask                          endpoint.Endpoint
	getTasksName                     endpoint.Endpoint
	getTaskDispatch                  endpoint.Endpoint
	getTaskTakers                    endpoint.Endpoint
	getPureTaskTakers                endpoint.Endpoint
	taskExists                       endpoint.Endpoint
	createGuideTask                  endpoint.Endpoint
	batchQueryTask                   endpoint.Endpoint
	getTaskRelevantTakers            endpoint.Endpoint
	updateTaskRelation               endpoint.Endpoint
	getTaskShare                     endpoint.Endpoint
	createUserInteract               endpoint.Endpoint
	updateUserInteract               endpoint.Endpoint
	createTaskActiveDetail           endpoint.Endpoint
	updateProjectRelation            endpoint.Endpoint
	getProjectMembers                endpoint.Endpoint
	getProjectsName                  endpoint.Endpoint
	getUsersMorningAndEveningProgram endpoint.Endpoint
	getTaskWithParent                endpoint.Endpoint
	getProject                       endpoint.Endpoint
	getWorkspace                     endpoint.Endpoint
	getWorkspaceMembers              endpoint.Endpoint
	getWorkspaceList                 endpoint.Endpoint
	getTaskChildren                  endpoint.Endpoint
	getWorkspaceFileSpace            endpoint.Endpoint
	getWorkspaceBindProject          endpoint.Endpoint
	getWorkspaceExternalMember       endpoint.Endpoint
	expireEquityUpdate               endpoint.Endpoint
	updateWorkspaceLevel             endpoint.Endpoint
	getTaskWorkspace                 endpoint.Endpoint
	getProjectWorkspace              endpoint.Endpoint
	addMemberEquityUpdate            endpoint.Endpoint
	batchGetWorkspaceDirectory       endpoint.Endpoint
	getWorkspaceMembersOfTask        endpoint.Endpoint
	getBatchWorkspaceInfo            endpoint.Endpoint
	createNewInvite                  endpoint.Endpoint
	getObjectiveMember               endpoint.Endpoint
	getObjective                     endpoint.Endpoint
	getScreen                        endpoint.Endpoint
	getScreenMembers                 endpoint.Endpoint
	getScreensByCardType             endpoint.Endpoint
}

func (c *clientBinding) GetTask(ctx context.Context, params GetTaskRequest) grpcbase.Response {
	if ctx == nil {
		ctx = context.Background()
		log.Println("GRPC：GetTask request context is nil，trace span将无法生效")
	}
	response, err := c.getTask(ctx, params)
	if err != nil {
		return grpcbase.Response{
			Message: err.Error(),
		}
	}
	r := response.(grpcbase.Response)
	return r
}
func (c *clientBinding) GetTasksName(ctx context.Context, params GetTasksNameRequest) grpcbase.Response {
	if ctx == nil {
		ctx = context.Background()
		log.Println("GRPC：GetTasksName request context is nil，trace span将无法生效")
	}
	response, err := c.getTasksName(ctx, params)
	if err != nil {
		return grpcbase.Response{
			Message: err.Error(),
		}
	}
	r := response.(grpcbase.Response)
	return r
}
func (c *clientBinding) GetTaskDispatch(ctx context.Context, params GetTaskDispatchRequest) grpcbase.Response {
	if ctx == nil {
		ctx = context.Background()
		log.Println("GRPC：GetTaskDispatch request context is nil，trace span将无法生效")
	}
	response, err := c.getTaskDispatch(ctx, params)
	if err != nil {
		return grpcbase.Response{
			Message: err.Error(),
		}
	}
	r := response.(grpcbase.Response)
	return r
}
func (c *clientBinding) GetTaskTakers(ctx context.Context, params GetTaskTakersRequest) grpcbase.Response {
	if ctx == nil {
		ctx = context.Background()
		log.Println("GRPC：GetTaskTakers request context is nil，trace span将无法生效")
	}
	response, err := c.getTaskTakers(ctx, params)
	if err != nil {
		return grpcbase.Response{
			Message: err.Error(),
		}
	}
	r := response.(grpcbase.Response)
	return r
}
func (c *clientBinding) GetPureTaskTakers(ctx context.Context, params GetPureTaskTakersRequest) grpcbase.Response {
	if ctx == nil {
		ctx = context.Background()
		log.Println("GRPC：GetPureTaskTakers request context is nil，trace span将无法生效")
	}
	response, err := c.getPureTaskTakers(ctx, params)
	if err != nil {
		return grpcbase.Response{
			Message: err.Error(),
		}
	}
	r := response.(grpcbase.Response)
	return r
}
func (c *clientBinding) TaskExists(ctx context.Context, params TaskExistsRequest) grpcbase.Response {
	if ctx == nil {
		ctx = context.Background()
		log.Println("GRPC：TaskExists request context is nil，trace span将无法生效")
	}
	response, err := c.taskExists(ctx, params)
	if err != nil {
		return grpcbase.Response{
			Message: err.Error(),
		}
	}
	r := response.(grpcbase.Response)
	return r
}
func (c *clientBinding) CreateGuideTask(ctx context.Context, params CreateGuideTaskRequest) grpcbase.Response {
	if ctx == nil {
		ctx = context.Background()
		log.Println("GRPC：CreateGuideTask request context is nil，trace span将无法生效")
	}
	response, err := c.createGuideTask(ctx, params)
	if err != nil {
		return grpcbase.Response{
			Message: err.Error(),
		}
	}
	r := response.(grpcbase.Response)
	return r
}
func (c *clientBinding) BatchQueryTask(ctx context.Context, params BatchQueryTaskRequest) grpcbase.Response {
	if ctx == nil {
		ctx = context.Background()
		log.Println("GRPC：BatchQueryTask request context is nil，trace span将无法生效")
	}
	response, err := c.batchQueryTask(ctx, params)
	if err != nil {
		return grpcbase.Response{
			Message: err.Error(),
		}
	}
	r := response.(grpcbase.Response)
	return r
}
func (c *clientBinding) GetTaskRelevantTakers(ctx context.Context, params GetTaskRelevantTakersRequest) grpcbase.Response {
	if ctx == nil {
		ctx = context.Background()
		log.Println("GRPC：GetTaskRelevantTakers request context is nil，trace span将无法生效")
	}
	response, err := c.getTaskRelevantTakers(ctx, params)
	if err != nil {
		return grpcbase.Response{
			Message: err.Error(),
		}
	}
	r := response.(grpcbase.Response)
	return r
}
func (c *clientBinding) UpdateTaskRelation(ctx context.Context, params UpdateTaskRelationRequest) grpcbase.Response {
	if ctx == nil {
		ctx = context.Background()
		log.Println("GRPC：UpdateTaskRelation request context is nil，trace span将无法生效")
	}
	response, err := c.updateTaskRelation(ctx, params)
	if err != nil {
		return grpcbase.Response{
			Message: err.Error(),
		}
	}
	r := response.(grpcbase.Response)
	return r
}
func (c *clientBinding) GetTaskShare(ctx context.Context, params GetTaskShareRequest) grpcbase.Response {
	if ctx == nil {
		ctx = context.Background()
		log.Println("GRPC：GetTaskShare request context is nil，trace span将无法生效")
	}
	response, err := c.getTaskShare(ctx, params)
	if err != nil {
		return grpcbase.Response{
			Message: err.Error(),
		}
	}
	r := response.(grpcbase.Response)
	return r
}
func (c *clientBinding) CreateUserInteract(ctx context.Context, params CreateUserInteractRequest) grpcbase.Response {
	if ctx == nil {
		ctx = context.Background()
		log.Println("GRPC：CreateUserInteract request context is nil，trace span将无法生效")
	}
	response, err := c.createUserInteract(ctx, params)
	if err != nil {
		return grpcbase.Response{
			Message: err.Error(),
		}
	}
	r := response.(grpcbase.Response)
	return r
}
func (c *clientBinding) UpdateUserInteract(ctx context.Context, params UpdateUserInteractRequest) grpcbase.Response {
	if ctx == nil {
		ctx = context.Background()
		log.Println("GRPC：UpdateUserInteract request context is nil，trace span将无法生效")
	}
	response, err := c.updateUserInteract(ctx, params)
	if err != nil {
		return grpcbase.Response{
			Message: err.Error(),
		}
	}
	r := response.(grpcbase.Response)
	return r
}
func (c *clientBinding) CreateTaskActiveDetail(ctx context.Context, params CreateTaskActiveDetailRequest) grpcbase.Response {
	if ctx == nil {
		ctx = context.Background()
		log.Println("GRPC：CreateTaskActiveDetail request context is nil，trace span将无法生效")
	}
	response, err := c.createTaskActiveDetail(ctx, params)
	if err != nil {
		return grpcbase.Response{
			Message: err.Error(),
		}
	}
	r := response.(grpcbase.Response)
	return r
}
func (c *clientBinding) UpdateProjectRelation(ctx context.Context, params UpdateProjectRelationRequest) grpcbase.Response {
	if ctx == nil {
		ctx = context.Background()
		log.Println("GRPC：UpdateProjectRelation request context is nil，trace span将无法生效")
	}
	response, err := c.updateProjectRelation(ctx, params)
	if err != nil {
		return grpcbase.Response{
			Message: err.Error(),
		}
	}
	r := response.(grpcbase.Response)
	return r
}
func (c *clientBinding) GetProjectMembers(ctx context.Context, params GetProjectMembersRequest) grpcbase.Response {
	if ctx == nil {
		ctx = context.Background()
		log.Println("GRPC：GetProjectMembers request context is nil，trace span将无法生效")
	}
	response, err := c.getProjectMembers(ctx, params)
	if err != nil {
		return grpcbase.Response{
			Message: err.Error(),
		}
	}
	r := response.(grpcbase.Response)
	return r
}
func (c *clientBinding) GetProjectsName(ctx context.Context, params GetProjectsNameRequest) grpcbase.Response {
	if ctx == nil {
		ctx = context.Background()
		log.Println("GRPC：GetProjectsName request context is nil，trace span将无法生效")
	}
	response, err := c.getProjectsName(ctx, params)
	if err != nil {
		return grpcbase.Response{
			Message: err.Error(),
		}
	}
	r := response.(grpcbase.Response)
	return r
}
func (c *clientBinding) GetUsersMorningAndEveningProgram(ctx context.Context, params GetUsersMorningAndEveningProgramRequest) grpcbase.Response {
	if ctx == nil {
		ctx = context.Background()
		log.Println("GRPC：GetUsersMorningAndEveningProgram request context is nil，trace span将无法生效")
	}
	response, err := c.getUsersMorningAndEveningProgram(ctx, params)
	if err != nil {
		return grpcbase.Response{
			Message: err.Error(),
		}
	}
	r := response.(grpcbase.Response)
	return r
}
func (c *clientBinding) GetTaskWithParent(ctx context.Context, params GetTaskWithParentRequest) grpcbase.Response {
	if ctx == nil {
		ctx = context.Background()
		log.Println("GRPC：GetTaskWithParent request context is nil，trace span将无法生效")
	}
	response, err := c.getTaskWithParent(ctx, params)
	if err != nil {
		return grpcbase.Response{
			Message: err.Error(),
		}
	}
	r := response.(grpcbase.Response)
	return r
}
func (c *clientBinding) GetProject(ctx context.Context, params GetProjectRequest) grpcbase.Response {
	if ctx == nil {
		ctx = context.Background()
		log.Println("GRPC：GetProject request context is nil，trace span将无法生效")
	}
	response, err := c.getProject(ctx, params)
	if err != nil {
		return grpcbase.Response{
			Message: err.Error(),
		}
	}
	r := response.(grpcbase.Response)
	return r
}
func (c *clientBinding) GetWorkspace(ctx context.Context, params GetWorkspaceRequest) grpcbase.Response {
	if ctx == nil {
		ctx = context.Background()
		log.Println("GRPC：GetWorkspace request context is nil，trace span将无法生效")
	}
	response, err := c.getWorkspace(ctx, params)
	if err != nil {
		return grpcbase.Response{
			Message: err.Error(),
		}
	}
	r := response.(grpcbase.Response)
	return r
}
func (c *clientBinding) GetWorkspaceMembers(ctx context.Context, params GetWorkspaceMembersRequest) grpcbase.Response {
	if ctx == nil {
		ctx = context.Background()
		log.Println("GRPC：GetWorkspaceMembers request context is nil，trace span将无法生效")
	}
	response, err := c.getWorkspaceMembers(ctx, params)
	if err != nil {
		return grpcbase.Response{
			Message: err.Error(),
		}
	}
	r := response.(grpcbase.Response)
	return r
}
func (c *clientBinding) GetWorkspaceList(ctx context.Context, params GetWorkspaceListRequest) grpcbase.Response {
	if ctx == nil {
		ctx = context.Background()
		log.Println("GRPC：GetWorkspaceList request context is nil，trace span将无法生效")
	}
	response, err := c.getWorkspaceList(ctx, params)
	if err != nil {
		return grpcbase.Response{
			Message: err.Error(),
		}
	}
	r := response.(grpcbase.Response)
	return r
}
func (c *clientBinding) GetTaskChildren(ctx context.Context, params GetTaskChildrenRequest) grpcbase.Response {
	if ctx == nil {
		ctx = context.Background()
		log.Println("GRPC：GetTaskChildren request context is nil，trace span将无法生效")
	}
	response, err := c.getTaskChildren(ctx, params)
	if err != nil {
		return grpcbase.Response{
			Message: err.Error(),
		}
	}
	r := response.(grpcbase.Response)
	return r
}
func (c *clientBinding) GetWorkspaceFileSpace(ctx context.Context, params GetWorkspaceFileSpaceRequest) grpcbase.Response {
	if ctx == nil {
		ctx = context.Background()
		log.Println("GRPC：GetWorkspaceFileSpace request context is nil，trace span将无法生效")
	}
	response, err := c.getWorkspaceFileSpace(ctx, params)
	if err != nil {
		return grpcbase.Response{
			Message: err.Error(),
		}
	}
	r := response.(grpcbase.Response)
	return r
}
func (c *clientBinding) GetWorkspaceBindProject(ctx context.Context, params GetWorkspaceBindProjectRequest) grpcbase.Response {
	if ctx == nil {
		ctx = context.Background()
		log.Println("GRPC：GetWorkspaceBindProject request context is nil，trace span将无法生效")
	}
	response, err := c.getWorkspaceBindProject(ctx, params)
	if err != nil {
		return grpcbase.Response{
			Message: err.Error(),
		}
	}
	r := response.(grpcbase.Response)
	return r
}
func (c *clientBinding) GetWorkspaceExternalMember(ctx context.Context, params GetWorkspaceExternalMemberRequest) grpcbase.Response {
	if ctx == nil {
		ctx = context.Background()
		log.Println("GRPC：GetWorkspaceExternalMember request context is nil，trace span将无法生效")
	}
	response, err := c.getWorkspaceExternalMember(ctx, params)
	if err != nil {
		return grpcbase.Response{
			Message: err.Error(),
		}
	}
	r := response.(grpcbase.Response)
	return r
}
func (c *clientBinding) ExpireEquityUpdate(ctx context.Context, params ExpireEquityUpdateRequest) grpcbase.Response {
	if ctx == nil {
		ctx = context.Background()
		log.Println("GRPC：ExpireEquityUpdate request context is nil，trace span将无法生效")
	}
	response, err := c.expireEquityUpdate(ctx, params)
	if err != nil {
		return grpcbase.Response{
			Message: err.Error(),
		}
	}
	r := response.(grpcbase.Response)
	return r
}
func (c *clientBinding) UpdateWorkspaceLevel(ctx context.Context, params UpdateWorkspaceLevelRequest) grpcbase.Response {
	if ctx == nil {
		ctx = context.Background()
		log.Println("GRPC：UpdateWorkspaceLevel request context is nil，trace span将无法生效")
	}
	response, err := c.updateWorkspaceLevel(ctx, params)
	if err != nil {
		return grpcbase.Response{
			Message: err.Error(),
		}
	}
	r := response.(grpcbase.Response)
	return r
}
func (c *clientBinding) GetTaskWorkspace(ctx context.Context, params GetTaskWorkspaceRequest) grpcbase.Response {
	if ctx == nil {
		ctx = context.Background()
		log.Println("GRPC：GetTaskWorkspace request context is nil，trace span将无法生效")
	}
	response, err := c.getTaskWorkspace(ctx, params)
	if err != nil {
		return grpcbase.Response{
			Message: err.Error(),
		}
	}
	r := response.(grpcbase.Response)
	return r
}
func (c *clientBinding) GetProjectWorkspace(ctx context.Context, params GetProjectWorkspaceRequest) grpcbase.Response {
	if ctx == nil {
		ctx = context.Background()
		log.Println("GRPC：GetProjectWorkspace request context is nil，trace span将无法生效")
	}
	response, err := c.getProjectWorkspace(ctx, params)
	if err != nil {
		return grpcbase.Response{
			Message: err.Error(),
		}
	}
	r := response.(grpcbase.Response)
	return r
}
func (c *clientBinding) AddMemberEquityUpdate(ctx context.Context, params AddMemberEquityUpdateRequest) grpcbase.Response {
	if ctx == nil {
		ctx = context.Background()
		log.Println("GRPC：AddMemberEquityUpdate request context is nil，trace span将无法生效")
	}
	response, err := c.addMemberEquityUpdate(ctx, params)
	if err != nil {
		return grpcbase.Response{
			Message: err.Error(),
		}
	}
	r := response.(grpcbase.Response)
	return r
}
func (c *clientBinding) BatchGetWorkspaceDirectory(ctx context.Context, params BatchGetWorkspaceDirectoryRequest) grpcbase.Response {
	if ctx == nil {
		ctx = context.Background()
		log.Println("GRPC：BatchGetWorkspaceDirectory request context is nil，trace span将无法生效")
	}
	response, err := c.batchGetWorkspaceDirectory(ctx, params)
	if err != nil {
		return grpcbase.Response{
			Message: err.Error(),
		}
	}
	r := response.(grpcbase.Response)
	return r
}
func (c *clientBinding) GetWorkspaceMembersOfTask(ctx context.Context, params GetWorkspaceMembersOfTaskRequest) grpcbase.Response {
	if ctx == nil {
		ctx = context.Background()
		log.Println("GRPC：GetWorkspaceMembersOfTask request context is nil，trace span将无法生效")
	}
	response, err := c.getWorkspaceMembersOfTask(ctx, params)
	if err != nil {
		return grpcbase.Response{
			Message: err.Error(),
		}
	}
	r := response.(grpcbase.Response)
	return r
}
func (c *clientBinding) GetBatchWorkspaceInfo(ctx context.Context, params GetBatchWorkspaceInfoRequest) grpcbase.Response {
	if ctx == nil {
		ctx = context.Background()
		log.Println("GRPC：GetBatchWorkspaceInfo request context is nil，trace span将无法生效")
	}
	response, err := c.getBatchWorkspaceInfo(ctx, params)
	if err != nil {
		return grpcbase.Response{
			Message: err.Error(),
		}
	}
	r := response.(grpcbase.Response)
	return r
}
func (c *clientBinding) CreateNewInvite(ctx context.Context, params CreateNewInviteRequest) grpcbase.Response {
	if ctx == nil {
		ctx = context.Background()
		log.Println("GRPC：CreateNewInvite request context is nil，trace span将无法生效")
	}
	response, err := c.createNewInvite(ctx, params)
	if err != nil {
		return grpcbase.Response{
			Message: err.Error(),
		}
	}
	r := response.(grpcbase.Response)
	return r
}
func (c *clientBinding) GetObjectiveMember(ctx context.Context, params GetObjectiveMemberRequest) grpcbase.Response {
	if ctx == nil {
		ctx = context.Background()
		log.Println("GRPC：GetObjectiveMember request context is nil，trace span将无法生效")
	}
	response, err := c.getObjectiveMember(ctx, params)
	if err != nil {
		return grpcbase.Response{
			Message: err.Error(),
		}
	}
	r := response.(grpcbase.Response)
	return r
}
func (c *clientBinding) GetObjective(ctx context.Context, params GetObjectiveRequest) grpcbase.Response {
	if ctx == nil {
		ctx = context.Background()
		log.Println("GRPC：GetObjective request context is nil，trace span将无法生效")
	}
	response, err := c.getObjective(ctx, params)
	if err != nil {
		return grpcbase.Response{
			Message: err.Error(),
		}
	}
	r := response.(grpcbase.Response)
	return r
}
func (c *clientBinding) GetScreen(ctx context.Context, params GetScreenRequest) grpcbase.Response {
	if ctx == nil {
		ctx = context.Background()
		log.Println("GRPC：GetScreen request context is nil，trace span将无法生效")
	}
	response, err := c.getScreen(ctx, params)
	if err != nil {
		return grpcbase.Response{
			Message: err.Error(),
		}
	}
	r := response.(grpcbase.Response)
	return r
}
func (c *clientBinding) GetScreenMembers(ctx context.Context, params GetScreenMembersRequest) grpcbase.Response {
	if ctx == nil {
		ctx = context.Background()
		log.Println("GRPC：GetScreenMembers request context is nil，trace span将无法生效")
	}
	response, err := c.getScreenMembers(ctx, params)
	if err != nil {
		return grpcbase.Response{
			Message: err.Error(),
		}
	}
	r := response.(grpcbase.Response)
	return r
}
func (c *clientBinding) GetScreensByCardType(ctx context.Context, params GetScreensByCardTypeRequest) grpcbase.Response {
	if ctx == nil {
		ctx = context.Background()
		log.Println("GRPC：GetScreensByCardType request context is nil，trace span将无法生效")
	}
	response, err := c.getScreensByCardType(ctx, params)
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

	c.getTask = grpcbase.CreateGRPCClientEndpoint(cc, "pb.Task",
		"GetTask",
		encodeGetTaskRequest,
		decodeGetTaskResponse)
	c.getTasksName = grpcbase.CreateGRPCClientEndpoint(cc, "pb.Task",
		"GetTasksName",
		encodeGetTasksNameRequest,
		decodeGetTasksNameResponse)
	c.getTaskDispatch = grpcbase.CreateGRPCClientEndpoint(cc, "pb.Task",
		"GetTaskDispatch",
		encodeGetTaskDispatchRequest,
		decodeGetTaskDispatchResponse)
	c.getTaskTakers = grpcbase.CreateGRPCClientEndpoint(cc, "pb.Task",
		"GetTaskTakers",
		encodeGetTaskTakersRequest,
		decodeGetTaskTakersResponse)
	c.getPureTaskTakers = grpcbase.CreateGRPCClientEndpoint(cc, "pb.Task",
		"GetPureTaskTakers",
		encodeGetPureTaskTakersRequest,
		decodeGetPureTaskTakersResponse)
	c.taskExists = grpcbase.CreateGRPCClientEndpoint(cc, "pb.Task",
		"TaskExists",
		encodeTaskExistsRequest,
		decodeTaskExistsResponse)
	c.createGuideTask = grpcbase.CreateGRPCClientEndpoint(cc, "pb.Task",
		"CreateGuideTask",
		encodeCreateGuideTaskRequest,
		decodeCreateGuideTaskResponse)
	c.batchQueryTask = grpcbase.CreateGRPCClientEndpoint(cc, "pb.Task",
		"BatchQueryTask",
		encodeBatchQueryTaskRequest,
		decodeBatchQueryTaskResponse)
	c.getTaskRelevantTakers = grpcbase.CreateGRPCClientEndpoint(cc, "pb.Task",
		"GetTaskRelevantTakers",
		encodeGetTaskRelevantTakersRequest,
		decodeGetTaskRelevantTakersResponse)
	c.updateTaskRelation = grpcbase.CreateGRPCClientEndpoint(cc, "pb.Task",
		"UpdateTaskRelation",
		encodeUpdateTaskRelationRequest,
		decodeUpdateTaskRelationResponse)
	c.getTaskShare = grpcbase.CreateGRPCClientEndpoint(cc, "pb.Task",
		"GetTaskShare",
		encodeGetTaskShareRequest,
		decodeGetTaskShareResponse)
	c.createUserInteract = grpcbase.CreateGRPCClientEndpoint(cc, "pb.Task",
		"CreateUserInteract",
		encodeCreateUserInteractRequest,
		decodeCreateUserInteractResponse)
	c.updateUserInteract = grpcbase.CreateGRPCClientEndpoint(cc, "pb.Task",
		"UpdateUserInteract",
		encodeUpdateUserInteractRequest,
		decodeUpdateUserInteractResponse)
	c.createTaskActiveDetail = grpcbase.CreateGRPCClientEndpoint(cc, "pb.Task",
		"CreateTaskActiveDetail",
		encodeCreateTaskActiveDetailRequest,
		decodeCreateTaskActiveDetailResponse)
	c.updateProjectRelation = grpcbase.CreateGRPCClientEndpoint(cc, "pb.Task",
		"UpdateProjectRelation",
		encodeUpdateProjectRelationRequest,
		decodeUpdateProjectRelationResponse)
	c.getProjectMembers = grpcbase.CreateGRPCClientEndpoint(cc, "pb.Task",
		"GetProjectMembers",
		encodeGetProjectMembersRequest,
		decodeGetProjectMembersResponse)
	c.getProjectsName = grpcbase.CreateGRPCClientEndpoint(cc, "pb.Task",
		"GetProjectsName",
		encodeGetProjectsNameRequest,
		decodeGetProjectsNameResponse)
	c.getUsersMorningAndEveningProgram = grpcbase.CreateGRPCClientEndpoint(cc, "pb.Task",
		"GetUsersMorningAndEveningProgram",
		encodeGetUsersMorningAndEveningProgramRequest,
		decodeGetUsersMorningAndEveningProgramResponse)
	c.getTaskWithParent = grpcbase.CreateGRPCClientEndpoint(cc, "pb.Task",
		"GetTaskWithParent",
		encodeGetTaskWithParentRequest,
		decodeGetTaskWithParentResponse)
	c.getProject = grpcbase.CreateGRPCClientEndpoint(cc, "pb.Task",
		"GetProject",
		encodeGetProjectRequest,
		decodeGetProjectResponse)
	c.getWorkspace = grpcbase.CreateGRPCClientEndpoint(cc, "pb.Task",
		"GetWorkspace",
		encodeGetWorkspaceRequest,
		decodeGetWorkspaceResponse)
	c.getWorkspaceMembers = grpcbase.CreateGRPCClientEndpoint(cc, "pb.Task",
		"GetWorkspaceMembers",
		encodeGetWorkspaceMembersRequest,
		decodeGetWorkspaceMembersResponse)
	c.getWorkspaceList = grpcbase.CreateGRPCClientEndpoint(cc, "pb.Task",
		"GetWorkspaceList",
		encodeGetWorkspaceListRequest,
		decodeGetWorkspaceListResponse)
	c.getTaskChildren = grpcbase.CreateGRPCClientEndpoint(cc, "pb.Task",
		"GetTaskChildren",
		encodeGetTaskChildrenRequest,
		decodeGetTaskChildrenResponse)
	c.getWorkspaceFileSpace = grpcbase.CreateGRPCClientEndpoint(cc, "pb.Task",
		"GetWorkspaceFileSpace",
		encodeGetWorkspaceFileSpaceRequest,
		decodeGetWorkspaceFileSpaceResponse)
	c.getWorkspaceBindProject = grpcbase.CreateGRPCClientEndpoint(cc, "pb.Task",
		"GetWorkspaceBindProject",
		encodeGetWorkspaceBindProjectRequest,
		decodeGetWorkspaceBindProjectResponse)
	c.getWorkspaceExternalMember = grpcbase.CreateGRPCClientEndpoint(cc, "pb.Task",
		"GetWorkspaceExternalMember",
		encodeGetWorkspaceExternalMemberRequest,
		decodeGetWorkspaceExternalMemberResponse)
	c.expireEquityUpdate = grpcbase.CreateGRPCClientEndpoint(cc, "pb.Task",
		"ExpireEquityUpdate",
		encodeExpireEquityUpdateRequest,
		decodeExpireEquityUpdateResponse)
	c.updateWorkspaceLevel = grpcbase.CreateGRPCClientEndpoint(cc, "pb.Task",
		"UpdateWorkspaceLevel",
		encodeUpdateWorkspaceLevelRequest,
		decodeUpdateWorkspaceLevelResponse)
	c.getTaskWorkspace = grpcbase.CreateGRPCClientEndpoint(cc, "pb.Task",
		"GetTaskWorkspace",
		encodeGetTaskWorkspaceRequest,
		decodeGetTaskWorkspaceResponse)
	c.getProjectWorkspace = grpcbase.CreateGRPCClientEndpoint(cc, "pb.Task",
		"GetProjectWorkspace",
		encodeGetProjectWorkspaceRequest,
		decodeGetProjectWorkspaceResponse)
	c.addMemberEquityUpdate = grpcbase.CreateGRPCClientEndpoint(cc, "pb.Task",
		"AddMemberEquityUpdate",
		encodeAddMemberEquityUpdateRequest,
		decodeAddMemberEquityUpdateResponse)
	c.batchGetWorkspaceDirectory = grpcbase.CreateGRPCClientEndpoint(cc, "pb.Task",
		"BatchGetWorkspaceDirectory",
		encodeBatchGetWorkspaceDirectoryRequest,
		decodeBatchGetWorkspaceDirectoryResponse)
	c.getWorkspaceMembersOfTask = grpcbase.CreateGRPCClientEndpoint(cc, "pb.Task",
		"GetWorkspaceMembersOfTask",
		encodeGetWorkspaceMembersOfTaskRequest,
		decodeGetWorkspaceMembersOfTaskResponse)
	c.getBatchWorkspaceInfo = grpcbase.CreateGRPCClientEndpoint(cc, "pb.Task",
		"GetBatchWorkspaceInfo",
		encodeGetBatchWorkspaceInfoRequest,
		decodeGetBatchWorkspaceInfoResponse)
	c.createNewInvite = grpcbase.CreateGRPCClientEndpoint(cc, "pb.Task",
		"CreateNewInvite",
		encodeCreateNewInviteRequest,
		decodeCreateNewInviteResponse)
	c.getObjectiveMember = grpcbase.CreateGRPCClientEndpoint(cc, "pb.Task",
		"GetObjectiveMember",
		encodeGetObjectiveMemberRequest,
		decodeGetObjectiveMemberResponse)
	c.getObjective = grpcbase.CreateGRPCClientEndpoint(cc, "pb.Task",
		"GetObjective",
		encodeGetObjectiveRequest,
		decodeGetObjectiveResponse)
	c.getScreen = grpcbase.CreateGRPCClientEndpoint(cc, "pb.Task",
		"GetScreen",
		encodeGetScreenRequest,
		decodeGetScreenResponse)
	c.getScreenMembers = grpcbase.CreateGRPCClientEndpoint(cc, "pb.Task",
		"GetScreenMembers",
		encodeGetScreenMembersRequest,
		decodeGetScreenMembersResponse)
	c.getScreensByCardType = grpcbase.CreateGRPCClientEndpoint(cc, "pb.Task",
		"GetScreensByCardType",
		encodeGetScreensByCardTypeRequest,
		decodeGetScreensByCardTypeResponse)
}
