# 替换 UpdateTime 为 UpdatedAt
find . -type f -exec sed -i '' 's/UpdateTime/UpdatedAt/g' {} +

# 替换 updateTime 为 updatedAt
find . -type f -exec sed -i '' 's/updateTime/updatedAt/g' {} +


find . -type f -exec sed -i '' 's/CreateTime/createdAt/g' {} +

# 替换 updateTime 为 updatedAt
find . -type f -exec sed -i '' 's/createTime/createdAt/g' {} +
find . -type f -exec sed -i '' 's/createBy/createdBy/g' {} +

find . -type f -exec sed -i '' 's/CreateBy/CreatedBy/g' {} +
find . -type f -exec sed -i '' 's/updateBy/updatedBy/g' {} +
find . -type f -exec sed -i '' 's/UpdateBy/UpdatedBy/g' {} +



表: sys_dict_item   sys_role  sys_dict
修改字段， create_by 为 created_by
         update_by 为 updated_by