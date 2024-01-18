package admin

import (
	"context"
	"errors"
	"fmt"
	"github.com/JingruiLea/ad_boost/common/logs"
	"github.com/JingruiLea/ad_boost/model"
	"github.com/JingruiLea/ad_boost/utils"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"net/http"
	"reflect"
)

var resources = map[string]Handler{
	"adgroup":        &CrudHandler[model.AdGroup]{},
	"ad":             &CrudHandler[model.Ad]{},
	"ad_report_item": &CrudHandler[model.AdReportItem]{},

	"online_ad": &OnlineAdHandler{},
}

type Filter map[string]interface{}

type GetManyFilter struct {
	Ids []int64 `json:"id"`
}

type Range struct {
	Start int `json:"start"`
	End   int `json:"end"`
}

type Sort struct {
	Field     string        `json:"field"`
	Direction SortDirection `json:"direction"`
}

type SortDirection string

const (
	ASC  SortDirection = "ASC"
	DESC SortDirection = "DESC"
)

// getList	GET http://my.api.url/posts?sort=["title","ASC"]&range=[0, 24]&filter={"title":"bar"}
// getOne	GET http://my.api.url/posts/123
// getMany	GET http://my.api.url/posts?filter={"id":[123,456,789]}
// getManyReference	GET http://my.api.url/posts?filter={"author_id":345}
// create	POST http://my.api.url/posts
// update	PUT http://my.api.url/posts/123
// delete	DELETE http://my.api.url/posts/123

type Handler interface {
	GetList(ctx context.Context, filter Filter, range_ Range, sort Sort) (interface{}, int64, error)
	GetOne(ctx context.Context, id int64) (interface{}, error)
	GetMany(ctx context.Context, filter GetManyFilter) (interface{}, error)
	GetManyReference(ctx context.Context, filter GetManyFilter) (interface{}, error)
	Create(ctx context.Context, data map[string]interface{}) (interface{}, error)
	Update(ctx context.Context, id int64, data interface{}) (interface{}, error)
	Delete(ctx context.Context, id int64) (interface{}, error)
}

func convert2Displayable(ctx context.Context, ins interface{}) interface{} {
	if displayable, ok := ins.(Displayable); ok {
		tmpDisplayable := displayable.Display()
		for i := range tmpDisplayable {
			if stringer, ok := tmpDisplayable[i].Value.(Stringer); ok {
				tmpDisplayable[i].Value = stringer.String()
			} else if i64, ok := tmpDisplayable[i].Value.(int64); ok {
				tmpDisplayable[i].Value = fmt.Sprintf("%d", i64)
			} else if i64, ok := tmpDisplayable[i].Value.(uint); ok {
				tmpDisplayable[i].Value = fmt.Sprintf("%d", i64)
			}
		}
		return tmpDisplayable
	} else {
		logs.CtxDebugf(ctx, "ins is not Displayable type: %v", reflect.TypeOf(ins))
		return ins
	}
}

