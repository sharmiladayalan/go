              CUSTOMER REQUESTS
                     │
                     ▼
        ┌─────────────────────────┐
        │ Buffered Channel Queue  │
        │ Capacity = 3            │
        └─────────────────────────┘
             │     │      │
             ▼     ▼      ▼
          Burger Pizza  Pasta

              (FULL)

          Sandwich waits...

                     │
                     ▼

        Worker consumes Burger

                     │
                     ▼

          Sandwich enters queue

                     │
                     ▼

        Worker processes remaining orders