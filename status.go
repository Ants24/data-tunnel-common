package datatunnelcommon

type JobStatus uint8

const (
	JobStatusWaiting  JobStatus = 1
	JobStatusRunning  JobStatus = 2
	JobStatusFinished JobStatus = 3
	JobStatusFailed   JobStatus = 4
	JobStatusStopped  JobStatus = 5
)

type JobVerifyStatus uint8

const (
	VerifyStatusWaiting      JobVerifyStatus = 1
	VerifyStatusRunning      JobVerifyStatus = 2
	VerifyStatusConsistent   JobVerifyStatus = 3
	VerifyStatusInconsistent JobVerifyStatus = 4
	VerifyStatusFailed       JobVerifyStatus = 5
)

type JobSection uint8

const (
	JobSectionWaiting   JobSection = 1
	JobSectionAuthCheck JobSection = 2
	JobSectionStruct    JobSection = 3
	JobSectionIncr      JobSection = 4
	JobSectionFull      JobSection = 5
	JobSectionVerify    JobSection = 6
	JobSectionDone      JobSection = 7
)