func parseActionByUrl(ctx *gin.Context, resourceName string, handler Handler) (interface{}, error) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Header("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
	//allow all headers
	ctx.Header("Access-Control-Allow-Headers", "*")
	ctx.Header("Access-Control-Allow-Credentials", "true")
	ctx.Header("Access-Control-Max-Age", "86400")
	ctx.Header("Access-Control-Expose-Headers", "Content-Range")

	var result interface{}
	var err error
	switch ctx.Request.Method {
	case http.MethodGet:
		if ctx.Param("id") != "" {
			idI64 := utils.Str2I64(ctx.Param("id"))
			if idI64 == 0 {
				logs.CtxErrorf(ctx, "Invalid getOne parameter.")
				return nil, errors.New("invalid getOne parameter")
			}
			result, err = handler.GetOne(ctx, idI64)
			if err != nil {
				logs.CtxErrorf(ctx, "GetOne error: %v", err)
				return nil, err
			}
			return convert2Displayable(ctx, result), err
		} else {
			params := ctx.Request.URL.Query()
			// 解析filter参数为JSON
			var filter Filter = map[string]interface{}{}
			filterParam := params.Get("filter")
			if filterParam != "" {
				if err = jsoniter.Unmarshal([]byte(filterParam), &filter); err != nil {
					logs.CtxErrorf(ctx, "Invalid filter parameter: %v", err)
					return nil, err
				}
			}

			// 解析range参数为JSON
			var range_ Range
			var rangeSlice []int
			rangeParam := params.Get("range")
			if rangeParam != "" {
				if err := jsoniter.Unmarshal([]byte(rangeParam), &rangeSlice); err != nil {
					logs.CtxErrorf(ctx, "Invalid range parameter: %v", err)
					return nil, err
				}
				if len(rangeSlice) != 2 {
					logs.CtxErrorf(ctx, "Invalid range slice: %v", rangeSlice)
					return nil, errors.New("invalid range slice")
				}
				range_.Start = rangeSlice[0]
				range_.End = rangeSlice[1]
			}

			// 解析sort参数为JSON
			var sort Sort
			var sortSlice []string
			sortParam := params.Get("sort")
			if sortParam != "" {
				if err := jsoniter.Unmarshal([]byte(sortParam), &sortSlice); err != nil {
					logs.CtxErrorf(ctx, "Invalid sort parameter: %v", err)
					return nil, err
				}
				if len(sortSlice) != 2 {
					logs.CtxErrorf(ctx, "Invalid sort slice: %v", sortSlice)
					return nil, errors.New("invalid sort slice")
				}
				sort.Field = sortSlice[0]
				sort.Direction = SortDirection(sortSlice[1])
			}
			result, total, err := handler.GetList(ctx, filter, range_, sort)
			if err != nil {
				logs.CtxErrorf(ctx, "GetList error: %v", err)
				return nil, err
			}
			resultValue := reflect.ValueOf(result)
			if resultValue.Kind() != reflect.Slice {
				logs.CtxErrorf(ctx, "GetList result is not slice: %v", result)
				return []interface{}{}, nil
			}
			//Content-Range: posts 0-4/27
			realEnd := range_.Start + resultValue.Len() - 1 //为什么减一？因为range是左闭右开区间
			retResult := make([]interface{}, resultValue.Len())
			for i := 0; i < resultValue.Len(); i++ {
				itemValue := resultValue.Index(i)
				ins := itemValue.Interface()
				retResult[i] = convert2Displayable(ctx, ins)
			}

			ctx.Header("Content-Range", fmt.Sprintf("%s %d-%d/%d", resourceName, range_.Start, realEnd, total))
			return retResult, err
		}
	case http.MethodPost:
		data := ctx.Request.Body
		bytess := make([]byte, ctx.Request.ContentLength)
		if _, err = data.Read(bytess); err != nil {
			logs.CtxErrorf(ctx, "Invalid create parameter: %v", err)
			return nil, err
		}
		createParams := make(map[string]interface{})
		if err = jsoniter.Unmarshal(bytess, &createParams); err != nil {
			logs.CtxErrorf(ctx, "Invalid create parameter: %v", err)
			return nil, err
		}
		result, err = handler.Create(ctx, createParams)
	case http.MethodPut:
		data := ctx.Request.Body
		bytess := make([]byte, ctx.Request.ContentLength)
		if _, err = data.Read(bytess); err != nil {
			logs.CtxErrorf(ctx, "Invalid update parameter: %v", err)
			return nil, err
		}
		updateParams := make(map[string]interface{})
		if err = jsoniter.Unmarshal(bytess, &updateParams); err != nil {
			logs.CtxErrorf(ctx, "Invalid update parameter: %v", err)
			return nil, err
		}
		id := utils.Str2I64(ctx.Param("id"))
		if id == 0 {
			logs.CtxErrorf(ctx, "Invalid update parameter.")
			return nil, errors.New("invalid update parameter")
		}
		result, err = handler.Update(ctx, id, updateParams)
	case http.MethodDelete:
		result, err = handler.Delete(ctx, 0)
	case http.MethodOptions:
		return nil, nil
	default:
		logs.CtxErrorf(ctx, "Invalid method: %s", ctx.Request.Method)
		return nil, errors.New("invalid method")
	}
	return result, err
}

func RegisterAdmin(router *gin.RouterGroup) {

	for resourceName, handler := range resources {
		resourceNameCopy := resourceName
		handlerCopy := handler
		router.Any(resourceName, func(ctx *gin.Context) {
			result, err := parseActionByUrl(ctx, resourceNameCopy, handlerCopy)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			ctx.JSON(http.StatusOK, result)
		})
		router.Any(resourceName+"/:id", func(ctx *gin.Context) {
			result, err := parseActionByUrl(ctx, resourceNameCopy, handlerCopy)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			ctx.JSON(http.StatusOK, result)
		})
	}
}

type Displayable interface {
	Display() utils.SortedList
}

type Stringer interface {
	String() string
}
