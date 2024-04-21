import 'package:isar/isar.dart';

//part 'tick_model.g.dart';

@collection
class MarketTickEntity {
  Id id = Isar.autoIncrement;

  @Index()
  late String symbol;

  late double lastPrice;
  late double open;
  late double high;
  late double low;
  late double close;
  late int volume;

  @Index()
  late DateTime timestamp;
}
