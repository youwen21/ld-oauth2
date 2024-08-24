.PHONY:build build-local

# 程序相关信息
PROGRAM_NAME = gofly-linux-amd64

REMOTE_HOST = myali
REMOTE_PROGRAM_PATH = /data/go-project/gofly

# 说明 V 参数的指定方式
#make build-local V=v1.2.3

buildPkg = gofly
buildTime = `date '+%Y-%m-%d %H:%M:%S'`
buildGo = `go version`
buildUser = `git config user.name`
buildVersion = `git rev-parse HEAD`
buildGitBranch=`git rev-parse --abbrev-ref HEAD`
buildLastCommitUser=`git log -n 1 --format="%an"`
buildLastCommitTime=`git log -1 --date=format:"%Y-%m-%d %T" --format="%ad"`
buildLastCommitId=`git log --format="%H" -n 1`
buildLastCommitMsg=`git log -n 1 --format="%s"`
#buildTag = `git describe --abbrev=0 --tags`  # https://gist.github.com/rponte/fdc0724dd984088606b0
appVersion=$(V)

ldtags = "-X '${buildPkg}/conf.BuildTime=${buildTime}' \
	 -X '${buildPkg}/conf.BuildGo=${buildGo}' \
	 -X '${buildPkg}/conf.BuildUser=${buildUser}' \
	 -X '${buildPkg}/conf.BuildVersion=${buildVersion}' \
	 -X '${buildPkg}/conf.BuildGitBranch=${buildGitBranch}' \
	 -X '${buildPkg}/conf.BuildLastCommitUser=${buildLastCommitUser}' \
	 -X '${buildPkg}/conf.BuildLastCommitTime=${buildLastCommitTime}' \
	 -X '${buildPkg}/conf.BuildLastCommitId=${buildLastCommitId}' \
	 -X '${buildPkg}/conf.BuildLastCommitMsg=${buildLastCommitMsg}' \
	 -X '${buildPkg}/conf.Version=${appVersion}' "

default: deploy-remote

build:
	GOOS=linux GOARCH=amd64 go build -ldflags ${ldtags} -o bin/$(PROGRAM_NAME)

build-local:
	go build -ldflags ${ldtags} -o bin/gofly

deploy-remote:build rsync-remote restart-remote
	echo "build , rsync,  restart finished"

rsync-remote:
	rsync bin/$(PROGRAM_NAME) $(REMOTE_HOST):$(REMOTE_PROGRAM_PATH)/
	rsync start_program.sh $(REMOTE_HOST):$(REMOTE_PROGRAM_PATH)/

# 远程重新启动程序
restart-remote:stop-remote start-remote

# 远程终止程序
stop-remote:
	ssh $(REMOTE_HOST) "pkill -f '$(PROGRAM_NAME)'" || true

# 远程启动程序
start-remote:
	ssh $(REMOTE_HOST) "cd $(REMOTE_PROGRAM_PATH) && ./start_program.sh"

