package service

import (
	"kitty/app/model"
	"time"
)

type jobInfoService struct {

}

func (this *jobInfoService)List() ([]model.JobInfo, error) {
	var infos []model.JobInfo
	_, err := ormer.QueryTable("job_info").Filter("state", 0).OrderBy("-create_time").All(&infos)
	return infos, err
}

// 新增
func (this *jobInfoService)Add(jobName, jobGroup, cron, url, params, phone, remark string, active int) error {

	jobInfo := &model.JobInfo{
		JobName:jobName,
		JobGroup:jobGroup,
		Cron:cron,
		Url:url,
		Active:active,
		Params:params,
		Phone:phone,
		Remark:remark,
	}

	_, err := ormer.Insert(jobInfo)
	return err

}

func (this *jobInfoService)FindJobInfoById(id int) (model.JobInfo, error) {

	var jobInfo model.JobInfo
	err := ormer.QueryTable("job_info").Filter("id", id).One(&jobInfo)
	return jobInfo, err

}

func (this *jobInfoService)UpdateJobInfo(id int, url, cron, params, phone, remark string) error {
	jobInfo := &model.JobInfo{Id:id, Url:url, Cron:cron, Params:params, Phone:phone, Remark:remark}
	_, err := ormer.Update(jobInfo, "url", "cron", "params", "phone", "remark")
	return err

}

func (this *jobInfoService)DeleteJobInfoById(id int) error {

	jobInfo := model.JobInfo{Id:id}
	_, err := ormer.Delete(&jobInfo)
	return err
}

type jobHistoryService struct {

}

type jobSnapshotService struct {

}

// 查询任务执行快照列表
func (this *jobSnapshotService)List(state int)([]model.JobSnapshot,error)  {

	var sanpshotList []model.JobSnapshot

	_,err := ormer.QueryTable("job_snapshot").All(&sanpshotList)

	return  sanpshotList,err
}

func (this *jobSnapshotService)FindJobSanpshotById(id ,state int)(model.JobSnapshot,error)  {
	var jobSnapshot model.JobSnapshot

	err:= ormer.QueryTable("job_snapshot").Filter("id",id).One(&jobSnapshot)
	return  jobSnapshot,err
}


func (this *jobSnapshotService)Add(jobSnapshot *model.JobSnapshot)  error {


	id,err := ormer.Insert(jobSnapshot)
	if err!= nil {
		return err
	}

	jobSnapshot.Id = int(id)

	return nil

}

func (this *jobSnapshotService)Update(id,state int,detail string,updateTime time.Time)  error {


	snapshot :=&model.JobSnapshot{
		Id:id,
		State:state,
		Detail:detail,
		UpdateTime:updateTime,
	}
	ormer.Update(snapshot,"state","detail","update_time")
	return nil

}



type jobSnapshotHistoryService struct {

}

