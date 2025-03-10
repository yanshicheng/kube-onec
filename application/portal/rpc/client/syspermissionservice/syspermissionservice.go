// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: portal.proto

package syspermissionservice

import (
	"context"

	"github.com/yanshicheng/kube-onec/application/portal/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AddSysDictItemReq           = pb.AddSysDictItemReq
	AddSysDictItemResp          = pb.AddSysDictItemResp
	AddSysDictReq               = pb.AddSysDictReq
	AddSysDictResp              = pb.AddSysDictResp
	AddSysMenuReq               = pb.AddSysMenuReq
	AddSysMenuResp              = pb.AddSysMenuResp
	AddSysOrganizationReq       = pb.AddSysOrganizationReq
	AddSysOrganizationResp      = pb.AddSysOrganizationResp
	AddSysPermissionReq         = pb.AddSysPermissionReq
	AddSysPermissionResp        = pb.AddSysPermissionResp
	AddSysPositionReq           = pb.AddSysPositionReq
	AddSysPositionResp          = pb.AddSysPositionResp
	AddSysRoleReq               = pb.AddSysRoleReq
	AddSysRoleResp              = pb.AddSysRoleResp
	AddSysUserReq               = pb.AddSysUserReq
	AddSysUserResp              = pb.AddSysUserResp
	BindRoleMenuReq             = pb.BindRoleMenuReq
	BindRoleMenuResp            = pb.BindRoleMenuResp
	BindRolePermissionReq       = pb.BindRolePermissionReq
	BindRolePermissionResp      = pb.BindRolePermissionResp
	BindRoleReq                 = pb.BindRoleReq
	BindRoleResp                = pb.BindRoleResp
	ChangePasswordReq           = pb.ChangePasswordReq
	ChangePasswordResp          = pb.ChangePasswordResp
	CheckDictItemCodeReq        = pb.CheckDictItemCodeReq
	CheckDictItemCodeResp       = pb.CheckDictItemCodeResp
	DelSysDictItemReq           = pb.DelSysDictItemReq
	DelSysDictItemResp          = pb.DelSysDictItemResp
	DelSysDictReq               = pb.DelSysDictReq
	DelSysDictResp              = pb.DelSysDictResp
	DelSysMenuReq               = pb.DelSysMenuReq
	DelSysMenuResp              = pb.DelSysMenuResp
	DelSysOrganizationReq       = pb.DelSysOrganizationReq
	DelSysOrganizationResp      = pb.DelSysOrganizationResp
	DelSysPermissionReq         = pb.DelSysPermissionReq
	DelSysPermissionResp        = pb.DelSysPermissionResp
	DelSysPositionReq           = pb.DelSysPositionReq
	DelSysPositionResp          = pb.DelSysPositionResp
	DelSysRoleReq               = pb.DelSysRoleReq
	DelSysRoleResp              = pb.DelSysRoleResp
	DelSysUserReq               = pb.DelSysUserReq
	DelSysUserResp              = pb.DelSysUserResp
	FrozenAccountsReq           = pb.FrozenAccountsReq
	FrozenAccountsResp          = pb.FrozenAccountsResp
	GetDictItemNameReq          = pb.GetDictItemNameReq
	GetDictItemTextResp         = pb.GetDictItemTextResp
	GetMenuTreeReq              = pb.GetMenuTreeReq
	GetMenuTreeResp             = pb.GetMenuTreeResp
	GetOrganizationTreeReq      = pb.GetOrganizationTreeReq
	GetRoleByUserIdReq          = pb.GetRoleByUserIdReq
	GetRoleByUserIdResp         = pb.GetRoleByUserIdResp
	GetSysDictByIdReq           = pb.GetSysDictByIdReq
	GetSysDictByIdResp          = pb.GetSysDictByIdResp
	GetSysDictItemByIdReq       = pb.GetSysDictItemByIdReq
	GetSysDictItemByIdResp      = pb.GetSysDictItemByIdResp
	GetSysMenuByIdReq           = pb.GetSysMenuByIdReq
	GetSysMenuByIdResp          = pb.GetSysMenuByIdResp
	GetSysOrganizationByIdReq   = pb.GetSysOrganizationByIdReq
	GetSysOrganizationByIdResp  = pb.GetSysOrganizationByIdResp
	GetSysPermissionByIdReq     = pb.GetSysPermissionByIdReq
	GetSysPermissionByIdResp    = pb.GetSysPermissionByIdResp
	GetSysPermissionTreeReq     = pb.GetSysPermissionTreeReq
	GetSysPermissionTreeResp    = pb.GetSysPermissionTreeResp
	GetSysPositionByIdReq       = pb.GetSysPositionByIdReq
	GetSysPositionByIdResp      = pb.GetSysPositionByIdResp
	GetSysRoleByIdReq           = pb.GetSysRoleByIdReq
	GetSysRoleByIdResp          = pb.GetSysRoleByIdResp
	GetSysUserByIdReq           = pb.GetSysUserByIdReq
	GetSysUserByIdResp          = pb.GetSysUserByIdResp
	GetTokenRequest             = pb.GetTokenRequest
	GetTokenResponse            = pb.GetTokenResponse
	GetUserInfoReq              = pb.GetUserInfoReq
	GetUserInfoResp             = pb.GetUserInfoResp
	LeaveReq                    = pb.LeaveReq
	LeaveResp                   = pb.LeaveResp
	LogoutRequest               = pb.LogoutRequest
	LogoutResponse              = pb.LogoutResponse
	MenuNode                    = pb.MenuNode
	RefreshTokenRequest         = pb.RefreshTokenRequest
	RefreshTokenResponse        = pb.RefreshTokenResponse
	ResetPasswordReq            = pb.ResetPasswordReq
	ResetPasswordResp           = pb.ResetPasswordResp
	RouteMeta                   = pb.RouteMeta
	SearchRoleMenuReq           = pb.SearchRoleMenuReq
	SearchRoleMenuResp          = pb.SearchRoleMenuResp
	SearchRolePermissionIdsReq  = pb.SearchRolePermissionIdsReq
	SearchRolePermissionIdsResp = pb.SearchRolePermissionIdsResp
	SearchRolePermissionReq     = pb.SearchRolePermissionReq
	SearchRolePermissionResp    = pb.SearchRolePermissionResp
	SearchSysDictItemReq        = pb.SearchSysDictItemReq
	SearchSysDictItemResp       = pb.SearchSysDictItemResp
	SearchSysDictReq            = pb.SearchSysDictReq
	SearchSysDictResp           = pb.SearchSysDictResp
	SearchSysMenuReq            = pb.SearchSysMenuReq
	SearchSysMenuResp           = pb.SearchSysMenuResp
	SearchSysOrganizationReq    = pb.SearchSysOrganizationReq
	SearchSysOrganizationResp   = pb.SearchSysOrganizationResp
	SearchSysPermissionReq      = pb.SearchSysPermissionReq
	SearchSysPermissionResp     = pb.SearchSysPermissionResp
	SearchSysPositionByUserReq  = pb.SearchSysPositionByUserReq
	SearchSysPositionByUserResp = pb.SearchSysPositionByUserResp
	SearchSysPositionReq        = pb.SearchSysPositionReq
	SearchSysPositionResp       = pb.SearchSysPositionResp
	SearchSysRoleReq            = pb.SearchSysRoleReq
	SearchSysRoleResp           = pb.SearchSysRoleResp
	SearchSysUserReq            = pb.SearchSysUserReq
	SearchSysUserResp           = pb.SearchSysUserResp
	SysDict                     = pb.SysDict
	SysDictItem                 = pb.SysDictItem
	SysMenu                     = pb.SysMenu
	SysOrganization             = pb.SysOrganization
	SysOrganizationSearch       = pb.SysOrganizationSearch
	SysPermission               = pb.SysPermission
	SysPermissionTree           = pb.SysPermissionTree
	SysPosition                 = pb.SysPosition
	SysRole                     = pb.SysRole
	SysUser                     = pb.SysUser
	TokenResponse               = pb.TokenResponse
	UpdateGlobalSysUserReq      = pb.UpdateGlobalSysUserReq
	UpdateGlobalSysUserResp     = pb.UpdateGlobalSysUserResp
	UpdateIconReq               = pb.UpdateIconReq
	UpdateIconResp              = pb.UpdateIconResp
	UpdateSysDictItemReq        = pb.UpdateSysDictItemReq
	UpdateSysDictItemResp       = pb.UpdateSysDictItemResp
	UpdateSysDictReq            = pb.UpdateSysDictReq
	UpdateSysDictResp           = pb.UpdateSysDictResp
	UpdateSysMenuReq            = pb.UpdateSysMenuReq
	UpdateSysMenuResp           = pb.UpdateSysMenuResp
	UpdateSysOrganizationReq    = pb.UpdateSysOrganizationReq
	UpdateSysOrganizationResp   = pb.UpdateSysOrganizationResp
	UpdateSysPermissionReq      = pb.UpdateSysPermissionReq
	UpdateSysPermissionResp     = pb.UpdateSysPermissionResp
	UpdateSysPositionReq        = pb.UpdateSysPositionReq
	UpdateSysPositionResp       = pb.UpdateSysPositionResp
	UpdateSysRoleReq            = pb.UpdateSysRoleReq
	UpdateSysRoleResp           = pb.UpdateSysRoleResp
	UpdateSysUserReq            = pb.UpdateSysUserReq
	UpdateSysUserResp           = pb.UpdateSysUserResp
	UploadImageRequest          = pb.UploadImageRequest
	UploadImageResponse         = pb.UploadImageResponse
	VerifyTokenRequest          = pb.VerifyTokenRequest
	VerifyTokenResponse         = pb.VerifyTokenResponse

	SysPermissionService interface {
		// -----------------------权限表-----------------------
		AddSysPermission(ctx context.Context, in *AddSysPermissionReq, opts ...grpc.CallOption) (*AddSysPermissionResp, error)
		UpdateSysPermission(ctx context.Context, in *UpdateSysPermissionReq, opts ...grpc.CallOption) (*UpdateSysPermissionResp, error)
		DelSysPermission(ctx context.Context, in *DelSysPermissionReq, opts ...grpc.CallOption) (*DelSysPermissionResp, error)
		GetSysPermissionById(ctx context.Context, in *GetSysPermissionByIdReq, opts ...grpc.CallOption) (*GetSysPermissionByIdResp, error)
		SearchSysPermission(ctx context.Context, in *SearchSysPermissionReq, opts ...grpc.CallOption) (*SearchSysPermissionResp, error)
		GetSysPermissionTree(ctx context.Context, in *GetSysPermissionTreeReq, opts ...grpc.CallOption) (*GetSysPermissionTreeResp, error)
	}

	defaultSysPermissionService struct {
		cli zrpc.Client
	}
)

