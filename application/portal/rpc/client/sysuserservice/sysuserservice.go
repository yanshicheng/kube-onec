// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: portal.proto

package sysuserservice

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

	SysUserService interface {
		// -----------------------账号信息表-----------------------
		AddSysUser(ctx context.Context, in *AddSysUserReq, opts ...grpc.CallOption) (*AddSysUserResp, error)
		UpdateSysUser(ctx context.Context, in *UpdateSysUserReq, opts ...grpc.CallOption) (*UpdateSysUserResp, error)
		DelSysUser(ctx context.Context, in *DelSysUserReq, opts ...grpc.CallOption) (*DelSysUserResp, error)
		GetSysUserById(ctx context.Context, in *GetSysUserByIdReq, opts ...grpc.CallOption) (*GetSysUserByIdResp, error)
		SearchSysUser(ctx context.Context, in *SearchSysUserReq, opts ...grpc.CallOption) (*SearchSysUserResp, error)
		UpdateGlobalSysUser(ctx context.Context, in *UpdateGlobalSysUserReq, opts ...grpc.CallOption) (*UpdateGlobalSysUserResp, error)
		ResetPassword(ctx context.Context, in *ResetPasswordReq, opts ...grpc.CallOption) (*ResetPasswordResp, error)
		FrozenAccounts(ctx context.Context, in *FrozenAccountsReq, opts ...grpc.CallOption) (*FrozenAccountsResp, error)
		ChangePassword(ctx context.Context, in *ChangePasswordReq, opts ...grpc.CallOption) (*ChangePasswordResp, error)
		Leave(ctx context.Context, in *LeaveReq, opts ...grpc.CallOption) (*LeaveResp, error)
		BindRole(ctx context.Context, in *BindRoleReq, opts ...grpc.CallOption) (*BindRoleResp, error)
		GetRoleByUserId(ctx context.Context, in *GetRoleByUserIdReq, opts ...grpc.CallOption) (*GetRoleByUserIdResp, error)
		GetUserInfo(ctx context.Context, in *GetUserInfoReq, opts ...grpc.CallOption) (*GetUserInfoResp, error)
		UpdateIcon(ctx context.Context, in *UpdateIconReq, opts ...grpc.CallOption) (*UpdateIconResp, error)
	}

	defaultSysUserService struct {
		cli zrpc.Client
	}
)

func NewSysUserService(cli zrpc.Client) SysUserService {
	return &defaultSysUserService{
		cli: cli,
	}
}

// -----------------------账号信息表-----------------------
func (m *defaultSysUserService) AddSysUser(ctx context.Context, in *AddSysUserReq, opts ...grpc.CallOption) (*AddSysUserResp, error) {
	client := pb.NewSysUserServiceClient(m.cli.Conn())
	return client.AddSysUser(ctx, in, opts...)
}

func (m *defaultSysUserService) UpdateSysUser(ctx context.Context, in *UpdateSysUserReq, opts ...grpc.CallOption) (*UpdateSysUserResp, error) {
	client := pb.NewSysUserServiceClient(m.cli.Conn())
	return client.UpdateSysUser(ctx, in, opts...)
}

func (m *defaultSysUserService) DelSysUser(ctx context.Context, in *DelSysUserReq, opts ...grpc.CallOption) (*DelSysUserResp, error) {
	client := pb.NewSysUserServiceClient(m.cli.Conn())
	return client.DelSysUser(ctx, in, opts...)
}

func (m *defaultSysUserService) GetSysUserById(ctx context.Context, in *GetSysUserByIdReq, opts ...grpc.CallOption) (*GetSysUserByIdResp, error) {
	client := pb.NewSysUserServiceClient(m.cli.Conn())
	return client.GetSysUserById(ctx, in, opts...)
}

func (m *defaultSysUserService) SearchSysUser(ctx context.Context, in *SearchSysUserReq, opts ...grpc.CallOption) (*SearchSysUserResp, error) {
	client := pb.NewSysUserServiceClient(m.cli.Conn())
	return client.SearchSysUser(ctx, in, opts...)
}

func (m *defaultSysUserService) UpdateGlobalSysUser(ctx context.Context, in *UpdateGlobalSysUserReq, opts ...grpc.CallOption) (*UpdateGlobalSysUserResp, error) {
	client := pb.NewSysUserServiceClient(m.cli.Conn())
	return client.UpdateGlobalSysUser(ctx, in, opts...)
}

func (m *defaultSysUserService) ResetPassword(ctx context.Context, in *ResetPasswordReq, opts ...grpc.CallOption) (*ResetPasswordResp, error) {
	client := pb.NewSysUserServiceClient(m.cli.Conn())
	return client.ResetPassword(ctx, in, opts...)
}

func (m *defaultSysUserService) FrozenAccounts(ctx context.Context, in *FrozenAccountsReq, opts ...grpc.CallOption) (*FrozenAccountsResp, error) {
	client := pb.NewSysUserServiceClient(m.cli.Conn())
	return client.FrozenAccounts(ctx, in, opts...)
}

func (m *defaultSysUserService) ChangePassword(ctx context.Context, in *ChangePasswordReq, opts ...grpc.CallOption) (*ChangePasswordResp, error) {
	client := pb.NewSysUserServiceClient(m.cli.Conn())
	return client.ChangePassword(ctx, in, opts...)
}

func (m *defaultSysUserService) Leave(ctx context.Context, in *LeaveReq, opts ...grpc.CallOption) (*LeaveResp, error) {
	client := pb.NewSysUserServiceClient(m.cli.Conn())
	return client.Leave(ctx, in, opts...)
}

func (m *defaultSysUserService) BindRole(ctx context.Context, in *BindRoleReq, opts ...grpc.CallOption) (*BindRoleResp, error) {
	client := pb.NewSysUserServiceClient(m.cli.Conn())
	return client.BindRole(ctx, in, opts...)
}

func (m *defaultSysUserService) GetRoleByUserId(ctx context.Context, in *GetRoleByUserIdReq, opts ...grpc.CallOption) (*GetRoleByUserIdResp, error) {
	client := pb.NewSysUserServiceClient(m.cli.Conn())
	return client.GetRoleByUserId(ctx, in, opts...)
}

func (m *defaultSysUserService) GetUserInfo(ctx context.Context, in *GetUserInfoReq, opts ...grpc.CallOption) (*GetUserInfoResp, error) {
	client := pb.NewSysUserServiceClient(m.cli.Conn())
	return client.GetUserInfo(ctx, in, opts...)
}

func (m *defaultSysUserService) UpdateIcon(ctx context.Context, in *UpdateIconReq, opts ...grpc.CallOption) (*UpdateIconResp, error) {
	client := pb.NewSysUserServiceClient(m.cli.Conn())
	return client.UpdateIcon(ctx, in, opts...)
}
