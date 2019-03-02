package dbops

import (
	"log"
)

/**
1.user->api  service->delete video
2.api service->scheduler->write video deleteion record
3.timer
4.timer->runner->read WVD->exec->delete video from folder
**/
//添加记录
func AddVideoDeletionRecord(vid string) error {

	stmtIns, err := dbConn.Prepare("INSERT INTO video_del_rec (video_id) VALUES(?)")

	if err != nil {
		return err
	}

	_, err = stmtIns.Exec(vid)
	if err != nil {
		log.Printf("AddVideoDeletionRecord error: %v", err)
		return err
	}

	defer stmtIns.Close()
	return nil
}
