package strategy

import "gitlab.33.cn/chain33/chain33/cmd/tools/tasks"

type simpleCreateExecProjStrategy struct {
	strategyBasic
}

func (this *simpleCreateExecProjStrategy) Run() error {
	mlog.Info("Begin run chain33 simple create dapp project.")
	defer mlog.Info("Run chain33 simple create dapp project finish.")
	if err := this.initMember(); err != nil {
		return err
	}
	return this.runImpl()
}

func (this *simpleCreateExecProjStrategy) runImpl() error {
	// 复制模板目录下的文件到指定的目标目录，同时替换掉文件名
	// 遍历目标文件夹内所有文件，替换内部标签
	// 执行shell命令，生成对应的 pb.go 文件
	// 更新引用文件
	var err error
	task := this.buildTask()
	for {
		if task == nil {
			break
		}
		err = task.Execute()
		if err != nil {
			mlog.Error("Execute command failed.", "error", err, "taskname", task.GetName())
			break
		}
		task = task.Next()
	}
	return err
}

func (this *simpleCreateExecProjStrategy) initMember() error {
	return nil
}

func (this *simpleCreateExecProjStrategy) buildTask() tasks.Task {
	return nil
}