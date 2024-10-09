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
                style: ElevatedButton.styleFrom(
                  backgroundColor: Color.fromARGB(255, 217, 217, 217),
                  foregroundColor: Color.fromARGB(255, 43, 42, 40),
                shape: RoundedRectangleBorder(
                  borderRadius: BorderRadius.circular(12),
                ),
                minimumSize: const Size(90, 90),
                padding: const EdgeInsets.all(16), 
                elevation: 8,
                shadowColor: Colors.black.withOpacity(0.5),  
              ), 
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
        title: const Text('AREA',
          style: TextStyle(
            fontWeight: FontWeight.bold,
          ),
        ),
      ),
      body: Container(
        color: Colors.white,
        alignment: Alignment.topCenter,
        child: Column(
        mainAxisSize: MainAxisSize.max,
        children: [
          const SizedBox(height: 80),
          const Text('Explore',
          style: TextStyle(
            fontSize: 24,
            fontWeight: FontWeight.bold
          ),),
          const SizedBox(height: 60),
          ElevatedButton(
            onPressed: () {
              showDialog(
                context: context,
                builder: (context) {
                  return const AlertDialog(
                    title: Text('Action!'),
                    content: Text('Bouton pressé!'),
                  );
                },
              );
            },
            child: const Text('Start Today',
              style: TextStyle(
                fontSize: 24,
                fontWeight: FontWeight.bold,
            ),),
            style: ElevatedButton.styleFrom(
              fixedSize: Size(317, 43),
              backgroundColor: const Color.fromARGB(255, 43, 42, 40),
              foregroundColor: const Color.fromARGB(255, 255, 255, 255),
              shape: RoundedRectangleBorder(
                borderRadius: BorderRadius.circular(48),
              ),
              // elevation: 8,
              // shadowColor: Colors.grey.withOpacity(0.5),
            ),
          ),
          const SizedBox(height: 60),
          Row(
            children:
              generateExpandedWidgets(context, platforms),
          ),
        ]),
      ),
    );
  }
}
