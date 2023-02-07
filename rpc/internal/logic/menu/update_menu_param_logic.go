package menu

import (
	"context"

	"github.com/suyuan32/simple-admin-core/pkg/ent"
	"github.com/suyuan32/simple-admin-core/pkg/ent/menu"
	"github.com/suyuan32/simple-admin-core/pkg/i18n"
	"github.com/suyuan32/simple-admin-core/pkg/msg/logmsg"
	"github.com/suyuan32/simple-admin-core/pkg/statuserr"
	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateMenuParamLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateMenuParamLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMenuParamLogic {
	return &UpdateMenuParamLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateMenuParamLogic) UpdateMenuParam(in *core.MenuParamInfo) (*core.BaseResp, error) {
	exist, err := l.svcCtx.DB.Menu.Query().Where(menu.IDEQ(in.MenuId)).Exist(l.ctx)
	if err != nil {
		logx.Errorw(logmsg.DatabaseError, logx.Field("detail", err.Error()))
		return nil, statuserr.NewInternalError(i18n.DatabaseError)
	}

	if !exist {
		logx.Errorw("menu not found", logx.Field("menuId", in.Id))
		return nil, statuserr.NewInvalidArgumentError("menu.menuNotExists")
	}

	err = l.svcCtx.DB.MenuParam.UpdateOneID(in.Id).
		SetType(in.Type).
		SetKey(in.Key).
		SetValue(in.Value).
		SetMenusID(in.MenuId).
		Exec(l.ctx)

	if err != nil {
		switch {
		case ent.IsNotFound(err):
			logx.Errorw(err.Error(), logx.Field("detail", in))
			return nil, statuserr.NewInvalidArgumentError(i18n.TargetNotFound)
		case ent.IsConstraintError(err):
			logx.Errorw(err.Error(), logx.Field("detail", in))
			return nil, statuserr.NewInvalidArgumentError(i18n.UpdateFailed)
		default:
			logx.Errorw(logmsg.DatabaseError, logx.Field("detail", err.Error()))
			return nil, statuserr.NewInternalError(i18n.DatabaseError)
		}
	}

	return &core.BaseResp{Msg: i18n.UpdateSuccess}, nil
}
