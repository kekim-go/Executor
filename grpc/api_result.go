package server

import (
	"context"

	handler "github.com/kekim-go/Executor/handler"
	"github.com/kekim-go/Executor/sqlbuilder"
	"github.com/kekim-go/Executor/sqlexecutor"
	grpc_executor "github.com/kekim-go/Protobuf/gen/proto/executor"
)

type apiResultServer struct {
	handler *handler.ApiResultHandler
}

func newApiResultServer(handler *handler.ApiResultHandler) grpc_executor.ApiResultServiceServer {
	return &apiResultServer{
		handler: handler,
	}
}

// GetApiResult : 실질적인 Api 호출 처리 기능
func (s *apiResultServer) GetApiResult(ctx context.Context, req *grpc_executor.ApiRequest) (*grpc_executor.ApiResult, error) {
	e := new(sqlexecutor.Executor)
	b := new(sqlbuilder.Builder)
	meta := b.GetMeta(s.handler.Ctx.MetaDB, req.GetStageId(), req.GetServiceId())

	//fmt.Printf("%+v", meta)

	searchSQL, matchSQL, countSQL, colType := b.BuildSQL(meta, req)
	data, matchCnt, totalCnt := e.Execute(s.handler.Ctx.DataDB, searchSQL, matchSQL, countSQL, colType)

	page, perPage := sqlbuilder.GetPage(req)

	apiResult := new(grpc_executor.ApiResult)
	apiResult.Data = data
	apiResult.Page = page
	apiResult.PerPage = perPage
	apiResult.CurrentCount = int32(len(data))
	apiResult.MatchCount = int32(matchCnt)
	apiResult.TotalCount = int32(totalCnt)

	return apiResult, nil
}
