// React 18 Market Depth (Level-2 Orderbook Ladder)
import React, { useState, useEffect } from 'react';

export default function MarketDepthLadder({ symbol = 'TCS' }) {
  const [bids, setBids] = useState([
    { price: 3520.50, qty: 120, orders: 4 },
    { price: 3520.00, qty: 350, orders: 8 },
    { price: 3519.50, qty: 500, orders: 12 },
    { price: 3519.00, qty: 820, orders: 15 },
    { price: 3518.50, qty: 1100, orders: 20 }
  ]);

  const [asks, setAsks] = useState([
    { price: 3521.00, qty: 180, orders: 5 },
    { price: 3521.50, qty: 420, orders: 9 },
    { price: 3522.00, qty: 650, orders: 11 },
    { price: 3522.50, qty: 900, orders: 14 },
    { price: 3523.00, qty: 1400, orders: 22 }
  ]);

  return (
    <div style={{ background: '#0f172a', color: '#f8fafc', padding: 16, borderRadius: 8, width: 480 }}>
      <div style={{ display: 'flex', justifyContent: 'space-between', marginBottom: 12 }}>
        <h4 style={{ margin: 0, color: '#38bdf8' }}>{symbol} Market Depth (L2)</h4>
        <span style={{ color: '#4ade80' }}>LTP: ₹3,520.80</span>
      </div>

      <div style={{ display: 'flex', gap: 12 }}>
        {/* Bid Side */}
        <div style={{ flex: 1 }}>
          <div style={{ fontSize: 12, color: '#94a3b8', borderBottom: '1px solid #334155', paddingBottom: 4 }}>BIDS</div>
          {bids.map((b, i) => (
            <div key={i} style={{ display: 'flex', justifyContent: 'space-between', padding: '4px 0', color: '#4ade80' }}>
              <span>{b.orders}</span>
              <span>{b.qty}</span>
              <strong>₹{b.price.toFixed(2)}</strong>
            </div>
          ))}
        </div>

        {/* Ask Side */}
        <div style={{ flex: 1 }}>
          <div style={{ fontSize: 12, color: '#94a3b8', borderBottom: '1px solid #334155', paddingBottom: 4 }}>ASKS</div>
          {asks.map((a, i) => (
            <div key={i} style={{ display: 'flex', justifyContent: 'space-between', padding: '4px 0', color: '#f87171' }}>
              <strong>₹{a.price.toFixed(2)}</strong>
              <span>{a.qty}</span>
              <span>{a.orders}</span>
            </div>
          ))}
        </div>
      </div>
    </div>
  );
}
