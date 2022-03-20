package task

import (
	"fmt"
	"github.com/mars-projects/mars/common/framework"
)

type TaskManager struct {
	*framework.TransactionEngine
}

func NewTaskManager(engine *framework.TransactionEngine, sender framework.MessageSender) (*TaskManager, error) {
	var err error
	var manager = TaskManager{engine}
	fmt.Println(framework.QueryMediaImageRequest)
	fmt.Println(framework.SnapshotResumedEvent)
	if err = manager.RegisterExecutor(framework.QueryMediaImageRequest,
		&QueryMediaImageExecutor{sender}); err != nil {
		return nil, err
	}
	//if err = task.RegisterExecutor(framework.GetMediaImageRequest,
	//	&GetMediaImageExecutor{sender, imageBiz}); err != nil {
	//	return nil, err
	//}
	//if err = task.RegisterExecutor(framework.CreateMediaImageRequest,
	//	&CreateMediaImageExecutor{sender, imageBiz}); err != nil {
	//	return nil, err
	//}
	//if err = task.RegisterExecutor(framework.ModifyMediaImageRequest,
	//	&ModifyMediaImageExecutor{sender, imageBiz}); err != nil {
	//	return nil, err
	//}
	//if err = task.RegisterExecutor(framework.DeleteMediaImageRequest,
	//	&DeleteMediaImageExecutor{sender, imageBiz}); err != nil {
	//	return nil, err
	//}
	//
	//if err = task.RegisterExecutor(framework.QueryDiskImageRequest,
	//	&QueryDiskImageExecutor{sender, imageBiz}); err != nil {
	//	return nil, err
	//}
	//if err = task.RegisterExecutor(framework.GetDiskImageRequest,
	//	&GetDiskImageExecutor{sender, imageBiz}); err != nil {
	//	return nil, err
	//}
	//if err = task.RegisterExecutor(framework.CreateDiskImageRequest,
	//	&CreateDiskImageExecutor{sender, imageBiz}); err != nil {
	//	return nil, err
	//}
	//if err = task.RegisterExecutor(framework.ModifyDiskImageRequest,
	//	&ModifyDiskImageExecutor{sender, imageBiz}); err != nil {
	//	return nil, err
	//}
	//if err = task.RegisterExecutor(framework.DeleteDiskImageRequest,
	//	&DeleteDiskImageExecutor{sender, imageBiz}); err != nil {
	//	return nil, err
	//}
	//if err = task.RegisterExecutor(framework.DiskImageUpdatedEvent,
	//	&DiskImageUpdateExecutor{sender, imageBiz}); err != nil {
	//	return nil, err
	//}
	//if err = task.RegisterExecutor(framework.SynchronizeDiskImageRequest,
	//	&SyncDiskImagesExecutor{
	//		Sender:      sender,
	//		ImageServer: imageBiz,
	//	}); err != nil {
	//	err = fmt.Errorf("register sync disk images fail: %s", err.Error())
	//	return nil, err
	//}
	//if err = task.RegisterExecutor(framework.SynchronizeMediaImageRequest,
	//	&SyncMediaImagesExecutor{
	//		Sender:      sender,
	//		ImageServer: imageBiz,
	//	}); err != nil {
	//	err = fmt.Errorf("register sync disk images fail: %s", err.Error())
	//	return nil, err
	//}

	return &manager, nil
}
