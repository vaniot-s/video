package taskrunner

import (
	"errors"
	"go-note/mooc/video/scheduler/dbops"
	"log"
	"os"
	"sync"
)

//删除文件
func deleteVideo(vid string) error {
	//ossfn := "videos/" + vid
	//bn := "avenssi-videos2"
	//ok := ossops.DeleteObject(ossfn, bn)
	//log.Printf(VIDEO_PATH + vid)
	err := os.Remove(VIDEO_PATH + vid)
	log.Printf("Deleting video error, %v", err)
	if err != nil && !os.IsNotExist(err) {
		log.Printf("Deleting video error, %v", err)
		return errors.New("Deleting video error")
	}

	return nil
}

//
func VideoClearDispatcher(dc dataChan) error {
	res, err := dbops.ReadVideoDeletionRecord(3) //设置为3每次读取
	if err != nil {
		log.Printf("Video clear dispatcher error: %v", err)
		return err
	}

	if len(res) == 0 {
		return errors.New("All tasks finished")
	}

	for _, id := range res {
		dc <- id
	}

	return nil
}

func VideoClearExecutor(dc dataChan) error {
	errMap := &sync.Map{}
	var err error

forloop:
	for {
		select {
		case vid := <-dc:
			go func(id interface{}) {
				if err := deleteVideo(id.(string)); err != nil {
					errMap.Store(id, err) //线程安全
					return
				}
				if err := dbops.DelVideoDeletionRecord(id.(string)); err != nil {
					errMap.Store(id, err)
					return
				}
			}(vid) //goroutine 不会暂存
		default:
			break forloop
		}
	}

	errMap.Range(func(k, v interface{}) bool {
		err = v.(error)
		if err != nil {
			return false
		}
		return true
	})

	return err
}
