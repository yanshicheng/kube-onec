#!/bin/bash

# 默认参数
HOST="172.16.1.61"
PORT="3307"
USER="root"
PASSWORD="123456"
SCHEMA="kube_onec"
TABLE="*"
PACKAGE='pb'
SERVICE_NAME="pb"
GO_PACKAGE="github.com/yanshicheng/kube-onec/application/portal/rpc/pb"
OUTPUT_FILE="./application/portal/rpc/portal1.proto"
# 打印帮助信息
usage() {
    echo "Usage: $0 [-h host] [-P port] [-u user] [-p password] [-s schema] [-k package] [-n service_name] [-g go_package] [-t table] [-o output_file]"
    echo "  -h   MySQL 主机 (默认: $HOST)"
    echo "  -P   MySQL 端口 (默认: $PORT)"
    echo "  -u   MySQL 用户名 (默认: $USER)"
    echo "  -p   MySQL 密码 (默认: 无)"
    echo "  -s   MySQL 数据库 (schema)"
    echo "  -k   protobuf package 名称 (默认: $PACKAGE)"
    echo "  -n   服务名称 (service_name)"
    echo "  -g   Go package 路径 (默认: $GO_PACKAGE)"
    echo "  -t   表名，多个表用','分隔 (默认: $TABLE)"
    echo "  -o   输出文件 (默认: $OUTPUT_FILE)"
    exit 1
}

# 解析命令行参数
while getopts "h:P:u:p:s:k:n:g:t:o:" opt; do
    case $opt in
        h) HOST="$OPTARG" ;;
        P) PORT="$OPTARG" ;;
        u) USER="$OPTARG" ;;
        p) PASSWORD="$OPTARG" ;;
        s) SCHEMA="$OPTARG" ;;
        k) PACKAGE="$OPTARG" ;;
        n) SERVICE_NAME="$OPTARG" ;;
        g) GO_PACKAGE="$OPTARG" ;;
        t) TABLE="$OPTARG" ;;
        o) OUTPUT_FILE="$OPTARG" ;;
        *) usage ;;
    esac
done

# 检查必须的参数
if [ -z "$SCHEMA" ] || [ -z "$SERVICE_NAME" ]; then
    echo "Error: -s (schema) 和 -n (service_name) 是必需的参数."
    usage
fi

# 执行 sql2pb 命令
sql2pb -host "$HOST" -port "$PORT" -user "$USER" -password "$PASSWORD" -schema "$SCHEMA" -package "$PACKAGE" -service_name "$SERVICE_NAME" -go_package "$GO_PACKAGE" -table "$TABLE" > "$OUTPUT_FILE"

# 检查命令是否成功执行
if [ $? -eq 0 ]; then
    echo "Proto 文件已成功生成: $OUTPUT_FILE"
else
    echo "错误: 生成 proto 文件失败."
    exit 1
fi
