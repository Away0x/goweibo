package services

import (
	userModel "gin_weibo/app/models/user"
	viewmodels "gin_weibo/app/view_models"
	"sync"
)

type userViewArr = []*viewmodels.UserViewModel
type idMap = map[uint]*viewmodels.UserViewModel

type userList struct {
	Lock  *sync.Mutex
	IdMap idMap // 由于用了协程，所以依赖这个 map(key 为 id) 来进行排序
}

// UserListService : 查询 user list 并转换为 view models
func UserListService(offset, limit int) userViewArr {
	var (
		// 最后返回的数据
		userViewModels = make(userViewArr, 0)
		// 用于最后排序的 id 列表
		ids = []uint{}

		// chan
		finished = make(chan bool, 1)
		wg       = sync.WaitGroup{}
	)

	ums, err := userModel.List(offset, limit)
	if err != nil {
		return userViewModels
	}

	// 获得 id 列表，记录顺序
	for _, u := range ums {
		ids = append(ids, u.ID)
	}

	userList := userList{
		Lock:  new(sync.Mutex),
		IdMap: make(idMap, len(ums)),
	}

	for _, u := range ums {
		wg.Add(1)

		// 对列表的每一项都做操作，如果操作复杂或条数太多，会造成 api 响应延迟，所以这里使用并行查询
		go func(u *userModel.User) {
			defer wg.Done()

			// 在并发处理中，更新同一个变量为了保证数据一致性，通常需要做锁处理
			userList.Lock.Lock()
			defer userList.Lock.Unlock()
			// 并发时数据被打乱了顺序，所以这里使用 map，id 为 key 以便后续排序
			userList.IdMap[u.ID] = viewmodels.NewUserViewModelSerializer(u) // create view model
		}(u)
	}

	go func() {
		wg.Wait() // 上面多个 goroutine 的并行处理完会发送消息给 finished
		close(finished)
	}()

	// 等待消息 (无可用 case 也无 default 会堵塞)
	select {
	case <-finished:
	}

	// 将 goroutine 中处理过的乱序数据排序
	for _, id := range ids {
		userViewModels = append(userViewModels, userList.IdMap[id])
	}

	return userViewModels
}
