// [[{101.goroutines]]

  /* 
  ```
  | Go worker pool arch. 
  | -----------------------------
  | Pending Tasks Channel   <·· Main goroutines will send new tasks to the cannel
  |
  | 
  | Result Task   Channel   <·· workwhere to sends results to
  | Main Goroutine                     <··  Setups workers at bootstrap,
  |                                         writes to pending task queue at runtime
  |                                         read from Restuls Task Channel
  |
  | Pending       Goroutine Worker 1       Results 
  |    task  ··>  Goroutine Worker 2  ··>     task
  | Channel       Goroutine Worker 3       Channel 
  |               ... 
  |               Goroutine Worker N
  |               └────────┬───────┘
  |              - The pool size is setup at bootstrap
  |              - Workers will read in a balanced way
  |                new tasks from the pending task channel
  |                and output to the result task channel
  ```
  */

package main

import (
    "fmt" 
    "time" 
    "sync" 
)

// STEP 1: Define Worker gorutine

type taskInputData struct {
    op string
    m1 int
    m2 int
}


var task_done = 0

func goworker(
    worker_id int,
    input_tasks    <-chan taskInputData, 
    output_results chan<- int ) {
  for task := range input_tasks {
      result :=  task.m1 * task.m2  // Send result back
      fmt.Printf("Worker %d %s %d %d %d: \n", worker_id, task.op, task.m1, task.m2, result)
      task_done = task_done + 1
      output_results <- result  // Send result back
      time.Sleep(6000)
  }
}

const (
  NUM_WORKERS = 10 // <·· Available gorutines ready fro new tasks.
  NUM_TASKS   = 100 // 
)
func main() {

    // STEP 2: Initialize the Task and Result Channels
    pending_tasks := make(chan taskInputData, 100 /*NUM_TASKS*/ )
    results       := make(chan int, 100 /* NUM_TASKS */ )
    // Create worker gorutines pool with reference to input/ouput/channel
    for i := 1; i <= NUM_WORKERS; i++ {
        fmt.Printf("creating goworker: %d \n", i)
        time.Sleep(10)
        go goworker(i, pending_tasks, results)
    }
 
    for j := 1; j <= NUM_TASKS; j++ {
        pending_tasks <- taskInputData { "op", j, j+1 } // <·· send task (just an int in example)
    }

    fmt.Println("A") 
    go func() { // <·· Launch anonymous Main (Orchestrator Goroutine)
           defer close(pending_tasks)
           defer close(results      )
        // WaitGroup avoid blocking main goroutine while waiting to 
        // finish all tasks. (This is not really needed when waiting for
        // completion in a server gorouting)
   //   fmt.Println("1") 
        wg := sync.WaitGroup{} 
        wg.Add(NUM_TASKS)
        for range results {
            <-results
            wg.Done()
        }
   //   wg.Wait()              // <·· main goroutine waits until all workers are done with their tasks
    }()
    fmt.Println("B") 

    for task_done != NUM_TASKS {
      fmt.Println("task_done: %d",task_done)
      time.Sleep(1 * time.Second)
    }

    /**
     * ^1:
     *  By closing the tasks channel, workers exit gracefully
     *  after processing all tasks. [[doc_has.keypoint]]
     */
}
// [[101.goroutines}]]