func NewSysPermissionService(cli zrpc.Client) SysPermissionService {
	return &defaultSysPermissionService{
		cli: cli,
	}
}

// -----------------------权限表-----------------------
func (m *defaultSysPermissionService) AddSysPermission(ctx context.Context, in *AddSysPermissionReq, opts ...grpc.CallOption) (*AddSysPermissionResp, error) {
	client := pb.NewSysPermissionServiceClient(m.cli.Conn())
	return client.AddSysPermission(ctx, in, opts...)
}

func (m *defaultSysPermissionService) UpdateSysPermission(ctx context.Context, in *UpdateSysPermissionReq, opts ...grpc.CallOption) (*UpdateSysPermissionResp, error) {
	client := pb.NewSysPermissionServiceClient(m.cli.Conn())
	return client.UpdateSysPermission(ctx, in, opts...)
}

func (m *defaultSysPermissionService) DelSysPermission(ctx context.Context, in *DelSysPermissionReq, opts ...grpc.CallOption) (*DelSysPermissionResp, error) {
	client := pb.NewSysPermissionServiceClient(m.cli.Conn())
	return client.DelSysPermission(ctx, in, opts...)
}

func (m *defaultSysPermissionService) GetSysPermissionById(ctx context.Context, in *GetSysPermissionByIdReq, opts ...grpc.CallOption) (*GetSysPermissionByIdResp, error) {
	client := pb.NewSysPermissionServiceClient(m.cli.Conn())
	return client.GetSysPermissionById(ctx, in, opts...)
}

func (m *defaultSysPermissionService) SearchSysPermission(ctx context.Context, in *SearchSysPermissionReq, opts ...grpc.CallOption) (*SearchSysPermissionResp, error) {
	client := pb.NewSysPermissionServiceClient(m.cli.Conn())
	return client.SearchSysPermission(ctx, in, opts...)
}

func (m *defaultSysPermissionService) GetSysPermissionTree(ctx context.Context, in *GetSysPermissionTreeReq, opts ...grpc.CallOption) (*GetSysPermissionTreeResp, error) {
	client := pb.NewSysPermissionServiceClient(m.cli.Conn())
	return client.GetSysPermissionTree(ctx, in, opts...)
}
