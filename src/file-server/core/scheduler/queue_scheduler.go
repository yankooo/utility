/*
@Time : 2019/9/18 14:28 
@Author : yanKoo
@File : queue_scheduler
@Software: GoLand
@Description:
*/
package scheduler

import "file-server/model"

// 使用队列来调度任务

type queuedScheduler struct {
	requestChan chan *model.InterphoneMsg		// Request channel
	// Worker channel, 其中每一个Worker是一个 chan model.FileSourceInfo 类型
	workerChan  chan chan *model.InterphoneMsg
}

func NewScheduler() *queuedScheduler {
	return &queuedScheduler{}
}

// 提交请求任务到 requestChannel
func (s *queuedScheduler) Submit(request *model.InterphoneMsg) {
	s.requestChan <- request
}

func (s *queuedScheduler) ConfigMasterWorkerChan(chan *model.InterphoneMsg) {
	panic("implement me")
}

// 告诉外界有一个 worker 可以接收 request
func (s *queuedScheduler) WorkerReady(w chan *model.InterphoneMsg) {
	s.workerChan <- w
}

func (s *queuedScheduler) Run() {
	// 生成channel
	s.workerChan = make(chan chan *model.InterphoneMsg)
	s.requestChan = make(chan *model.InterphoneMsg)
	go func() {
		// 创建请求队列和工作队列
		var requestQ []*model.InterphoneMsg
		var workerQ []chan *model.InterphoneMsg
		for {
			var activeWorker chan *model.InterphoneMsg
			var activeRequest *model.InterphoneMsg

			// 当requestQ和workerQ同时有数据时
			if len(requestQ) > 0 && len(workerQ) > 0 {
				activeWorker = workerQ[0]
				activeRequest = requestQ[0]
			}

			select {
			case r := <-s.requestChan: // 当 requestChan 收到数据
				requestQ = append(requestQ, r)
			case w := <-s.workerChan: // 当 workerChan 收到数据
				workerQ = append(workerQ, w)
			case activeWorker <- activeRequest: // 当请求队列和认读队列都不为空时，给任务队列分配任务
				requestQ = requestQ[1:]
				workerQ = workerQ[1:]
			}
		}
	}()
}
