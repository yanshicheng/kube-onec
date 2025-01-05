package code

import "github.com/yanshicheng/kube-onec/common/handler/errorx"

var (
	// 通用错误类型（第四位为 0）
	AccountRequiredParams       = errorx.New(101001, "用户名，账号，图标，手机号，邮箱，工号，入职时间为必传参数!")
	ParameterIllegal            = errorx.New(101002, "参数不合法!")
	GenerateUUIDErr             = errorx.New(101003, "生成 UUID 失败!")
	PasswordIllegal             = errorx.New(101004, "密码必须大于 6 位，包含数字、字母、特殊字符!")
	ActionIllegal               = errorx.New(101005, "action 只能是 GET POST PUT DELETE *!")
	LevelIllegal                = errorx.New(101006, "level 只能是 1 2!")
	ParentPermissionNotExist    = errorx.New(101007, "父级权限不存在!")
	DeletePermissionHasChildErr = errorx.New(101008, "删除权限失败，请先清空子权限!")

	// 账号相关错误（第四位为 1）
	EncryptPasswordErr      = errorx.New(101101, "密码加密失败!")
	CreateAccountErr        = errorx.New(101102, "创建账号失败!")
	FindAccountErr          = errorx.New(101103, "查询账号失败!")
	DecodeBase64PasswordErr = errorx.New(101104, "解码密码失败!")
	LoginErr                = errorx.New(101105, "用户名或密码校验失败!")
	PasswordNotMatchErr     = errorx.New(101106, "密码不匹配或新旧密码一致!")
	ChangePasswordErr       = errorx.New(101107, "修改密码失败!")
	FrozenAccountsErr       = errorx.New(101108, "冻结账号操作失败!")
	AccountLockedErr        = errorx.New(101109, "账号被冻结，请联系管理员!")
	AccountLockedTip        = errorx.New(101110, "账号已经设置为离职状态!")
	ResetPasswordErr        = errorx.New(101111, "重置密码失败!")
	ResetPasswordTip        = errorx.New(101112, "登录失败请重置密码!")
	LeaveErr                = errorx.New(101113, "离职操作失败!")
	UpdateSysUserErr        = errorx.New(101114, "更新用户信息失败!")
	FindSysUserListErr      = errorx.New(101115, "查询用户列表失败!")
	NewPasswordNotMatchErr  = errorx.New(101117, "新密码不匹配!")

	// 角色相关错误（第四位为 2）
	CreateRoleErr         = errorx.New(101201, "创建角色失败!")
	FindRoleErr           = errorx.New(101202, "查询角色失败!")
	BindRoleErr           = errorx.New(101203, "绑定角色失败!")
	FindRolePermissionErr = errorx.New(101204, "查询角色权限失败!")
	//批量查询角色权限关系失败
	FindRolePermissionListErr = errorx.New(101205, "批量查询角色权限关系失败!")
	DelBindRolePermissionErr  = errorx.New(101206, "删除角色权限关系失败!")
	BindRolePermissionErr     = errorx.New(101207, "绑定角色权限失败!")
	// 权限相关错误（第四位为 3）
	GetChildPermissionErr = errorx.New(101301, "获取子集权限异常!")
	FindDictItemsErr      = errorx.New(101302, "查询字典数据失败!")
	DictHasItemsErr       = errorx.New(101303, "字典无法删除，请先删除字典数据!")

	// 机构相关错误（第四位为 4）
	CreateOrganizationErr        = errorx.New(101401, "创建机构层级失败!")
	DeleteOrganizationNotNullErr = errorx.New(101402, "删除机构失败，请先清空子机构!")
	GetOrganizationInfoErr       = errorx.New(101403, "机构信息获取失败!")
	GetOrganizationErr           = errorx.New(101404, "机构信息查询出错!")
	GetPositionInfoErr           = errorx.New(101405, "职位信息获取失败!")
	GetPositionErr               = errorx.New(101406, "职位信息查询出错!")

	// Minio 相关错误（第四位为 5）
	MinioCheckErr  = errorx.New(101501, "Minio 检查失败!")
	MinioUploadErr = errorx.New(101502, "Minio 文件上传失败，请检查 Minio 服务!")

	// Token 相关错误（第四位为 6）
	GenerateJWTTokenErr    = errorx.New(101601, "生成 JWT Token 失败!")
	RedisStorageErr        = errorx.New(101602, "Redis 存储失败!")
	UUIDExistErr           = errorx.New(101603, "UUID 已经存在，Token 获取异常!")
	UUIDNotExistErr        = errorx.New(101604, "UUID 查询不到，退出异常!")
	UUIDDeleteErr          = errorx.New(101605, "UUID 删除失败，退出异常!")
	UUIDQueryErr           = errorx.New(101606, "UUID 查询报错，请联系管理员处理!")
	RefreshTokenEmptyErr   = errorx.New(101607, "刷新令牌不能为空!")
	RefreshTokenExpiredErr = errorx.New(101608, "刷新令牌过期!")

	// 数据字典相关
	
	DictNotExistErr = errorx.New(101701, "字典不存在!")
)
