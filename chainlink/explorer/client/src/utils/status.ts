import { titleCase } from 'title-case'
import { noCase } from 'change-case'
import { JobRun, TaskRun } from 'explorer/models'

const hasUnfulfilledEthTx = (jobRun: JobRun) => {
  return (jobRun.taskRuns || []).some((tr: TaskRun) => {
    return (
      tr.type === 'ethtx' &&
      tr.status === 'completed' &&
      tr.transactionStatus !== 'fulfilledRunLog'
    )
  })
}

export default (jobRun: JobRun): [string, boolean] => {
  let unfulfilledEthTx = false
  let text

  if (jobRun.status === 'in_progress') {
    text = 'Pending'
    unfulfilledEthTx = hasUnfulfilledEthTx(jobRun)
  } else if (jobRun.status === 'error') {
    text = 'Errored'
    unfulfilledEthTx = hasUnfulfilledEthTx(jobRun)
  } else if (jobRun.status === 'completed') {
    text = 'Complete'
    unfulfilledEthTx = hasUnfulfilledEthTx(jobRun)
  } else {
    text = titleCase(noCase(jobRun.status))
  }

  return [text, unfulfilledEthTx]
}
