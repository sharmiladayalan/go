                 START
                   │
        ┌──────────┴──────────┐
        │                     │
        ▼                     ▼
  Rider Location        Order Status
   (3 seconds)          (1 second)
        │                     │
        ▼                     ▼

        ─────── SELECT WAITS ───────

        ┌───────────────────────────┐
        │ case 1: location update   │
        │ case 2: status update     │
        │ case 3: timeout (2s)      │
        └──────────┬────────────────┘
                   │
        ┌──────────┴──────────┐
        │                     │
        ▼                     ▼

   1 sec wins ✔        (timeout ignored)