module github.com/xlab/catcher/plugins/suplog

go 1.22

replace github.com/bugsnag/bugsnag-go => github.com/xlab/suplog/hooks/bugsnag/bugsnag-go v0.0.0-20220928170903-92b70aa2b602

require (
	github.com/bugsnag/bugsnag-go v1.5.3
	github.com/xlab/catcher v1.1.0
	github.com/xlab/suplog v1.4.4
)

require (
	github.com/bugsnag/panicwrap v1.3.4 // indirect
	github.com/gofrs/uuid v4.2.0+incompatible // indirect
	github.com/kardianos/osext v0.0.0-20190222173326-2bc1f35cddc0 // indirect
	github.com/sirupsen/logrus v1.9.0 // indirect
	github.com/xlab/closer v1.0.0 // indirect
	golang.org/x/sys v0.0.0-20220715151400-c0bba94af5f8 // indirect
)
