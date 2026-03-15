# High-Performance Stock Trading Engine & Multi-Platform Terminal

A low-latency electronic trading platform built for institutional and retail brokerage (AC Agarwal tenure, Sep 2023 – Present).

## Core Architecture
- **Order Matching Engine (Go)**: Price-time priority (FIFO) matching algorithm handling 150,000+ orders/second with microsecond internal latency.
- **Market Data Feed Parser (Go)**: Binary decoder unpacking multicast UDP tick packets (NSE/BSE simulation).
- **Streaming Gateway (Node.js)**: WebSocket fan-out distributing Level-2 Market Depth to thousands of concurrent trader clients.
- **Web Terminals**:
  - React 18 Trading Terminal with Level-2 ladder.
  - SolidJS ultra-lightweight reactive ticker widget benchmarking zero-VDOM UI updates at 60fps.
- **Mobile App (Flutter + Isar DB)**: Mobile trading application with **Isar NoSQL Database** for local tick caching, offline charts, and fast watchlists.
- **Desktop Prototype**: Low-latency Rust + Wails / Tauri trading terminal experiment.
- **Cloud & Streaming**: Apache Kafka event streaming and Kubernetes HPA deployment manifests.

Created by Dharmesh Patel <dharmkp@gmail.com> (CTO, AC Agarwal).
