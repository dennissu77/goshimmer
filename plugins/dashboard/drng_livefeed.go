package dashboard

import (
	"encoding/hex"
	"time"

	"github.com/iotaledger/hive.go/daemon"
	"github.com/iotaledger/hive.go/events"
	"github.com/iotaledger/hive.go/workerpool"

	drngpkg "github.com/iotaledger/goshimmer/packages/drng"
	"github.com/iotaledger/goshimmer/packages/shutdown"
	"github.com/iotaledger/goshimmer/plugins/drng"
)

var (
	drngLiveFeedWorkerCount     = 1
	drngLiveFeedWorkerQueueSize = 50
	drngLiveFeedWorkerPool      *workerpool.WorkerPool
)

type drngMsg struct {
	Instance      uint32 `json:"instance"`
	Name          string `json:"name"`
	DistributedPK string `json:"dpk"`
	Round         uint64 `json:"round"`
	Randomness    string `json:"randomness"`
	Timestamp     string `json:"timestamp"`
}

func configureDrngLiveFeed() {
	drngLiveFeedWorkerPool = workerpool.New(func(task workerpool.Task) {
		newRandomness := task.Param(0).(*drngpkg.State)

		// assign the name of the instance based on its instanceID
		var name string
		switch newRandomness.Committee().InstanceID {
		case drng.Pollen:
			name = "Pollen"
		case drng.XTeam:
			name = "X-Team"
		case drng.Community:
			name = "Community"
		default:
			name = "Custom"
		}

		broadcastWsMessage(&wsmsg{MsgTypeDrng, &drngMsg{
			Instance:      newRandomness.Committee().InstanceID,
			Name:          name,
			DistributedPK: hex.EncodeToString(newRandomness.Committee().DistributedPK),
			Round:         newRandomness.Randomness().Round,
			Randomness:    hex.EncodeToString(newRandomness.Randomness().Randomness[:32]),
			Timestamp:     newRandomness.Randomness().Timestamp.Format("2 Jan 2006 15:04:05"),
		}})

		task.Return(nil)
	}, workerpool.WorkerCount(drngLiveFeedWorkerCount), workerpool.QueueSize(drngLiveFeedWorkerQueueSize))
}

func runDrngLiveFeed() {
	if err := daemon.BackgroundWorker("Dashboard[DRNGUpdater]", func(shutdownSignal <-chan struct{}) {
		newMsgRateLimiter := time.NewTicker(time.Second / 10)
		defer newMsgRateLimiter.Stop()

		notifyNewRandomness := events.NewClosure(func(message *drngpkg.State) {
			select {
			case <-newMsgRateLimiter.C:
				drngLiveFeedWorkerPool.TrySubmit(message)
			default:
			}
		})
		drng.Instance().Events.Randomness.Attach(notifyNewRandomness)

		drngLiveFeedWorkerPool.Start()
		defer drngLiveFeedWorkerPool.Stop()

		<-shutdownSignal
		log.Info("Stopping Dashboard[DRNGUpdater] ...")
		drng.Instance().Events.Randomness.Detach(notifyNewRandomness)
		log.Info("Stopping Dashboard[DRNGUpdater] ... done")
	}, shutdown.PriorityDashboard); err != nil {
		log.Panicf("Failed to start as daemon: %s", err)
	}
}
