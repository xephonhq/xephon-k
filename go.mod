module github.com/xephonhq/xephon-k

go 1.13

require (
	github.com/dyweb/go.ice v0.0.0-20180225063146-f8221af2fd71
	github.com/dyweb/gommon v0.0.13
	github.com/gocql/gocql v0.0.0-20200115135732-617765adbe2d
	github.com/gogo/protobuf v1.3.1
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b // indirect
	github.com/golang/protobuf v1.0.0 // indirect
	github.com/libtsdb/libtsdb-go v0.0.1
	github.com/opentracing/opentracing-go v1.0.2 // indirect
	github.com/pkg/errors v0.9.1
	github.com/spf13/cobra v0.0.5
	github.com/stretchr/testify v1.4.0
	golang.org/x/net v0.0.0-20180218175443-cbe0f9307d01
	golang.org/x/sync v0.0.0-20190911185100-cd5d95a43a6e // indirect
	google.golang.org/genproto v0.0.0-20180226182557-2d9486acae19 // indirect
	google.golang.org/grpc v1.10.0
	gopkg.in/yaml.v2 v2.2.7
)

replace github.com/dyweb/go.ice => ../../dyweb/go.ice

replace github.com/libtsdb/libtsdb-go => ../../libtsdb/libtsdb-go
