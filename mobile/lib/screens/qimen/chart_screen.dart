import 'package:flutter/material.dart';
import 'package:google_fonts/google_fonts.dart';
import '../../services/api_service.dart';

class ChartScreen extends StatefulWidget {
  const ChartScreen({super.key});

  @override
  State<ChartScreen> createState() => _ChartScreenState();
}

class _ChartScreenState extends State<ChartScreen> {
  final ApiService _apiService = ApiService();
  DateTime _selectedDateTime = DateTime.now();
  Map<String, dynamic>? _chartData;
  bool _isLoading = false;

  Future<void> _generateChart() async {
    setState(() {
      _isLoading = true;
    });

    try {
      final response = await _apiService.generateChart(_selectedDateTime);
      setState(() {
        _chartData = response['data'];
      });
    } catch (e) {
      if (mounted) {
        ScaffoldMessenger.of(context).showSnackBar(
          SnackBar(content: Text(e.toString())),
        );
      }
    } finally {
      if (mounted) {
        setState(() {
          _isLoading = false;
        });
      }
    }
  }

  @override
  void initState() {
    super.initState();
    _generateChart();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('奇門盤'),
        centerTitle: true,
        actions: [
          IconButton(
            icon: const Icon(Icons.calendar_today),
            onPressed: () async {
              final date = await showDatePicker(
                context: context,
                initialDate: _selectedDateTime,
                firstDate: DateTime(1900),
                lastDate: DateTime(2100),
              );
              if (date != null) {
                setState(() {
                  _selectedDateTime = DateTime(
                    date.year,
                    date.month,
                    date.day,
                    _selectedDateTime.hour,
                    _selectedDateTime.minute,
                  );
                });
                _generateChart();
              }
            },
          ),
          IconButton(
            icon: const Icon(Icons.access_time),
            onPressed: () async {
              final time = await showTimePicker(
                context: context,
                initialTime: TimeOfDay.fromDateTime(_selectedDateTime),
              );
              if (time != null) {
                setState(() {
                  _selectedDateTime = DateTime(
                    _selectedDateTime.year,
                    _selectedDateTime.month,
                    _selectedDateTime.day,
                    time.hour,
                    time.minute,
                  );
                });
                _generateChart();
              }
            },
          ),
        ],
      ),
      body: Column(
        children: [
          Padding(
            padding: const EdgeInsets.all(16.0),
            child: Text(
              '選擇時間：${_selectedDateTime.year}/${_selectedDateTime.month}/${_selectedDateTime.day} ${_selectedDateTime.hour}:${_selectedDateTime.minute.toString().padLeft(2, '0')}',
              style: Theme.of(context).textTheme.titleMedium,
            ),
          ),
          if (_isLoading)
            const Expanded(
              child: Center(
                child: CircularProgressIndicator(),
              ),
            )
          else if (_chartData != null)
            Expanded(
              child: GridView.builder(
                padding: const EdgeInsets.all(16.0),
                gridDelegate: const SliverGridDelegateWithFixedCrossAxisCount(
                  crossAxisCount: 3,
                  childAspectRatio: 1.0,
                  crossAxisSpacing: 8.0,
                  mainAxisSpacing: 8.0,
                ),
                itemCount: 9,
                itemBuilder: (context, index) {
                  final palace = _chartData!['palace_${index + 1}'];
                  return Card(
                    child: Container(
                      decoration: BoxDecoration(
                        border: Border.all(color: Colors.grey),
                        borderRadius: BorderRadius.circular(8.0),
                      ),
                      child: Column(
                        mainAxisAlignment: MainAxisAlignment.center,
                        children: [
                          Text(
                            '宮位 ${index + 1}',
                            style: Theme.of(context).textTheme.titleSmall,
                          ),
                          const SizedBox(height: 8),
                          Text('天干：${palace['heavenly_stem']}'),
                          Text('地支：${palace['earthly_branch']}'),
                          Text('九星：${palace['star']}'),
                          Text('八門：${palace['door']}'),
                          Text('八神：${palace['god']}'),
                          Text('暗干：${palace['hidden_stem']}'),
                        ],
                      ),
                    ),
                  );
                },
              ),
            )
          else
            const Expanded(
              child: Center(
                child: Text('無法載入奇門盤數據'),
              ),
            ),
        ],
      ),
      floatingActionButton: FloatingActionButton(
        onPressed: _isLoading ? null : _generateChart,
        child: const Icon(Icons.refresh),
      ),
    );
  }
} 