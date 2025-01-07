# portal
goctl rpc protoc  -I=application/portal/rpc/proto/ -I=$(go list -f '{{.Dir}}' -m github.com/zeromicro/go-zero) --go_out=application/portal/rpc/ --go_opt=module="github.com/yanshicheng/kube-onec/application/portal/rpc" --go-grpc_out=application/portal/rpc/ --go-grpc_opt=module="github.com/yanshicheng/kube-onec/application/portal/rpc" --zrpc_out=application/portal/rpc/ -m  application/portal/rpc/proto/portal.proto
goctl model mysql datasource -url="root:123456@tcp(172.16.1.61:3307)/kube_onec" -table="sys_user,sys_organizations,sys_positions,sys_role,sys_user_role,sys_menu,sys_role_menu,sys_permission,sys_role_permission,sys_dict,sys_dict_item" -dir="./application/portal/rpc/internal/model" -cache=true --style=goZero
goctl model mysql ddl --src=./sql/user.sql --style=goZero --dir=./application/portal/rpc/internal/model  -c


copier


goctl api go -api=./application/portal/api/portal.api -dir=./application/portal/api/ -style=goZero

bash script/mysql.sh -t sys_user,sys_organization,sys_position,sys_role,sys_user_role,sys_menu,sys_role_menu,sys_permission,sys_role_permission,sys_dict,sys_dict_item
goctl model mysql datasource -url="root:123456@tcp(172.16.1.61:3307)/kube_onec" -table="sys_user,sys_user_position,sys_organization,sys_position,sys_role,sys_user_role,sys_menu,sys_role_menu,sys_permission,sys_role_permission,sys_dict,sys_dict_item" -dir="./application/portal/rpc/internal/model" -cache=true --style=goZero
goctl rpc protoc \
  -I=application/portal/rpc/ \
  -I=$(go list -f '{{.Dir}}' -m github.com/zeromicro/go-zero) \
  --go_out=application/portal/rpc/pb/ \
  --go_opt=module="github.com/yanshicheng/kube-onec/application/portal/rpc/pb" \
  --go-grpc_out=application/portal/rpc/pb/ \
  --go-grpc_opt=module="github.com/yanshicheng/kube-onec/application/portal/rpc/pb"  \
  --zrpc_out=application/portal/rpc/ -m \
  application/portal/rpc/portal.proto

go run application/portal/rpc/portal.go -f application/portal/rpc/etc/portal.yaml
go run application/portal/api/portal.go -f application/portal/api/etc/portal.yaml


# manager
goctl model mysql datasource -url="root:123456@tcp(172.16.1.61:3307)/kube_onec" -table="onec_cluster,onec_cluster_conn_info,onec_node,onec_project,onec_project_admin,onec_project_quota,onec_project_application,onec_resource_taints,onec_resource_annotations,onec_resource_labels" -dir="./application/manager/rpc/internal/model" -cache=true --style=goZero
bash script/mysql.sh -t onec_cluster,onec_cluster_conn_info,onec_node,onec_project,onec_project_admin,onec_project_quota,onec_project_application -o "./application/manager/rpc/manager1.proto" -g "github.com/yanshicheng/kube-onec/application/manager/rpc/pb"

goctl rpc protoc \
  -I=application/manager/rpc/ \
  -I=$(go list -f '{{.Dir}}' -m github.com/zeromicro/go-zero) \
  --go_out=application/manager/rpc/pb/ \
  --go_opt=module="github.com/yanshicheng/kube-onec/application/manager/rpc/pb" \
  --go-grpc_out=application/manager/rpc/pb/ \
  --go-grpc_opt=module="github.com/yanshicheng/kube-onec/application/manager/rpc/pb"  \
  --zrpc_out=application/manager/rpc/ -m \
  application/manager/rpc/manager.proto


goctl api go -api=./application/manager/api/manager.api -dir=./application/manager/api/ -style=goZero
go run application/manager/rpc/manager.go -f application/manager/rpc/etc/manager.yaml
go run application/manager/api/manager.go -f application/manager/api/etc/manager-api.yaml