                JOB QUEUE
         ┌──────────────────┐
         │ Order-1          │
         │ Order-2          │
         │ Order-3          │
         │ Order-4          │
         │ Order-5          │
         └──────────────────┘
                 │
      ┌──────────┼──────────┐
      ▼          ▼          ▼
   Worker1    Worker2    Worker3
      │          │          │
      ▼          ▼          ▼
 Process     Process     Process
      │          │          │
      └──────────┼──────────┘
                 ▼
          Send completion signal