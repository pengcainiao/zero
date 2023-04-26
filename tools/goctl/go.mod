module github.com/zeromicro/go-zero/tools/goctl

go 1.15

replace gitlab.flyele.vip/server-side/go-zero/v2 v2.0.71 => ../../../go-zero

require (
	github.com/fatih/structtag v1.2.0
	github.com/go-sql-driver/mysql v1.6.0
	github.com/logrusorgru/aurora v2.0.3+incompatible
	github.com/stretchr/testify v1.7.1
	github.com/urfave/cli v1.22.5
	github.com/zeromicro/antlr v0.0.1
	github.com/zeromicro/ddl-parser v0.0.0-20210712021150-63520aca7348
	github.com/zeromicro/go-zero v1.3.0
	gitlab.flyele.vip/server-side/go-zero/v2 v2.0.71
	k8s.io/api v0.20.12
	k8s.io/apimachinery v0.23.4
	k8s.io/client-go v0.20.12
)
