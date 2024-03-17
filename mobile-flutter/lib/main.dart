import 'package:flutter/material.dart';

void main() {
  runApp(const TradingApp());
}

class TradingApp extends StatelessWidget {
  const TradingApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'AC Trading Desk',
      theme: ThemeData.dark().copyWith(
        scaffoldBackgroundColor: const Color(0xFF0F172A),
        primaryColor: const Color(0xFF0284C7),
      ),
      home: const Scaffold(
        body: Center(
          child: Text('AC Agarwal Mobile Trading Desk'),
        ),
      ),
    );
  }
}
