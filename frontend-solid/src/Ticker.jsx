// SolidJS Zero-VDOM Reactive Ticker Benchmark
// Fine-grained reactivity allows updating 100+ ticker cells per frame without React diffing overhead
import { createSignal, onMount, onCleanup } from 'solid-js';

export function Ticker(props) {
  const [price, setPrice] = createSignal(props.initialPrice || 1000.0);
  const [change, setChange] = createSignal(0);

  let timer;
  onMount(() => {
    // Simulate 50 ticks per second
    timer = setInterval(() => {
      const delta = (Math.random() - 0.49) * 2;
      setPrice(p => parseFloat((p + delta).toFixed(2)));
      setChange(delta);
    }, 20);
  });

  onCleanup(() => clearInterval(timer));

  return (
    <div class="ticker-box" style={{ display: 'inline-flex', gap: '8px', padding: '6px 12px', background: '#1e293b' }}>
      <span style={{ 'font-weight': 'bold', color: '#e2e8f0' }}>{props.symbol}</span>
      <span style={{ color: change() >= 0 ? '#4ade80' : '#f87171', 'font-family': 'monospace' }}>
        ₹{price().toFixed(2)}
      </span>
    </div>
  );
}
