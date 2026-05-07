START
  │
  ▼
Create channels:
  messages, alerts
  │
  ▼
Start goroutines:
  ├── messages arrives in 3 sec
  └── alerts arrives in 5 sec
  │
  ▼
Loop (every 1 sec):
  │
  ├── select checks:
  │     ├─ messages ready? → print
  │     ├─ alerts ready?   → print
  │     └─ nothing ready   → default
  │
  ▼
Continue loop...
  │
  ▼
END after 7 seconds