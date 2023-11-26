// Node.js High-Throughput WebSocket Gateway for Market Depth
const express = require('express');
const http = require('http');
const WebSocket = require('ws');

const app = express();
const server = http.createServer(app);
const wss = new WebSocket.Server({ server });

const subscriptions = new Map(); // symbol -> Set of ws clients

wss.on('connection', (ws) => {
  console.log('Trading client connected');

  ws.on('message', (message) => {
    try {
      const data = JSON.parse(message);
      if (data.action === 'SUBSCRIBE') {
        const symbol = data.symbol;
        if (!subscriptions.has(symbol)) subscriptions.set(symbol, new Set());
        subscriptions.get(symbol).add(ws);
        console.log(`Client subscribed to ${symbol}`);
      }
    } catch (e) {
      console.error('Invalid WS message', e);
    }
  });

  ws.on('close', () => {
    for (const clients of subscriptions.values()) {
      clients.delete(ws);
    }
  });
});

// Broadcast Market Depth update to subscribed frontend clients
function broadcastMarketDepth(symbol, depth) {
  const clients = subscriptions.get(symbol);
  if (!clients || clients.size === 0) return;

  const payload = JSON.stringify({ type: 'DEPTH_UPDATE', symbol, depth });
  for (const client of clients) {
    if (client.readyState === WebSocket.OPEN) {
      client.send(payload);
    }
  }
}

server.listen(8080, () => {
  console.log('Trading WebSocket Gateway listening on port 8080');
});
