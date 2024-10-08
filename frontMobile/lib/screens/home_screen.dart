import 'package:flutter/material.dart';

class HomeScreen extends StatelessWidget {
  const HomeScreen({super.key});

  List<String> callFunctionServices() {
    return ["Microsoft", "GitHub", "Google"];
  }

  generateExpandedWidgets(BuildContext context, List<String> platforms) {
    return platforms.map((platform) {
      return Expanded(
        child: Column(
          children: [
            ElevatedButton(
              onPressed: () {
                showDialog(
                  context: context,
                  builder: (context) {
                    return AlertDialog(
                      title: Text(platform),
                      content: Text('Call to API $platform'),
                      actions: [
                        TextButton(
                          onPressed: () {
                            Navigator.of(context).pop();
                          },
                          child: const Text('Close'),
                        ),
                      ],
                    );
                  },
                );
              },
              child: Text(platform),
            ),
          ],
        ),
      );
    }).toList();
  }

  @override
  Widget build(BuildContext context) {
    final List<String> platforms = callFunctionServices();

    return Scaffold(
      appBar: AppBar(
        backgroundColor: Colors.white,
        title: const Text('Home'),
      ),
      body: Container(
        color: Colors.deepPurple,
        alignment: Alignment.topCenter,
        child: Row(
          children: generateExpandedWidgets(context, platforms),
        ),
      ),
    );
  }
}
