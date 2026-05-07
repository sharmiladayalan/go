                ┌──────────────────────┐
                │       main()         │
                │  create channels     │
                │  start goroutines    │
                └─────────┬────────────┘
                          │
        ┌─────────────────┴─────────────────┐
        │                                   │
        ▼                                   ▼
┌──────────────────┐              ┌──────────────────┐
│ Goroutine 1      │              │ Goroutine 2      │
│                  │              │                  │
│ sleep 2 sec      │              │ sleep 1 sec      │
│ send → c1        │              │ send → c2        │
└─────────┬────────┘              └─────────┬────────┘
          │                                 │
          ▼                                 ▼

        (timing flow)

        0 sec: both waiting
        1 sec: c2 READY ✅
        2 sec: c1 READY ✅


────────────────────────────────────────────────────────

                 SELECT LOOP (main)

────────────────────────────────────────────────────────

        ┌──────────────────────────────┐
        │   select (first iteration)    │
        └─────────────┬────────────────┘
                      │
        ┌─────────────┴─────────────┐
        │                           │
        ▼                           ▼
   c1 NOT READY               c2 READY ✅
                                │
                                ▼
                    "Text message received"
                                │
                                ▼
                     PRINT OUTPUT #1


────────────────────────────────────────────────────────

        ┌──────────────────────────────┐
        │   select (second iteration)   │
        └─────────────┬────────────────┘
                      │
        ┌─────────────┴─────────────┐
        │                           │
        ▼                           ▼
   c1 READY ✅               c2 already used
        │
        ▼
   "Photo uploaded"
        │
        ▼
   PRINT OUTPUT #2

────────────────────────────────────────────────────────

FINAL OUTPUT:

✔ Text message received
✔ Photo uploaded