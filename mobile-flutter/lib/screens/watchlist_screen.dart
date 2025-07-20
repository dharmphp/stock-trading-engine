import 'package:flutter/material.dart';

class WatchlistScreen extends StatelessWidget {
  final List<String> symbols = const ['RELIANCE', 'TCS', 'HDFCBANK', 'INFY', 'ICICIBANK', 'ITC', 'SBIN'];

  const WatchlistScreen({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('AC Agarwal Trading Watchlist'),
        backgroundColor: const Color(0xFF1E293B),
      ),
      body: ListView.separated(
        itemCount: symbols.length,
        separatorBuilder: (_, __) => const Divider(color: Color(0xFF334155), height: 1),
        itemBuilder: (ctx, i) {
          return ListTile(
            title: Text(symbols[i], style: const TextStyle(fontWeight: FontWeight.bold)),
            subtitle: const Text('NSE EQ'),
            trailing: Column(
              mainAxisAlignment: MainAxisAlignment.center,
              crossAxisAlignment: CrossAxisAlignment.end,
              children: const [
                Text('₹2,450.00', style: TextStyle(color: Color(0xFF4ADE80), fontWeight: FontWeight.bold)),
                Text('+1.25%', style: TextStyle(color: Color(0xFF4ADE80), fontSize: 12)),
              ],
            ),
          );
        },
      ),
    );
  }
}
