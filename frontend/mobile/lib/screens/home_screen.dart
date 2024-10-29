import 'dart:convert';
import 'package:flutter/material.dart';
import 'package:flutter_dotenv/flutter_dotenv.dart';
import 'package:http/http.dart' as http;
import '../component/appBar.dart';
import '../component/navBar.dart';

class HomeScreen extends StatefulWidget {
  const HomeScreen({super.key});

  @override
  State<HomeScreen> createState() => _HomeScreenState();
}

class _HomeScreenState extends State<HomeScreen> {
  List<String> callFunctionServices() {
    return ["Microsoft", "GitHub", "Google"];
  }

  Future<List<dynamic>> getApplets() async {
    String? token = dotenv.env['DEV_TOKEN'];
    String apiAppletUrl = dotenv.env['API_APPLET_URL']!;
    if (token == null) {
      throw Exception("Token non disponible");
    }

    final response = await http.get(
      Uri.parse(apiAppletUrl),
      headers: {
        'access_token': token,
      },
    );

    if (response.statusCode == 200) {
      final data = jsonDecode(response.body);
      return data['applet_array'];
    } else {
      throw Exception("Failed to load applets");
    }
  }

  int _selectedIndex = 0;

  void _onItemTapped(int index) {
    setState(() {
      _selectedIndex = index;
    });
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
              style: ElevatedButton.styleFrom(
                backgroundColor: const Color.fromARGB(255, 217, 217, 217),
                foregroundColor: const Color.fromARGB(255, 43, 42, 40),
                shape: RoundedRectangleBorder(
                  borderRadius: BorderRadius.circular(12),
                ),
                minimumSize: const Size(90, 90),
                padding: const EdgeInsets.all(16),
                elevation: 8,
                shadowColor: Colors.black.withOpacity(0.5),
              ),
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
    appBar: const CustomAppBar(),
      body: Container(
        color: Colors.white,
        alignment: Alignment.topCenter,
        child: Column(
          mainAxisSize: MainAxisSize.max,
          children: [
            Column(
            mainAxisSize: MainAxisSize.max,
            children: [
              const SizedBox(height: 80),
              const Text(
                'Explore',
                style: TextStyle(fontSize: 30, fontWeight: FontWeight.bold),
              ),
              const SizedBox(height: 10),
              const Text(
                'Automate to save time and get more done',
                style: TextStyle(fontSize: 15, fontWeight: FontWeight.bold),
                textAlign: TextAlign.center,
              ),
              const SizedBox(height: 40),
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
                style: ElevatedButton.styleFrom(
                  fixedSize: const Size(317, 43),
                  backgroundColor: const Color.fromARGB(255, 43, 42, 40),
                  foregroundColor: const Color.fromARGB(255, 255, 255, 255),
                  shape: RoundedRectangleBorder(
                    borderRadius: BorderRadius.circular(48),
                  ),
                ),
                child: const Text(
                  'Start Today',
                  style: TextStyle(
                    fontSize: 24,
                    fontWeight: FontWeight.bold,
                  ),
                ),
              ),
              const SizedBox(height: 60),
              Row(
                children: generateExpandedWidgets(context, platforms),
              ),
              const SizedBox(height: 40),
              FutureBuilder<List<dynamic>>(
                future: getApplets(),
                builder: (context, snapshot) {
                  if (snapshot.connectionState == ConnectionState.waiting) {
                    return const CircularProgressIndicator();
                  } else if (snapshot.hasError) {
                    return Text("Erreur : ${snapshot.error}");
                  } else if (!snapshot.hasData || snapshot.data!.isEmpty) {
                    return const Text("Aucun applet trouvé");
                  } else {
                    final applets = snapshot.data!;
                    return Expanded(
                      child: ListView.builder(
                        itemCount: applets.length,
                        shrinkWrap: true,
                        itemBuilder: (context, index) {
                          final applet = applets[index];
                          return ElevatedButton(
                            onPressed: () {
                              showDialog(
                                context: context,
                                builder: (context) {
                                  return AlertDialog(
                                    title: const Text("More information"),
                                    content: Text(
                                        'Status: ${applet['IsOn'] ? 'On' : 'Off'}'),
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
                            style: ElevatedButton.styleFrom(
                                backgroundColor:
                                    const Color.fromARGB(255, 217, 217, 217),
                                foregroundColor:
                                    const Color.fromARGB(255, 43, 42, 40),
                                elevation: 8,
                                shadowColor: Colors.black.withOpacity(0.5),
                                alignment: Alignment.center),
                            child: Text(
                              "ACTION : ${applet['If']}\nREACTION : ${applet['That']}",
                              textAlign: TextAlign.center,
                            ),
                          );
                        },
                      ),
                    );
                  }
                },
              ),
            ]),
          ],
        ),
      ),
      bottomNavigationBar: CustomNavBar(
        currentIndex: _selectedIndex,
        onTap: _onItemTapped,
      ),
    );
  }
}
