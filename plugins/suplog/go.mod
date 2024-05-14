module github.com/xlab/catcher/plugins/suplog

go 1.22

replace github.com/bugsnag/bugsnag-go => github.com/xlab/suplog/hooks/bugsnag/bugsnag-go v0.0.0-20220928170903-92b70aa2b602

require github.com/bugsnag/bugsnag-go v0.0.0-00010101000000-000000000000

require (
	github.com/bugsnag/panicwrap v1.3.4 // indirect
	github.com/gofrs/uuid v4.2.0+incompatible // indirect
	github.com/kardianos/osext v0.0.0-20190222173326-2bc1f35cddc0 // indirect
)
