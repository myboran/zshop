# 数据库配置
host=127.0.0.1
port=3306
dbname=zshop
username=root
passwd=123456

#生成的表名
tables=goods
modeldir=./goodsModel
goctl model mysql datasource -url="${username}:${passwd}@tcp(${host}:${port})/${dbname}" -table="${tables}"  -dir="${modeldir}" -cache=false --style=goZero

tables=inventory
modeldir=./inventoryModel
goctl model mysql datasource -url="${username}:${passwd}@tcp(${host}:${port})/${dbname}" -table="${tables}"  -dir="${modeldir}" -cache=false --style=goZero

tables=order_info
modeldir=./orderInfoModel
goctl model mysql datasource -url="${username}:${passwd}@tcp(${host}:${port})/${dbname}" -table="${tables}"  -dir="${modeldir}" -cache=false --style=goZero

tables=order_goods
modeldir=./orderGoodsModel
goctl model mysql datasource -url="${username}:${passwd}@tcp(${host}:${port})/${dbname}" -table="${tables}"  -dir="${modeldir}" -cache=false --style=goZero